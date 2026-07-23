package services

import (
	"fmt"
	"weather/config"
)

func makeURL(cfg config.Config, endpoint, city string) string {
	return fmt.Sprintf("%s%s?q=%s&units=metric&appid=%s",
		cfg.BaseURL,
		endpoint,
		city,
		cfg.APIKey,
	)
}
