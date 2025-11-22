package handler

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Embed all templates at build time (no filesystem access needed at runtime)
//
//go:embed templates/*
var templates embed.FS

var tmpl *template.Template

// Parse once at cold start – cheap and safe in serverless
func init() {
	var err error
	tmpl, err = template.New("").ParseFS(templates, "templates/*.html")
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/index.html"
	} else {
		path = "/" + strings.TrimPrefix(path, "/") + ".html"
		if path == "/about.html" { // only allow the pages you actually have
			// ok
		} else if path != "/index.html" {
			http.NotFound(w, r)
			return
		}
	}

	// If you use a layout/base template, change the name below accordingly
	err := tmpl.ExecuteTemplate(w, strings.TrimPrefix(path, "/"), nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		fmt.Println("template error:", err)
	}
}
