/*
SIMPLE GO WEBPAGE APPLICATION
This is a minimal web server built with Go and the Fiber framework.
It serves a simple webpage with light/dark mode toggle functionality.

Architecture:
- Uses Fiber v2 as the web framework (fast HTTP router)
- HTML templates for dynamic content rendering
- Single homepage with dynamic server time
- No database - everything is static/stateless
- Pure CSS styling, no frameworks
*/

package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// main() function - entry point of our Go application
// This function sets up the web server and defines all the routes
func main() {

	/*
		TEMPLATE ENGINE SETUP
		Fiber uses HTML templates to render dynamic content. The template engine:
		- Looks for template files in ./templates/ directory
		- Uses .html file extension
		- Supports Go template syntax with {{.VariableName}} notation
	*/
	engine := html.New("./templates", ".html")

	/*
		FIBER APPLICATION SETUP
		Creates a new Fiber web application instance with configuration.
		Fiber is similar to Express.js but written in Go for high performance.
	*/
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	/*
		ROUTE DEFINITIONS
		Here we define all the URL routes that our web server responds to.
		Each route connects a URL path to a handler function.
	*/

	// HOMEPAGE ROUTE: GET request to "/" (root URL)
	// This is the main page visitors see when they visit the website
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "My Go Webpage",
			"Message": "Welcome to my simple Go webpage!",
			"Time":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // local dev
	}
	log.Fatal(app.Listen(":" + port)) // ← change from hardcoded :3000

	/*
		SERVER STARTUP
		Start the web server and begin accepting HTTP requests.
		The server will run indefinitely until stopped with Ctrl+C.
	*/

	log.Printf("🚀 Server starting on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
