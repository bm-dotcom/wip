// Go Module Configuration
// This go.mod file defines the dependencies for our simple webpage project.
// Only essential dependencies are included to keep the project lightweight.

module simple-webpage-go

// Go version requirement - using Go 1.25.4 for modern features and security
go 1.25.4

// Direct dependencies - explicitly added by the developer
require (
	// Fiber web framework v2 - fast HTTP router for Go
	// https://github.com/gofiber/fiber
	// Lightweight, fast web framework inspired by Express.js but for Go
	github.com/gofiber/fiber/v2 v2.52.10

	// HTML template engine for Fiber - handles template rendering
	// https://github.com/gofiber/template
	// Allows us to use Go templates for dynamic HTML generation
	github.com/gofiber/template/html/v2 v2.1.3
)

// Indirect dependencies - automatically added by Go when running 'go mod tidy'
// These are required by our direct dependencies but we don't import them directly
require (
	// Compression libraries for HTTP responses
	github.com/andybalholm/brotli v1.1.0 // indirect  // Brotli compression algorithm
	github.com/klauspost/compress v1.17.9 // indirect // General compression utilities

	// Fiber core utilities and template helpers
	github.com/gofiber/template v1.8.3 // indirect    // Base template engine
	github.com/gofiber/utils v1.1.0 // indirect       // Utility functions for Fiber

	// UUID generation - used by Fiber for request IDs
	github.com/google/uuid v1.6.0 // indirect  // Universally Unique IDentifier (UUID) generator

	// Terminal/console output utilities - used for server logs
	github.com/mattn/go-colorable v0.1.13 // indirect  // Cross-platform color terminal support
	github.com/mattn/go-isatty v0.0.20 // indirect      // Determines if output is a terminal
	github.com/mattn/go-runewidth v0.0.16 // indirect   // Unicode rune width calculation
	github.com/rivo/uniseg v0.2.0 // indirect           // Unicode text segmentation

	// Fast HTTP implementation - core of Fiber's performance
	github.com/valyala/bytebufferpool v1.0.0 // indirect // Efficient byte buffer pooling
	github.com/valyala/fasthttp v1.51.0 // indirect       // High-performance HTTP library
	github.com/valyala/tcplisten v1.0.0 // indirect       // Optimized TCP listener

	// System-level utilities
	golang.org/x/sys v0.28.0 // indirect  // System calls and os operations
)
