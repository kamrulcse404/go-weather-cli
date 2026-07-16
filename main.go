package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.openweathermap.org/data/2.5/weather?q=Dhaka&units=metric&appid=eb3a7249b3e996fc89bd942cc168daa0"

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Something error: %v\n", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed. Status: %d\n", res.StatusCode)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Something error: %v\n", err)
		return
	}

}
