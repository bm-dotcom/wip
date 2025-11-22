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

## Local Development
```bash
go mod tidy
go run api/handler.go
```
Open http://localhost:808
