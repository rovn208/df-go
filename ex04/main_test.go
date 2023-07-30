package main

import (
	"github.com/rovn208/df-go/ex04/counter"
	"testing"
)

func TestGetDifferentIntegers(t *testing.T) {
	testCases := []struct {
		name          string
		word          string
		expectedCount int
	}{
		{"empty string", "", 0},
		{"happy case", "a123bc34d8ef34", 3},
		{"word with float number", "a123.311bc34d8ef34", 2},
		{"word with special characters", "a123.311bEc34d語8ef34", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			simpleCounter := counter.NewCounter(tc.word)
			actualCount := simpleCounter.GetDifferentIntegers()
			if actualCount != tc.expectedCount {
				t.Errorf("expected %v, got %v", tc.expectedCount, actualCount)
			}
		})
	}
}
