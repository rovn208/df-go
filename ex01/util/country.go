package util

import (
	"fmt"
)

// GetCountryCode returns a country is its supported, otherwise return empty
func GetCountryCode(args []string) (string, bool) {
	country := args[len(args)-1]
	if !isValidCountry(country) {
		fmt.Printf("Country code %s is not supported\n", country)
		return "", false
	}
	return country, true
}

func isValidCountry(countryCode string) bool {
	supportedCountryCodes := []string{"VN", "US"}
	for _, ctCode := range supportedCountryCodes {
		if ctCode == countryCode {
			return true
		}
	}

	return false
}

type NameOrder string

const (
	// WESTERN aka Western name order has the order given name(s), family name
	WESTERN NameOrder = "western"
	// EASTERN aka Western name order has the order family name, given name(s)
	EASTERN NameOrder = "eastern"
)

// GetNameOrder returns NameOrder based on country code. Using EASTERN as default value.
func GetNameOrder(countryCode string) NameOrder {
	switch countryCode {
	case "US":
		return WESTERN
	case "VN":
		return EASTERN
	default:
		return EASTERN
	}
}
