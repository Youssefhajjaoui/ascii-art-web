package web

import (
	"html/template"
	"net/http"
)

func Toget(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Handleerr(w, http.StatusNotFound, "path incorect", http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodGet {
		Handleerr(w, http.StatusMethodNotAllowed, "Method not allowed", http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	// Parse the index.html template
	t, err := template.ParseFiles("templates/index.html")
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
