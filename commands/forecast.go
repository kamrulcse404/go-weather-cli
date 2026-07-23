package commands

import (
	"fmt"
	"strings"
	"weather/config"
	"weather/services"
	"weather/utils"
)

func HandleForecast(args []string, cfg config.Config) {
	if len(args) < 3 {
		utils.PrintUsage()
		return
	}

	city := strings.Join(args[2:], " ")

	forecast, err := services.GetForecast(cfg, city)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.PrintForecast(forecast)
}
