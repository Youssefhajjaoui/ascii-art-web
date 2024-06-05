package web

import (
	"html/template"
	"net/http"
)

func Handleerr(w http.ResponseWriter, statuscode int, message string, statustext string) {
	t, err := template.ParseFiles("templates/warningmsg.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type errors struct {
		Statustext string
		Statuscode int
		Message    string // Corrected field name
	}
	error := errors{
		Statustext: statustext,
		Statuscode: statuscode,
		Message:    message,
	}
	w.WriteHeader(statuscode)
	err = t.Execute(w, error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// http.Error(w, message, http.StatusInternalServerError)
}
