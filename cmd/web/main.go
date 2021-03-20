package main


import (
	"fmt"
	"github.com/tklara86/lp_collection_go/pkg/config"
	"github.com/tklara86/lp_collection_go/pkg/handlers"
	"github.com/tklara86/lp_collection_go/pkg/render"
	"log"
	"net/http"
)


const portNumber = ":8080"

// main is the main application function
func main() {
	// variable App of type AppConfig
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/albums", handlers.Repo.Albums)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}