package commands

import (
	"fmt"
	"strings"
	"weather/config"
	"weather/services"
	"weather/utils"
)

func HandleCurrentWeather(args []string, cfg config.Config) {
	if len(args) < 3 {
		utils.PrintUsage()
		return
	}

	city := strings.Join(args[2:], " ")

	weather, err := services.GetCurrentWeather(cfg, city)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.PrintWeather(weather)
}
