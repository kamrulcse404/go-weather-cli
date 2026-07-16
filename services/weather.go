package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather/models"
)

const (
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
	apiKey  = "eb3a7249b3e996fc89bd942cc168daa0"
)

func GetWeather(city string) (models.WeatherResponse, error) {
	url := fmt.Sprintf("%s?q=%s&units=metric&appid=%s", baseURL, city, apiKey)
	res, err := http.Get(url)

	if err != nil {
		return models.WeatherResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.WeatherResponse{}, fmt.Errorf("failed to fetch weather: %s", res.Status)
	}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return models.WeatherResponse{}, err
	}

	var weather models.WeatherResponse
	err = json.Unmarshal(body, &weather)

	if err != nil {
		return models.WeatherResponse{}, err
	}

	return weather, nil
}
