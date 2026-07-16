package main

import (
	"fmt"
	"weather/config"
	"weather/services"
	"weather/utils"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	weather, err := services.GetWeather(cfg, "dhaka")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	utils.PrintWeather(weather)
}
