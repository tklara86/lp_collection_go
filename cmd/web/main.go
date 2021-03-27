package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tklara86/lp_collection_go/pkg/config"
	"github.com/tklara86/lp_collection_go/pkg/handlers"
	"github.com/tklara86/lp_collection_go/pkg/render"
	"log"
	"net/http"
	"time"
)


const portNumber = ":8080"
// variable App of type AppConfig
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour  // 24h
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)


	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	serve := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}