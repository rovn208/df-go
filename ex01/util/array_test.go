package util

import (
	"testing"
)

func TestReverseArray(t *testing.T) {
	arr := []string{"a", "b", "c"}
	expectedArr := []string{"c", "b", "a"}

	arr = ReverseArray(arr)

	if len(arr) != len(expectedArr) {
		t.Errorf("Expected %v, got %v", expectedArr, arr)
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != expectedArr[i] {
			t.Errorf("Expected %v, got %v", expectedArr, arr)
		}
	}
}
