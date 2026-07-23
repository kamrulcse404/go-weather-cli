package utils

import (
	"fmt"
	"weather/models"
)

func PrintCurrentWeather(weather models.WeatherResponse) {
	fmt.Printf("%-15s %s\n", "Name:", weather.Name)
	fmt.Printf("%-15s %.1f °C\n", "Temperature:", weather.Main.Temp)
	fmt.Printf("%-15s %d %%\n", "Humidity:", weather.Main.Humidity)
	fmt.Printf("%-15s %.1f m/s\n", "Wind Speed:", weather.Wind.Speed)

	if len(weather.Weather) > 0 {
		fmt.Printf("%-15s %s\n", "Condition:", weather.Weather[0].Main)
		fmt.Printf("%-15s %s\n", "Description:", weather.Weather[0].Description)
	}
}
