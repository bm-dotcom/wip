// Package main implements a production-hardened web server using Fiber v2.
// Handles CSRF protection, rate limiting, gzip compression, and view rendering.
// Serves a monitoring dashboard with responsive UI.

package main

// Imports: context for timeouts, fmt/slog/log for structured logging, os/signal for graceful shutdown,
// Fiber framework and middleware for web server functionality, html template engine for views.
import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

// main initializes the web server with middleware and routes, starts listening, and handles graceful shutdown.
func main() {
	// Initialize HTML template engine to load and render template files from ./templates directory.
	engine := html.New("./templates", ".html")

	// Create Fiber app instance with custom configuration for views and error handling.
	app := fiber.New(fiber.Config{
		Views: engine, // Attach template engine for rendering HTML views.
		// Custom error handler for unhandled errors, logs with structured data and returns 500.
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			slog.Error("request error", slog.String("error", err.Error()), slog.Int("status", code))
			return c.Status(code).SendString("Internal Server Error")
		},
	})

	// Middleware: general HTTP request logger formatting time, IP, status, latency, method, path.
	app.Use(logger.New(logger.Config{
		Format: "${time} [${ip}] ${status} ${latency} ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone: "Local",
	}))

	// Middleware: rate limiter allowing 10 requests per second per IP to prevent abuse.
	app.Use(limiter.New(limiter.Config{
		Max:        10, // Max requests per IP per expiration window.
		Expiration: 1 * time.Second, // Reset window.
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Use client IP as key.
		},
		LimitReached: func(c *fiber.Ctx) error {
			slog.Warn("rate limit exceeded", slog.String("ip", c.IP()))
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	}))

	// Middleware: security headers including strict Content Security Policy for CDNs.
	app.Use(helmet.New(helmet.Config{
		ContentSecurityPolicy: "default-src 'none'; script-src 'self' https://cdn.tailwindcss.com https://unpkg.com https://cdn.jsdelivr.net https://cdnjs.cloudflare.com 'unsafe-inline' 'unsafe-eval'; style-src 'self' https://cdn.tailwindcss.com https://unpkg.com https://*.cartocdn.com https://*.basemaps.cartocdn.com 'unsafe-inline'; img-src data: https: blob: https://*.cartocdn.com https://*.basemaps.cartocdn.com; connect-src 'self' https: ws: wss: data: blob:; object-src 'none'; base-uri 'self'; form-action 'self'; frame-ancestors 'none';",
	}))

	// Middleware: gzip compression for responses to improve bandwidth.
	app.Use(compress.New())

	// Middleware: enables Cross-Origin Resource Sharing for required domains if needed.
	app.Use(cors.New())

	// Route: GET / renders the index template with layout for the monitoring dashboard.
	app.Get("/", func(c *fiber.Ctx) error {
		// Generate unique request ID and measure processing time for logging.
		reqID := fmt.Sprintf("%d", time.Now().UnixNano())
		start := time.Now()
		defer func() {
			dur := time.Since(start)
			// Log structured request details after processing.
			slog.Info("request",
				slog.String("req_id", reqID),
				slog.String("method", c.Method()),
				slog.String("path", c.Path()),
				slog.Int("status", c.Response().StatusCode()),
				slog.Duration("duration", dur),
				slog.String("ip", c.IP()),
			)
		}()

		// Render index.html template using layout.html.
		return c.Render("index", fiber.Map{}, "layout")
	})

	// Route: GET /dashboard renders the dashboard template with layout.
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		// Generate unique request ID and measure processing time for logging.
		reqID := fmt.Sprintf("%d", time.Now().UnixNano())
		start := time.Now()
		defer func() {
			dur := time.Since(start)
			// Log structured request details after processing.
			slog.Info("request",
				slog.String("req_id", reqID),
				slog.String("method", c.Method()),
				slog.String("path", c.Path()),
				slog.Int("status", c.Response().StatusCode()),
				slog.Duration("duration", dur),
				slog.String("ip", c.IP()),
			)
		}()

		// Render dashboard.html template using layout.html.
		return c.Render("dashboard", fiber.Map{}, "layout")
	})

	// Determine server port from environment or default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("server starting", slog.String("addr", ":"+port))

	// Start server in background to allow signal handling.
	go func() {
		if err := app.Listen(":" + port); err != nil {
			slog.Error("server error", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	// Handle shutdown signals (Ctrl+C, SIGTERM) for graceful termination.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // Block until signal received.
	slog.Info("shutting down gracefully")

	// Shutdown with timeout to allow connections to close gracefully.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("shutdown error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
