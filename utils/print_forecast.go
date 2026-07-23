package utils

import (
	"fmt"
	"strings"
	"time"
	"weather/models"
)

func PrintForecast(forecasts models.ForecastResponse) {
	fmt.Printf("Forecast for %s\n\n", forecasts.City.Name)
	currentDate := ""

	for _, item := range forecasts.List {
		timeAndDate := strings.Split(item.DateTime, " ")
		if len(timeAndDate) != 2 {
			continue
		}

		sDate := timeAndDate[0]
		sTime := timeAndDate[1][:5]

		if sDate != currentDate {
			currentDate = sDate
			t, err := time.Parse("2006-01-02", sDate)

			if err != nil {
				continue
			}

			fmt.Println()
			fmt.Println(t.Format("Mon, Jan 02"))
			fmt.Println()

			fmt.Printf("%-6s %-8s %-10s %-12s %-20s\n", "Time", "Temp", "Humidity", "Wind Speed", "Weather")
			fmt.Println(strings.Repeat("-", 60))
		}

		description := "N/A"

		if len(item.Weather) > 0 {
			description = item.Weather[0].Description
		}

		temp := fmt.Sprintf("%.0f°C", item.Main.Temp)
		wind := fmt.Sprintf("%.1f m/s", item.Wind.Speed)

		fmt.Printf(
			"%-6s %-8s %-10d %-12s %-20s\n",
			sTime,
			temp,
			item.Main.Humidity,
			wind,
			description,
		)
	}

	fmt.Println()
}
