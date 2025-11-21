package handler

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := tmpl.ExecuteTemplate(w, "layout.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
