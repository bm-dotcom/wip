package handler

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed ../templates/*
var templateFS embed.FS

var engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	// Uncomment the line below temporarily if you want to see request logs in Vercel
	// router.Use(gin.Logger())

	// This fs.Sub trick makes the embed path bulletproof – works every single time on Vercel
	templates, err := fs.Sub(templateFS, "templates")
	if err != nil {
		panic(fmt.Sprintf("template sub FS error: %v", err))
	}

	tmpl := template.Must(template.ParseFS(templates, "*.html"))
	router.SetHTMLTemplate(tmpl)

	router.NoRoute(func(c *gin.Context) {
		path := strings.Trim(c.Request.URL.Path, "/")

		if path == "" {
			path = "index"
		}

		templateName := path + ".html"

		// Whitelist only existing templates
		if templateName != "index.html" && templateName != "about.html" {
			c.String(http.StatusNotFound, "404 - Page not found")
			return
		}

		title := "Home"
		if path != "" {
			title = strings.Title(path)
		}

		c.HTML(http.StatusOK, templateName, gin.H{
			"Title": title,
		})
	})

	engine = router
}

func Handler(w http.ResponseWriter, r *http.Request) {
	engine.ServeHTTP(w, r)
}
