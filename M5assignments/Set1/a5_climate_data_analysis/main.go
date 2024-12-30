package main

import (
	"fmt"
	"strings"
)

type CityData struct {
	CityName       string
	AvgTemperature float64
	Rainfall       float64
}

func highestTemperature(cities []CityData) (string, float64) {
	var highestCity string
	var highestTemp float64
	for _, city := range cities {
		if city.AvgTemperature > highestTemp {
			highestTemp = city.AvgTemperature
			highestCity = city.CityName
		}
	}
	return highestCity, highestTemp
}

func lowestTemperature(cities []CityData) (string, float64) {
	var lowestCity string
	var lowestTemp float64 = cities[0].AvgTemperature
	for _, city := range cities {
		if city.AvgTemperature < lowestTemp {
			lowestTemp = city.AvgTemperature
			lowestCity = city.CityName
		}
	}
	return lowestCity, lowestTemp
}

func averageRainfall(cities []CityData) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []CityData, threshold float64) {
	fmt.Println("\nCities with rainfall above", threshold, "mm:")
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("%s: %.2f mm\n", city.CityName, city.Rainfall)
		}
	}
}

func searchByCityName(cities []CityData, cityName string) {
	cityName = strings.ToLower(cityName)
	found := false
	for _, city := range cities {
		if strings.ToLower(city.CityName) == cityName {
			fmt.Printf("\nCity: %s\nAverage Temperature: %.2f°C\nRainfall: %.2f mm\n", city.CityName, city.AvgTemperature, city.Rainfall)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("City not found.")
	}
}

func main() {
	cities := []CityData{
		{"New York", 15.2, 120.4},
		{"London", 10.8, 114.5},
		{"Tokyo", 16.3, 120.8},
		{"Paris", 12.5, 90.1},
		{"Sydney", 18.4, 102.3},
	}

	highestCity, highestTemp := highestTemperature(cities)
	lowestCity, lowestTemp := lowestTemperature(cities)
	fmt.Printf("\nCity with the highest temperature: %s (%.2f°C)\n", highestCity, highestTemp)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowestCity, lowestTemp)

	avgRainfall := averageRainfall(cities)
	fmt.Printf("\nAverage rainfall across all cities: %.2f mm\n", avgRainfall)

	var threshold float64
	fmt.Print("\nEnter the rainfall threshold (in mm) to filter cities: ")
	_, err := fmt.Scan(&threshold)
	if err != nil || threshold < 0 {
		fmt.Println("Invalid input. Please enter a positive number for rainfall threshold.")
		return
	}
	filterCitiesByRainfall(cities, threshold)

	var cityName string
	fmt.Print("\nEnter the city name to search: ")
	fmt.Scan(&cityName)
	searchByCityName(cities, cityName)
}
