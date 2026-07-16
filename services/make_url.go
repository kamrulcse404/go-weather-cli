package services

import "fmt"

func makeUrl(baseURL string, city string, apiKey string) string {
	return fmt.Sprintf("%s?q=%s&units=metric&appid=%s", baseURL, city, apiKey)
}
