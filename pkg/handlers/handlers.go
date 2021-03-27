package handlers

import (
	"github.com/tklara86/lp_collection_go/pkg/config"
	"github.com/tklara86/lp_collection_go/pkg/models"
	"github.com/tklara86/lp_collection_go/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request)  {
	stringMap := make(map[string]string)
	stringMap["title"] = "Home Page"

	render.RenderTemplate(w,"home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Albums page handler
func (m *Repository) Albums(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Albums Page"
	render.RenderTemplate(w, "albums.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
