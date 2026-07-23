package commands

import (
	"fmt"
	"os"
	"strings"
	"weather/config"
	"weather/services"
	"weather/utils"
)

func HandleCurrentWeather(args []string, cfg config.Config) {
	if len(os.Args) < 3 {
		utils.PrintUsage()
		return
	}

	city := strings.Join(os.Args[2:], " ")

	weather, err := services.GetWeather(cfg, city)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.PrintWeather(weather)
}
