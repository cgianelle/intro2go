package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildItinerary(cities []string) (string, string) {
	itinerary := ""
	airports := map[string]string{
		"ATL": "Atlanta",
		"PEK": "Beijing",
		"ORD": "Chicago",
		"LHR": "London",
		"HND": "Tokyo",
		"LAX": "Los Angeles",
		"CDG": "Paris",
		"DFW": "Dallas",
		"FRA": "Frankfurt",
		"DEN": "Denver",
		"HKG": "Hong Kong",
		"MAD": "Madrid",
		"DXB": "Dubai",
		"JFK": "New York",
		"AMS": "Amsterdam",
		"CGK": "Jakarta",
		"BKK": "Bangkok",
		"SIN": "Singapore",
		"CAN": "Guangzhou",
		"PVG": "Shanghai",
	}

	insertCity := func(city string) {
		if itinerary == "" {
			itinerary = city
		} else {
			itinerary = itinerary + " TO " + city
		}
	}

	for _, c := range cities {
		city := airports[c]
		if city == "" {
			a := "Unknown destination, " + c
			return a, ""
		}
		insertCity(city)
	}
	return "", itinerary
}

func main() {
	var cities []string
	fmt.Print("Enter the airport codes of your trip:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	cities = strings.Split(scanner.Text(), " ")
	errStr, itin := buildItinerary(cities)
	if errStr != "" {
		fmt.Println(errStr)
	} else {
		fmt.Println(itin)
	}
}
