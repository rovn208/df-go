package util

import (
	"fmt"
)

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
	WESTERN NameOrder = "western"
	EASTERN NameOrder = "eastern"
)

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
