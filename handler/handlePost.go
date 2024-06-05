package web

import (
	"html/template"
	"net/http"

	process "web/Ascii"
	Outil "web/Outil"
)

func Topost(w http.ResponseWriter, r *http.Request) {
	myoutils := new(Outil.Outils)
	if r.URL.Path != "/ascii-art" {
		Handleerr(w, http.StatusNotFound, "path incorect", http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodPost {
		Handleerr(w, http.StatusMethodNotAllowed, "Method not allowed", http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	text := r.PostFormValue("text")
	banner := r.PostFormValue("banner")
	myoutils.WordToChange = text
	myoutils.Banner = banner
	if text == "" || len(text) > 2500 {
		Handleerr(w, http.StatusBadRequest, "string erreur", http.StatusText(http.StatusBadRequest))
		return
	}
	process.Process_Text(myoutils)
	if myoutils.Erreur != "" {
		Handleerr(w, http.StatusBadRequest, "string erreur", http.StatusText(http.StatusBadRequest))
		return
	}
	type FormData struct {
		Text   string
		Banner string // Adjust field name as needed
	}
	formData := FormData{
		Text: myoutils.Result,
	}
	err = t.Execute(w, formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
