// api/index.go
package handler

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed ../templates/*
var templateFS embed.FS

var tmpl = template.Must(template.New("").ParseFS(templateFS, "templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "oops", 500)
		}
		return
	}
	http.NotFound(w, r)
}