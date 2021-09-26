package main

import (
	"fmt"
	config2 "github.com/GiovannyHdz/bookings/internal/config"
	handlers2 "github.com/GiovannyHdz/bookings/internal/handlers"
	render2 "github.com/GiovannyHdz/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config2.AppConfig
var session *scs.SessionManager

func main() {
	// changes this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render2.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers2.NewRepo(&app)
	handlers2.NewHandlers(repo)
	render2.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
