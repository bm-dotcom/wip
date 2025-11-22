package handler

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var templateFS embed.FS

var router *gin.Engine

func init() {
	// Use gin.Default() so you get Logger + Recovery middleware for free (very useful in serverless)
	router = gin.Default()

	// Parse all templates once at cold start
	tmpl := template.Must(template.ParseFS(templateFS, "templates/*.html"))
	router.SetHTMLTemplate(tmpl)

	// Catch-all handler (thanks to vercel.json rewrite, every request hits this function)
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// Normalize path: remove leading/trailing slashes and handle empty
		cleanPath := strings.Trim(path, "/")
		if cleanPath == "" {
			cleanPath = "index"
		}

		tmplName := cleanPath + ".html"

		// Security: only allow the templates you actually have
		if tmplName != "index.html" && tmplName != "about.html" {
			c.String(http.StatusNotFound, "404 page not found")
			return
		}

		// If you later pass data to templates, use gin.H{"title": "About"} etc.
		c.HTML(http.StatusOK, tmplName, nil)
	})
}

// Vercel Go runtime expects this exact signature
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
