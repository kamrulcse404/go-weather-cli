package utils

import "fmt"

func PrintUsage() {
	fmt.Println()
	fmt.Println("Weather CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  weather current <city>")
	fmt.Println("  weather forecast <city>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  weather current Dhaka")
	fmt.Println("  weather current \"New York\"")
	fmt.Println("  weather forecast Dhaka")
	fmt.Println("  weather forecast \"New York\"")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help      Show this help message")
	fmt.Println()
}
