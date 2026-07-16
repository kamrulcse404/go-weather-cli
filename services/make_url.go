package services

import (
	"fmt"
	"weather/config"
)

func makeURL(cfg config.Config, city string) string {
	return fmt.Sprintf("%s?q=%s&units=metric&appid=%s", cfg.BaseURL, city, cfg.APIKey)
}
