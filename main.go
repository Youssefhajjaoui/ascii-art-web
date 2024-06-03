package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func toget(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method faile", http.StatusMethodNotAllowed)
	}
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
	if r.Method != http.MethodPost {
		http.Error(w, "method faile", http.StatusMethodNotAllowed)
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Out := r.PostFormValue("text")
	fmt.Println(Out)
	err = t.Execute(w, Out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// http.ServeFile(w, r, "index.html")
	// http.Redirect(w, r, "/", 8080)
}

func main() {
	http.HandleFunc("/", toget)
	http.HandleFunc("/ascii-art", topost)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
