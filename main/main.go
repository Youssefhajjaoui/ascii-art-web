package main

import (
	web "ascii/pkg/AsciiWeb"
	"log"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(web.Web_Proc),
		// ReadTimeout:  1000,
		// WriteTimeout: 1000,
	}

	log.Fatal(s.ListenAndServe())
}
