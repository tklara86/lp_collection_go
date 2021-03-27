package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tklara86/lp_collection_go/pkg/config"
	"github.com/tklara86/lp_collection_go/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/albums", handlers.Repo.Albums)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

