package main

import (
	"fmt"
	"os"
	"strings"
	"weather/config"
	"weather/services"
	"weather/utils"
)

func main() {
	if len(os.Args) < 2{
		utils.PrintUsage()
		return
	}

	city := strings.Join(os.Args[1:], " ")

	cfg, err := config.LoadEnv()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	weather, err := services.GetWeather(cfg, city)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	utils.PrintWeather(weather)
}
