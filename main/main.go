package main

import (
	"log"
	"net/http"

	web "ascii/pkg/AsciiWeb"
)

func main() {
	http.HandleFunc("/", web.Web_Proc)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
