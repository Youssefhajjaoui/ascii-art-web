package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func toget(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "PATH INALOWED", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the index.html template
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an empty struct to hold form data (optional for cleaner separation)
	// Execute the template without initial data if it's a GET request
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func topost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		handleerr(w, r, http.StatusNotFound, "path incorect", http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodPost {
		handleerr(w, r, http.StatusNotFound, "path incorect", http.StatusText(http.StatusNotFound))
		return
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	text := r.PostFormValue("text")
	banner := r.PostFormValue("banner") // Adjust field name as needed
	editedText := strings.ToUpper(text)
	type FormData struct {
		Text   string
		Banner string // Adjust field name as needed
	}
	formData := FormData{
		Text:   editedText,
		Banner: banner,
	}
	fmt.Println("Received text:", banner)
	// Execute the template with the edited text
	err = t.Execute(w, formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", toget)
	http.HandleFunc("/ascii-art", topost)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func handleerr(w http.ResponseWriter, r *http.Request, statuscode int, message string, statustext string) {
	t, err := template.ParseFiles("warningmsg.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type errors struct {
		Text string
		code int // Adjust field name as needed
	}
	error := errors{
		Text: statustext,
		code: statuscode,
	}
	err = t.Execute(w, error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, message, http.StatusInternalServerError)
}
