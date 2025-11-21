// api/index.go
package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tmpl = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

func Handler(w http.ResponseWriter, r *http.Request) {
	// Accept root path in any form Vercel sends it
	if r.URL.Path == "/" || r.URL.Path == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err := tmpl.ExecuteTemplate(w, "layout.html", nil)
		if err != nil {
			http.Error(w, "Template error", 500)
			log.Printf("ExecuteTemplate error: %v", err)
		}
		return
	}

	// Everything else = 404
	http.NotFound(w, r)
}