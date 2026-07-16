package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather/models"
)

func main() {
	url := "https://api.openweathermap.org/data/2.5/weather?q=Dhaka&units=metric&appid=eb3a7249b3e996fc89bd942cc168daa0"

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Something error: %v\n", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed. Status: %d\n", res.StatusCode)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Something error: %v\n", err)
		return
	}

	var weather models.WeatherResponse

	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("json convert problem: %v\n", err)
		return
	}

	fmt.Printf("%-15s %s\n", "Name:", weather.Name)
	fmt.Printf("%-15s %.1f °C\n", "Temperature:", weather.Main.Temp)
	fmt.Printf("%-15s %d %%\n", "Humidity:", weather.Main.Humidity)
	fmt.Printf("%-15s %.1f m/s\n", "Wind Speed:", weather.Wind.Speed)
}
