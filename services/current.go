package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"weather/config"
	"weather/models"
)

func GetWeather(cfg config.Config, city string) (models.WeatherResponse, error) {
	escapedCity := url.QueryEscape(city)
	requestURL := makeURL(cfg, escapedCity)

	res, err := http.Get(requestURL)

	if err != nil {
		return models.WeatherResponse{}, err
	}
	defer res.Body.Close()

	var apiErr models.ErrorResponse
	if res.StatusCode != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil {
			return models.WeatherResponse{}, err
		}

		return models.WeatherResponse{}, fmt.Errorf("weather API (%s): %s", apiErr.Cod, apiErr.Message)
	}

	var weather models.WeatherResponse

	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	return weather, nil
}
