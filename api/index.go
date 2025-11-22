package main

import (
	"embed"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/template/html/v2"
)

//go:embed ../templates/*
var templatesFS embed.FS

var app *fiber.App
var handler http.Handler // global adapted handler (built once)

func init() {
	engine := html.NewFileSystem(http.FS(templatesFS), ".html")

	app = fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "My Go Webpage",
			"Message": "Welcome to my simple Go webpage!",
			"Time":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "About",
		})
	})

	handler = adaptor.FiberApp(app) // built once, zero overhead
}

func main() {
	// local dev only
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}

// Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	// Fiber needs RequestURI for correct routing on Vercel
	if r.RequestURI == "" {
		r.RequestURI = r.URL.String()
	}
	handler.ServeHTTP(w, r)
}
