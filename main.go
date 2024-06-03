package main

import (
	"fmt"
	"net/http"
)

func toget(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method faile", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "index.html")
}

func topost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method faile", http.StatusMethodNotAllowed)
	}
	Out := r.PostFormValue("text")
	// t, err := template.ParseFiles("index.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = t.Execute(w, Out)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	http.ServeFile(w, r, "index.html")
	fmt.Fprintln(w, Out)
	// http.Redirect(w, r, "/", 8080)
}

func main() {
	http.HandleFunc("/", toget)
	http.HandleFunc("/ascii-art", topost)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
