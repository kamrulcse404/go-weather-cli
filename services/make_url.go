package services

import (
	"fmt"
	"weather/models"
)

func makeUrl(config models.Config, city string) string {
	return fmt.Sprintf("%s?q=%s&units=metric&appid=%s", config.BaseURL, city, config.APIKey)
}
