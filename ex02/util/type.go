package util

import (
	"flag"
)

type InputType string

const (
	INTEGER InputType = "int"
	FLOAT   InputType = "float"
	STRING  InputType = "string"
	MIX     InputType = "mix"
)

var SupportedInputTypes = []InputType{INTEGER, FLOAT, STRING, MIX}

// GetInputType returns the input type
func GetInputType() InputType {
	for _, t := range SupportedInputTypes {
		if isInputTypeSet(t) {
			return t
		}
	}
	return MIX
}

// ValidateOneRequired validates one required flag
func ValidateOneRequired() error {
	count := 0
	for _, inputType := range SupportedInputTypes {
		if isInputTypeSet(inputType) {
			count++
		}
	}

	if count == 0 {
		return InputTypeMissingError
	}

	if count > 1 {
		return MoreThanOneInputTypeError
	}

	return nil
}

func isInputTypeSet(inputType InputType) bool {
	if t := flag.Lookup(string(inputType)); t != nil && t.Value.String() == "true" {
		return true
	}
	return false
}
