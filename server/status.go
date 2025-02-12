package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, statusCode int) {
	templateFile := fmt.Sprintf("static/status/%d.html", statusCode)
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}
	w.WriteHeader(statusCode)
	tmpl.Execute(w, nil)
}
