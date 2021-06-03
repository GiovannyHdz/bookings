package render

import (
	"bytes"
	"fmt"
	"github.com/GiovannyHdz/bookings/pkg/config"
	"github.com/GiovannyHdz/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//functions MAP of functions that could be used in template (HTML file like formatting date)
var functions = template.FuncMap{

}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData set the default data of all pages
func AddDefaultData(td *models.TemplateDate) *models.TemplateDate {

	return td
}

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateDate) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// You can also do it as follows -> myCache[tmpl].Execute(w, nil) instead of creating a new buffer

	//Used to write until w
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	// this gets a list of all files ending with page.tmpl, and stores
	// it in a slice of strings called pages
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// now we loop through the slice of strings, which has two
	// entries: "home.page.tmpl" and "about.page.tmpl"
	for _, page := range pages {
		//get only the file name NOT the full path
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// here, we are checking to see if there are any files at all that
		// end with layout.tmpl. THere is only one, but if there were more
		// than one, we we get them all and store them in a slice of strings
		// named matches
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// if the length of matches is > 0, we go through the slice
		// and parse all of the layouts available to us. We might not use
		// any of them in this iteration through the loop, but if the current
		// template we are working on (home.page.tmpl the first time through)
		// does use a layout, we need to have it available to us before we add it
		// to our template set
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}