package util

import (
	"strings"
)

func ReOrderName(order NameOrder, args []string) string {
	firstName, middleName, lastName := getNameParts(args)

	if middleName == "" {
		return getName(order, firstName, lastName)
	}
	return getName(order, firstName, middleName, lastName)
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
