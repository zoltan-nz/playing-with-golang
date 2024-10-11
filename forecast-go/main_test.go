package main

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLatLong(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"results":[{"latitude":51.5074,"longitude":-0.1278}]}`

	url := fmt.Sprintf(GEO_API_URL, "London")
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, mockResponse))

	latLong, err := getLatLong("London")
	assert.NoError(t, err)
	assert.NotNil(t, latLong)
	assert.Equal(t, 51.5074, latLong.Latitude)
	assert.Equal(t, -0.1278, latLong.Longitude)
}

func TestGetWeather(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"hourly":{"temperature_2m":[{"value":10.0}]}}`
	url := fmt.Sprintf(FORECAST_API_URL, 51.5074, -0.1278)
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, mockResponse))

	latLong := LatLong{Latitude: 51.5074, Longitude: -0.1278}
	temperature, err := getWeather(latLong)
	assert.NoError(t, err)
	assert.Equal(t, "{\"hourly\":{\"temperature_2m\":[{\"value\":10.0}]}}", temperature)
}
