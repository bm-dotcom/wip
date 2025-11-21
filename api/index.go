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

// Parse all templates from the embedded FS (paths relative to project root)
var tmpl = template.Must(template.ParseFS(templateFS, "template/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	// Vercel sends "" or "/" for the root domain — accept both
	if r.URL.Path == "/" || r.URL.Path == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
			log.Printf("Template error: %v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
		return
	}

	// Everything else is 404
	http.NotFound(w, r)
}
