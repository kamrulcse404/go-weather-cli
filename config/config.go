package config

import (
	"os"
	"weather/models"

	"github.com/joho/godotenv"
)

func LoadEnv() (models.Config, error) {
	err := godotenv.Load()

	if err != nil {
		return models.Config{}, err
	}

	baseURL  := os.Getenv("BASE_URL")
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	var config models.Config
	config.BaseURL = baseURL
	config.APIKey = apiKey
	
	return config, nil
}
