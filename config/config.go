package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL string
	APIKey  string
}


func LoadEnv() (Config, error) {
	err := godotenv.Load()

	if err != nil {
		return Config{}, err
	}

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	return Config{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}, nil
}
