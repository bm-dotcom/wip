package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl = template.Must(template.ParseGlob("../templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	// Only allow exact root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Optional: cache template parse in dev (Vercel cold starts anyway)
	if os.Getenv("VERCEL") == "1" {
		// Re-parse on cold start is fine — Vercel does it once per instance
		t, err := template.ParseGlob("../templates/*.html")
		if err != nil {
			http.Error(w, "Template parse error", 500)
			log.Printf("Template parse error: %v", err)
			return
		}
		tmpl = template.Must(t, nil)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Printf("Template execution failed: %v", err)
	}
}