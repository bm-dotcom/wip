package handler

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

//go:embed ../templates/*
var templateFS embed.FS

// Parse templates once at cold start (bulletproof with fs.Sub)
var templates *template.Template

func init() {
	subFS, err := fs.Sub(templateFS, "templates")
	if err != nil {
		panic(err)
	}
	templates = template.Must(template.ParseFS(subFS, "*.html"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Get clean path (e.g. "", "/", "/about", "/about/")
	p := strings.TrimSuffix(strings.Trim(r.URL.Path, "/"), "/")

	// Map path → template name
	var tmplName string
	switch p {
	case "", "index":
		tmplName = "index.html"
	case "about":
		tmplName = "about.html"
	default:
		http.NotFound(w, r)
		return
	}

	// Serve the template (your styling will be fully visible)
	err := templates.ExecuteTemplate(w, tmplName, map[string]any{
		"Title": strings.Title(p == "" ? "home" : p),
	})
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}