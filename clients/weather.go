package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherClient struct {
	apiKey string
}

func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{apiKey: apiKey}
}

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func (с *WeatherClient) GetWeatherData(city string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, с.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
