package main

import (
	"errors"
	"flag"
	"github.com/rovn208/df-go/ex02/cmd"
	"github.com/rovn208/df-go/ex02/util"
	"os"
	"strings"
	"testing"
)

func TestMainExecutor(t *testing.T) {
	testCases := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  error
	}{
		{"sort integer", toArray("-int 1 4 2"), "1 2 4 ", nil},
		{"sort float", toArray("-float 1 4.4 0 4.2"), "0 1 4.2 4.4 ", nil},
		{"sort string", toArray("-string apple orange banana"), "apple banana orange ", nil},
		{"sort mix", toArray("-mix apple 1 orange 4 banana 4.2"), "1 4 4.2 apple banana orange ", nil},
		{"empty args", toArray("-mix"), "", errors.New("requires at least 1 arg(s), only received 0")},
		{"wrong type of integer", toArray("-int 1 4 2.2"), "", errors.New("invalid argument: 2.2")},
		{"wrong type of integer", toArray("-int 1 4 a"), "", errors.New("invalid argument: a")},
		{"wrong type of float", toArray("-float 2.2 a"), "", errors.New("invalid argument: a")},
		{"wrong type of string", toArray("-string 2.2 a"), "", errors.New("invalid argument: 2.2")},
		{"wrong type of string", toArray("-string 2 a"), "", errors.New("invalid argument: 2")},
		{"input type is missing", toArray("apple"), "", util.InputTypeMissingError},
		{"more than one input types", toArray("-mix -int apple 1 orange 4 banana 4.2"), "1 4 4.2 apple banana orange ",
			errors.New("more than one input type is set")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resetFlags()
			os.Args = tc.args
			sortedString, err := cmd.Execute()

			if tc.expectedError != nil {
				if err == nil || (err.Error() != tc.expectedError.Error()) {
					t.Fatalf("expected \"%s\" got \"%s\"", tc.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil got \"%s\"", err)
			}

			if sortedString != tc.expectedOutput {
				t.Fatalf("expected \"%s\" got \"%s\"", tc.expectedOutput, sortedString)
			}
		})
	}
}

func toArray(str string) []string {
	args := []string{"test"}
	return append(args, strings.Split(str, " ")...)
}

func resetFlags() {
	for _, inputType := range util.SupportedInputTypes {
		if t := flag.Lookup(string(inputType)); t != nil && t.Value.String() == "true" {
			err := t.Value.Set("false")
			if err != nil {
				return
			}
		}
	}
}
