package utils

import "fmt"

func PrintUsage() {
	fmt.Println()
	fmt.Println("Weather CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  weather <city>")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  weather Dhaka")
	fmt.Println("Options:")
	fmt.Println("  -h, --help    Show this help message")
}