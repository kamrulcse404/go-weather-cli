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
	if len(os.Args) < 2 {
		utils.PrintUsage()
		return
	}

	command := strings.ToLower(os.Args[1])

	if command == "--help" || command == "-h" {
		utils.PrintUsage()
		return
	}

	switch command {
	case "current":
		if len(os.Args) < 3 {
			utils.PrintUsage()
			return
		}

		cfg, err := config.LoadEnv()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		city := strings.Join(os.Args[2:], " ")

		weather, err := services.GetWeather(cfg, city)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		utils.PrintWeather(weather)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		utils.PrintUsage()
	}

}
