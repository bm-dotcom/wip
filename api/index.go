// api/index.go
package handler

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed ../templates/*
var templateFS embed.FS

var tmpl = template.Must(template.New("").ParseFS(templateFS, "templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	// Accept both "" and "/" – Vercel sometimes sends empty string
	if r.URL.Path == "/" || r.URL.Path == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
			http.Error(w, "template error", 500)
			return
		}
		return
	}
	http.NotFound(w, r)
}