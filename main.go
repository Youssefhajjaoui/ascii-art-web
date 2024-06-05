package main

import (
	"fmt"
	"html/template"
	"net/http"

	process "web/Ascii"
	"web/Outil"
)

func toget(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleerr(w, http.StatusNotFound, "path incorect", http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodGet {
		handleerr(w, http.StatusMethodNotAllowed, "path incorect", http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	// Parse the index.html template
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func topost(w http.ResponseWriter, r *http.Request) {
	myoutils := new(Outil.Outils)
	if r.URL.Path != "/ascii-art" {
		handleerr(w, http.StatusNotFound, "path incorect", http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodPost {
		handleerr(w, http.StatusMethodNotAllowed, "path incorect", http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	text := r.PostFormValue("text")
	banner := r.PostFormValue("banner")
	myoutils.WordToChange = text
	myoutils.Banner = banner
	if text == "" {
		handleerr(w, http.StatusBadRequest, "string erreur", http.StatusText(http.StatusBadRequest))
		return
	}
	process.Process_Text(myoutils)
	if myoutils.Erreur != "" {
		handleerr(w, http.StatusBadRequest, "string erreur", http.StatusText(http.StatusBadRequest))
		return
	}
	type FormData struct {
		Text   string
		Banner string // Adjust field name as needed
	}
	formData := FormData{
		Text:   myoutils.Result,
		Banner: myoutils.Banner,
	}
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

func handleerr(w http.ResponseWriter, statuscode int, message string, statustext string) {
	t, err := template.ParseFiles("warningmsg.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type errors struct {
		Statustext string
		Statuscode int // Corrected field name
	}
	error := errors{
		Statustext: statustext,
		Statuscode: statuscode, // Corrected field name
	}
	w.WriteHeader(statuscode)
	err = t.Execute(w, error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, message, http.StatusInternalServerError)
}
