package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr: ":8080",
	}

	http.Handle("/foo", nil)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(s.ListenAndServe())

	// fmt.Println("project valide inchallah with team khaliha 3ala allah")
}
