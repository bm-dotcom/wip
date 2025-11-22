package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Title   string
	Message string
	Time    string
	Content string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")
	router.Use(gin.Recovery())

	now := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")

	router.GET("/", func(c *gin.Context) {
		data := PageData{
			Title:   "Go Webpage",
			Message: "Welcome from the Gin-powered Go serverless handler on Vercel!",
			Time:    now,
		}
		c.HTML(http.StatusOK, "index.html", data)
	})

	router.GET("/about", func(c *gin.Context) {
		data := PageData{
			Title:   "About",
			Content: "This is a simple webpage created with Go, Gin framework, and HTML templates, deployed as serverless function on Vercel.",
		}
		c.HTML(http.StatusOK, "about.html", data)
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 - Page Not Found")
	})

	router.ServeHTTP(w, r)
}
