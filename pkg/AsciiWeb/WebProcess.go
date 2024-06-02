package asciiweb

import (
	"fmt"
	"net/http"
	"text/template"
)

func Web_Proc(w http.ResponseWriter, r *http.Request) {
	// MainWeb := ""
	ErrWeb := "Error/index-err.html"
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello ROOT")
	} else {
		template.ParseFiles(ErrWeb)
	}
}
