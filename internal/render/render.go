package render

import (
	"bytes"
	"errors"
	config2 "github.com/GiovannyHdz/bookings/internal/config"
	models2 "github.com/GiovannyHdz/bookings/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//functions MAP of functions that could be used in template (HTML file like formatting date)
var functions = template.FuncMap{

}

var app *config2.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config2.AppConfig) {
	app = a
}

// AddDefaultData set the default data of all pages
func AddDefaultData(td *models2.TemplateDate, r *http.Request) *models2.TemplateDate {
	td.CSRFToken = nosurf.Token(r)
	return td
}

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models2.TemplateDate) error {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Println("Could not get template from template cache")
		return errors.New("Could not get template from template cache")
	}

	//Used to write until w
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser::", err)
		return errors.New("Error writing template to browser::" + err.Error())
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

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