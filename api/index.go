package handler

import (
	"html/template"
	"net/http"
	"time"
)

type PageData struct {
	Title   string
	Message string
	Time    string
	Content string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	now := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")

	path := r.URL.Path

	switch {
	case path == "/about":
		data := PageData{
			Title:   "About",
			Content: "This is a simple webpage created with Go standard library and HTML templates, deployed as serverless function on Vercel.",
		}
		if err := tmpl.ExecuteTemplate(w, "about.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case path == "/" || path == "":
		data := PageData{
			Title:   "Go Webpage",
			Message: "Welcome from the Go serverless handler on Vercel!",
			Time:    now,
		}
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.NotFound(w, r)
	}
}
