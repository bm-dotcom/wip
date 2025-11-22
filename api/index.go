package handler

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed ../templates/*
var templateFS embed.FS

var templates *template.Template

func init() {
	subFS, err := fs.Sub(templateFS, "templates")
	if err != nil {
		panic(err)
	}
	templates = template.Must(template.ParseFS(subFS, "*.html"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Clean the path
	p := strings.Trim(r.URL.Path, "/")
	if p == "" {
		p = "index"
	}

	var tmplName string
	switch p {
	case "index", "":
		tmplName = "index.html"
	case "about":
		tmplName = "about.html"
	default:
		http.NotFound(w, r)
		return
	}

	// Fix for the syntax error: no ternary in Go!
	title := "Home"
	if p != "" && p != "index" {
		title = strings.Title(p)
	}

	err := templates.ExecuteTemplate(w, tmplName, map[string]string{
		"Title": title,
	})
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}