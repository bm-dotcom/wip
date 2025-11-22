# Global Monitoring Site (Phase 1 - UI Shell)

Minimal production-grade single-page global monitoring site shell.

## Features
- Full-viewport responsive Leaflet map (world center or geolocation)
- Dark/light mode toggle (localStorage, Tailwind `dark`)
- CartoDB Voyager (light) / Dark Matter (dark) tiles
- CDN assets (Tailwind, htmx, Alpine, Leaflet) with SRI hashes (latest Nov 2025: htmx 2.0.8, alpine 3.15.2, leaflet 1.9.4)
- Security headers (CSP strict for CDNs, X-Content-Type-Options, Referrer-Policy, Permissions-Policy geolocation)
- In-memory IP rate limiting (sync.Map, 10 req/sec)
- Structured JSON logging (req ID UUID, method, path, status, duration, IP)
- Pure net/http + html/template embed (no frameworks)
- Vercel serverless ready (single handler, vercel.json rewrite)

Ready for phase 2 (htmx SSE/WebSocket markers).

## Code Explanations

<!-- Explanation of functionality in files -->

### main.go
- main func initializes Fiber app with HTML templates.
- Middleware: logger for request logging, limiter (10 req/sec/IP), helmet (CSP for CDNs), compress (gzip), cors.
- Route GET / renders index template with layout.
- Logs structured info per request with reqID, method, path, status, duration, IP.
- Graceful shutdown on SIGTERM with 30s timeout.

### go.mod
- Module monitorpro with Go 1.23.
- Requires Fiber v2 for web framework, html template for views.

### templates/layout.html
- HTML base template with CDN links for Tailwind, Leaflet, Alpine.
- Navbar with logo, dark mode toggle/sun-moon icon, mobile menu.
- Main content section for nested templates.
- Footer with copyright and GitHub link.
- Alpine for dark mode state management with localStorage.

### templates/index.html
- Defines content block: hero title, description, map/feature grid.
- Leaflet map init with CartoDB tiles, geolocation, marker.
- Alpine for dark/light toggle.

## Local Development
```bash
go mod tidy
go run main.go
```
Open http://localhost:8080
