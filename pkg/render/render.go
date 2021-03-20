package render

import (
	"bytes"
	"fmt"
	"github.com/tklara86/lp_collection_go/pkg/config"
	"github.com/tklara86/lp_collection_go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates (a *config.AppConfig) {
	app = a
}

func AddDefaultData(td * models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, tmplData *models.TemplateData) {
	var tc map[string]*template.Template
	// Get the template cache from teh app config
	if app.UseCache {
		// Get the template cache from teh app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buff := new(bytes.Buffer)

	tmplData = AddDefaultData(tmplData)

	_ = t.Execute(buff, tmplData)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing to browser")
	}

}


// CreateTemplateCache creates a template cache as a mp
func CreateTemplateCache() (map[string]*template.Template ,error) {
	// creating map which will hold all templates
	templateCache := map[string]*template.Template{}
	// find all pages file in templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return templateCache,err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return templateCache,err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return templateCache,err
			}
		}

		templateCache[name] =  templateSet
	}
	return templateCache,nil
}