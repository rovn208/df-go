package main

import (
	"github.com/rovn208/df-go/ex03/counter"
	"testing"
)

func TestCountRectangles(t *testing.T) {
	testCases := []struct {
		name          string
		arr           [][]int
		expectedCount int
	}{
		{"empty array", [][]int{}, 0},
		{"1x1, nx1, mxn",
			[][]int{
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 1}},
			6,
		},
		{"adjacent rectangles",
			[][]int{
				{1, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1, 1, 0},
				{1, 1, 0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0}},
			0,
		},
		{"invalid form rectangle",
			[][]int{
				{1, 1, 1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 0, 1, 1, 0},
				{0, 0, 0, 1, 1, 1, 0},
				{1, 1, 0, 1, 0, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 0, 0, 0}},
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rectangleCounter := counter.NewCounter(tc.arr)
			rectangles := rectangleCounter.CountRectangles()

			if rectangles != tc.expectedCount {
				t.Fatalf("Expected %d, got %d", tc.expectedCount, rectangles)
			}
		})
	}
}
