package main

import (
	"fmt"
	"net/http"
)



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello World")

		if err != nil {
			fmt.Printf("Error is %s", err)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}