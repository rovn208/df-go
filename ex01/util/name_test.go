package util

import "testing"

func TestReOrderName(t *testing.T) {
	testCases := []struct {
		args         []string
		expectedName string
	}{
		{[]string{"Ro", "Ngoc", "Vo", "VN"}, "Vo Ngoc Ro"},
		{[]string{"Ro", "Vo", "VN"}, "Vo Ro"},
		{[]string{"Ro", "Ngoc", "Vo", "US"}, "Ro Ngoc Vo"},
		{[]string{"Ro", "Vo", "US"}, "Ro Vo"},
		{[]string{"Ro", "Ngoc", "Vo", "VNN"}, ""},
	}

	for _, testCase := range testCases {
		actualName, _ := ReOrderName(testCase.args)

		if actualName != testCase.expectedName {
			t.Errorf("Expected %s, got %s", testCase.expectedName, actualName)
		}
	}
}
