package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WeatherClient represents a client for fetching weather data from an external API.
type WeatherClient struct {
	apiKey string
}

// NewWeatherClient creates a new instance of WeatherClient with the provided API key.
func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{apiKey: apiKey}
}

// WeatherData represents the structure for unmarshaling weather data from the external API response.
type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// API URL for fetching weather data from the OpenWeatherMap API.
const openWeatherMapURL = "http://api.openweathermap.org/data/2.5/weather"

// GetWeatherData retrieves weather data for a specific city using the OpenWeatherMap API.
func (c *WeatherClient) GetWeatherData(city string) (*WeatherData, error) {
	// Build the URL for the API request
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherMapURL, city, c.apiKey)

	// Perform the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into the WeatherData struct
	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}
	return &weatherData, nil
}
