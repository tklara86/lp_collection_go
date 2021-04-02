package handlers

import (
	"encoding/json"
	"github.com/tklara86/lp_collection_go/internal/config"
	"github.com/tklara86/lp_collection_go/internal/models"
	"github.com/tklara86/lp_collection_go/internal/render"
	"log"
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

	render.RenderTemplate(w,r,"home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}


// Albums page handler
func (m *Repository) Albums(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["title"] = "Albums Page"
	render.RenderTemplate(w,r, "albums.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}



//type jsonResponse struct {
//	Name string `json:"name"`
//	Email string `json:"email"`
//	Description  string `json:"description"`
//}

// PostAlbum page handler
//func (m *Repository) PostAlbum(w http.ResponseWriter, r *http.Request) {
//	name := r.Form.Get("name")
//	email := r.Form.Get("email")
//	description := r.Form.Get("description")
//
//	response := jsonResponse{
//		Name: name,
//		Email: email,
//		Description: description,
//	}
//
//	result, err := json.MarshalIndent(response, "", "    ")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(string(result))
//	w.Header().Set("Content-Type", "application/json")
//	_, err = w.Write(result)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	//_, err := w.Write([]byte(fmt.Sprintf("Name is %s, email is %s and description is %s", name, email,description)))
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//}


type jsonResponse struct {
	Name string `json:"name"`
	Description  string `json:"description"`
}

//SearchAlbums page handler
func (m *Repository) SearchAlbums(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Name: "Bitches Brew",
		Description: "Miles Davis' Album",
	}

	out, err := json.MarshalIndent(resp, "", "     ")

	if err != nil {
		log.Println(err)
	}
	// Set the header
	w.Header().Set("Content-Type", "application/json")
	_ ,err = w.Write(out)

	if err != nil {
		log.Fatal(err)
	}
}