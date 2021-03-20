package handlers

import (
	"github.com/tklara86/lp_collection_go/pkg/render"
	"net/http"
)

// Home is home page handler
func Home(w http.ResponseWriter, r *http.Request)  {
	render.RenderTemplate(w,"home.page.tmpl")
}

// Albums page handler
func Albums(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "albums.page.tmpl")
}
