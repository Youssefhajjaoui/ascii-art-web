package main

import (
	"fmt"
	"net/http"

	Handle "web/handler"
)

func main() {
	http.HandleFunc("/", Handle.Toget)
	http.HandleFunc("/ascii-art", Handle.Topost)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
