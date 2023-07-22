package main

import (
	"bytes"
	"errors"
	"github.com/rovn208/df-go/ex02/cmd"
	"github.com/rovn208/df-go/ex02/util"
	"io"
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
		{"sort integer", toArray("-i 1 4 2"), "1 2 4 ", nil},
		{"sort float", toArray("--float 1 4.4 0 4.2"), "0 1 4.2 4.4 ", nil},
		{"sort string", toArray("-s apple orange banana"), "apple banana orange ", nil},
		{"sort mix", toArray("-m apple 1 orange 4 banana 4.2"), "1 4 4.2 apple banana orange ", nil},
		{"empty args", toArray("-m"), "", errors.New("requires at least 1 arg(s), only received 0")},
		{"wrong type of integer", toArray("-i 1 4 2.2"), "", errors.New("invalid argument: 2.2")},
		{"wrong type of integer", toArray("-i 1 4 a"), "", errors.New("invalid argument: a")},
		{"wrong type of float", toArray("-f 2.2 a"), "", errors.New("invalid argument: a")},
		{"wrong type of string", toArray("-s 2.2 a"), "", errors.New("invalid argument: 2.2")},
		{"wrong type of string", toArray("-s 2 a"), "", errors.New("invalid argument: 2")},
		{"input type is missing", toArray("apple"), "", util.InputTypeMissingError},
		{"more than one input types", toArray("-m -i apple 1 orange 4 banana 4.2"), "1 4 4.2 apple banana orange ",
			errors.New("if any flags in the group [int float string mix] are set none of the others can be; [int mix] were all set")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cli := cmd.New()
			b := bytes.NewBufferString("")
			cli.SetOut(b)
			cli.SetArgs(tc.args)
			err := cli.Execute()

			if tc.expectedError != nil {
				if err == nil || (err.Error() != tc.expectedError.Error()) {
					t.Fatalf("expected \"%s\" got \"%s\"", tc.expectedError, err)
				}
				return
			}

			out, err := io.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}

			if string(out) != tc.expectedOutput {
				t.Fatalf("expected \"%s\" got \"%s\"", tc.expectedOutput, string(out))
			}
		})
	}
}

func toArray(str string) []string {
	return strings.Split(str, " ")
}
