package main

import (
	"fmt"
	"html/template"
	"net/http"
)


const portNumber = ":8080"

// Home is home page handler
func Home(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w,"home.page.tmpl")
}

// Albums page handler
func Albums(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "albums.page.tmpl")
}

// renderTemplate render templates handler
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Printf("error %s", err)
	}

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/albums", Albums)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}