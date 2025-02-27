package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	GEO_API_URL      = "https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json"
	FORECAST_API_URL = "https://api.open-meteo.com/v1/forecast?latitude=%.6f&longitude=%.6f&hourly=temperature_2m"
)

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoResponse struct {
	Results []LatLong `json:"results"`
}

func getLatLong(city string) (*LatLong, error) {
	endpoint := fmt.Sprintf(GEO_API_URL, url.QueryEscape(city))
	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, fmt.Errorf("error making request to Geo API: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body")
		}
	}(resp.Body)

	var response GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(response.Results) < 1 {
		return nil, errors.New("no results found")
	}

	return &response.Results[0], nil
}

func getWeather(latLong LatLong) (string, error) {
	endpoint := fmt.Sprintf(FORECAST_API_URL, latLong.Latitude, latLong.Longitude)
	resp, err := http.Get(endpoint)

	if err != nil {
		return "", fmt.Errorf("error making request to Weather API: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func main() {
	latlong, err := getLatLong("Toronto, Canada")
	if err != nil {
		log.Fatalf("Failed to get latitude and longitude: %s", err)
	}
	fmt.Printf("Latitude: %f, Longitude: %f\n\n", latlong.Latitude, latlong.Longitude)

	weather, err := getWeather(*latlong)
	if err != nil {
		log.Fatalf("Failed to get weather: %s", err)
	}
	fmt.Printf("Weather: %s\n", weather)
}
