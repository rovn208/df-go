package util

import (
	"strings"
)

// ReOrderName returns a full name which has order based on the country Code
func ReOrderName(args []string) (string, bool) {
	countryCode, ok := GetCountryCode(args)
	if !ok {
		return "", false
	}

	nameOrder := GetNameOrder(countryCode)
	firstName, middleName, lastName := getNameParts(args)

	if middleName == "" {
		return getName(nameOrder, firstName, lastName), true
	}
	return getName(nameOrder, firstName, middleName, lastName), true
}

func getName(order NameOrder, nameParts ...string) string {
	if order == WESTERN {
		return strings.Join(nameParts, " ")
	}
	return strings.Join(ReverseArray(nameParts), " ")
}

func getNameParts(args []string) (firstName string, middleName string, lastName string) {
	firstName = args[0]
	lastName = args[len(args)-2]
	middleName = strings.Join(args[1:len(args)-2], " ")
	return
}
