package main

import (
	"fmt"
	"os"
	"strings"
	"weather/commands"
	"weather/config"
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
		cfg, err := config.LoadEnv()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		commands.HandleCurrentWeather(os.Args, cfg)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		utils.PrintUsage()
	}

}
