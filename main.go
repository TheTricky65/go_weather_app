package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const apiKey = "my_api_key"

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city name>")
		return
	}

	city := strings.Join(os.Args[1:], " ")
	getWeather(city)
}

func getWeather(city string) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("API Error: %d - %s\n", resp.StatusCode, resp.Status)
		return
	}

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	fmt.Printf("\nWeather in %s:\n", data.Name)
	fmt.Printf("Temperature: %.1f°C\n", data.Main.Temp)
	fmt.Printf("Condition: %s\n", strings.Title(data.Weather[0].Description))
	fmt.Printf("Humidity: %d%%\n", data.Main.Humidity)
	fmt.Printf("Wind Speed: %.1f km/h\n", data.Wind.Speed)
}
