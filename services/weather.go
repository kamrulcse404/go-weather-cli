package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather/config"
	"weather/models"
)


func GetWeather(city string) (models.WeatherResponse, error) {
	config, err := config.LoadEnv()

	if err != nil {
		return models.WeatherResponse{}, err 
	}

	url := makeUrl(config, city)

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
