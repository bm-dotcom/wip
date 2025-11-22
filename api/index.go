package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Period struct {
	Name          string `json:"name"`
	ShortForecast string `json:"shortForecast"`
	Temperature   string `json:"temperature"`
	WindSpeed     string `json:"windSpeed"`
}

type ForecastProperties struct {
	Periods []Period `json:"periods"`
}

type Forecast struct {
	Properties ForecastProperties `json:"properties"`
}

type PointProperties struct {
	Forecast string `json:"forecast"`
}

type Point struct {
	Properties PointProperties `json:"properties"`
}

type PageData struct {
	Title   string
	Content string
	Periods []Period
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func getForecast(ctx context.Context, lat, lon float64) ([]Period, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	pointsURL := fmt.Sprintf("https://api.weather.gov/points/%.4f,%.4f", lat, lon)
	req, err := http.NewRequestWithContext(ctx, "GET", pointsURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "GoWeatherApp (contact@example.com)")
	req.Header.Set("Accept", "application/geo+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("points API error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var point Point
	if err := json.Unmarshal(body, &point); err != nil {
		return nil, err
	}

	forecastURL := point.Properties.Forecast
	req, err = http.NewRequestWithContext(ctx, "GET", forecastURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "GoWeatherApp (contact@example.com)")
	req.Header.Set("Accept", "application/geo+json")

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("forecast API error: %d", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var forecast Forecast
	if err := json.Unmarshal(body, &forecast); err != nil {
		return nil, err
	}

	return forecast.Properties.Periods, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	lat := 40.7128 // NYC default
	lon := -74.0060

	if latStr := r.URL.Query().Get("lat"); latStr != "" {
		if f, err := strconv.ParseFloat(latStr, 64); err == nil {
			lat = f
		}
	}
	if lonStr := r.URL.Query().Get("lon"); lonStr != "" {
		if f, err := strconv.ParseFloat(lonStr, 64); err == nil {
			lon = f
		}
	}

	ctx := r.Context()
	periods, err := getForecast(ctx, lat, lon)
	if err != nil {
		fmt.Printf("Weather fetch error: %v\n", err)
	}

	data := PageData{
		Title:   "Weather Home",
		Content: "Real-time weather forecast from the National Weather Service API. Customize with ?lat=39.7456&lon=-97.0892 (US center).",
		Periods: periods,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "About",
		Content: "This is a simple Go web application deployed on Vercel, integrating the public NWS (National Weather Service) API for real-time forecasts.",
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "about.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/style.css"):
		http.ServeFile(w, r, "public/style.css")
	case r.URL.Path == "/":
		homeHandler(w, r)
	case r.URL.Path == "/about":
		aboutHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}
