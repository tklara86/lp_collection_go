package main


import (
	"fmt"
	"github.com/tklara86/lp_collection_go/pkg/handlers"
	"net/http"
)


const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/albums", handlers.Albums)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}