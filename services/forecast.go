package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"weather/config"
	"weather/models"
)

func GetForecast(cfg config.Config, city string) (models.ForecastResponse, error) {
	escapedCity := url.QueryEscape(city)
	requestURL := makeURL(cfg, "forecast", escapedCity)

	res, err := http.Get(requestURL)

	if err != nil {
		return models.ForecastResponse{}, err
	}
	defer res.Body.Close()

	var apiErr models.ErrorResponse
	if res.StatusCode != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil {
			return models.ForecastResponse{}, err
		}

		return models.ForecastResponse{}, fmt.Errorf("weather API (%s): %s", apiErr.Cod, apiErr.Message)
	}

	var forecast models.ForecastResponse

	err = json.NewDecoder(res.Body).Decode(&forecast)
	if err != nil {
		return models.ForecastResponse{}, err
	}

	return forecast, nil
}
