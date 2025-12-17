package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/KBrian85/weather-brief/internal/openmeteo"
)

type BriefPage struct {
	Location  string
	Today     openmeteo.DailyBrief
	TodayText string
}

func main() {
	indexTmpl := template.Must(template.ParseFiles(filepath.Join("web", "templates", "index.html")))
	briefTmpl := template.Must(template.ParseFiles(filepath.Join("web", "templates", "brief.html")))
	om := openmeteo.New()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		_ = indexTmpl.Execute(w, nil)
	})

	mux.HandleFunc("/brief", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad form submission", http.StatusBadRequest)
			return
		}

		city := strings.TrimSpace(r.FormValue("city"))
		if city == "" {
			http.Error(w, "City is required", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		loc, err := om.Geocode(ctx, city)
		if err != nil {
			http.Error(w, "City not found. Try another.", http.StatusBadRequest)
			return
		}

		forecast, err := om.DailyForecast(ctx, loc.Lat, loc.Lon, 1)
		if err != nil || len(forecast) == 0 {
			http.Error(w, "Could not load forecast. Try again.", http.StatusBadRequest)
			return
		}

		locationName := loc.Name
		if loc.Admin1 != "" {
			locationName = locationName + ", " + loc.Admin1
		}
		locationName = locationName + ", " + loc.Country

		data := BriefPage{
			Location:  locationName,
			Today:     forecast[0],
			TodayText: openmeteo.WeatherText(forecast[0].WeatherCode),
		}

		_ = briefTmpl.Execute(w, data)
	})

	log.Println("Listening on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", mux))
}