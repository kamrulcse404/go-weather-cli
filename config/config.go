package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", "", err
	}

	baseURL  := os.Getenv("BASE_URL")
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	return baseURL, apiKey, nil
}
