package openmeteo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	HTTP *http.Client
}

func New() *Client {
	return &Client{HTTP: &http.Client{Timeout: 8 * time.Second}}
}

type Location struct {
	Name    string
	Admin1  string
	Country string
	Lat     float64
	Lon     float64
}

type geocodeResponse struct {
	Results []struct {
		Name    string  `json:"name"`
		Admin1  string  `json:"admin1"`
		Country string  `json:"country"`
		Lat     float64 `json:"latitude"`
		Lon     float64 `json:"longitude"`
	} `json:"results"`
}

func (c *Client) Geocode(ctx context.Context, city string) (Location, error) {
	if len(city) < 2 {
		return Location{}, errors.New("city too short")
	}

	u, _ := url.Parse("https://geocoding-api.open-meteo.com/v1/search")
	q := u.Query()
	q.Set("name", city)
	q.Set("count", "1")
	q.Set("language", "en")
	q.Set("format", "json")
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("geocoding returned %s", resp.Status)
	}

	var decoded geocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return Location{}, err
	}
	if len(decoded.Results) == 0 {
		return Location{}, errors.New("no matches")
	}

	r := decoded.Results[0]
	return Location{
		Name:    r.Name,
		Admin1:  r.Admin1,
		Country: r.Country,
		Lat:     r.Lat,
		Lon:     r.Lon,
	}, nil
}
type forecastResponse struct {
	Daily struct {
		Time             []string  `json:"time"`
		WeatherCode      []int     `json:"weather_code"`
		TempMax          []float64 `json:"temperature_2m_max"`
		TempMin          []float64 `json:"temperature_2m_min"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
	} `json:"daily"`
}

type DailyBrief struct {
	Date        string
	WeatherCode int
	TempMax     float64
	TempMin     float64
	PrecipMM    float64
}

func (c *Client) DailyForecast(ctx context.Context, lat, lon float64, days int) ([]DailyBrief, error) {
	u, _ := url.Parse("https://api.open-meteo.com/v1/forecast")
	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%.6f", lat))
	q.Set("longitude", fmt.Sprintf("%.6f", lon))

	// daily variables. When requesting daily, provide timezone. Using "auto" is simplest.
	q.Set("daily", "weather_code,temperature_2m_max,temperature_2m_min,precipitation_sum")
	q.Set("timezone", "auto")
	q.Set("forecast_days", fmt.Sprintf("%d", days))

	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("forecast returned %s", resp.Status)
	}

	var decoded forecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return nil, err
	}

	n := len(decoded.Daily.Time)
	if n == 0 {
		return nil, errors.New("no daily data")
	}

	out := make([]DailyBrief, 0, n)
	for i := 0; i < n; i++ {
		b := DailyBrief{Date: decoded.Daily.Time[i]}
		if i < len(decoded.Daily.WeatherCode) {
			b.WeatherCode = decoded.Daily.WeatherCode[i]
		}
		if i < len(decoded.Daily.TempMax) {
			b.TempMax = decoded.Daily.TempMax[i]
		}
		if i < len(decoded.Daily.TempMin) {
			b.TempMin = decoded.Daily.TempMin[i]
		}
		if i < len(decoded.Daily.PrecipitationSum) {
			b.PrecipMM = decoded.Daily.PrecipitationSum[i]
		}
		out = append(out, b)
	}
	return out, nil
}

func WeatherText(code int) string {
	switch code {
	case 0:
		return "Clear sky"
	case 1, 2, 3:
		return "Cloudy"
	case 45, 48:
		return "Fog"
	case 61, 63, 65:
		return "Rain"
	case 71, 73, 75:
		return "Snow"
	case 80, 81, 82:
		return "Showers"
	case 95:
		return "Thunderstorm"
	default:
		return "Mixed conditions"
	}
}
