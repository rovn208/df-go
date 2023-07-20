package util

import "testing"

func TestGetCountryCode(t *testing.T) {
	args := []string{"Ro", "Ngoc", "Vo", "VN"}
	countryCode, ok := GetCountryCode(args)
	if !ok {
		t.Errorf("Expected get country code successfully, got false")
	}
	if countryCode != "VN" {
		t.Errorf("Expected VN, got %s", countryCode)
	}
}

func TestGetCountryCodeUnsupported(t *testing.T) {
	args := []string{"Ro", "Ngoc", "Vo", "VNN"}
	_, ok := GetCountryCode(args)

	if ok {
		t.Errorf("Expected false, got true")
	}
}

func TestGetNameOrder(t *testing.T) {
	testCases := []struct {
		countryCode string
		expected    NameOrder
	}{
		{"US", WESTERN},
		{"VN", EASTERN},
		{"DE", EASTERN},
	}

	for _, testCase := range testCases {
		actual := GetNameOrder(testCase.countryCode)
		if actual != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, actual)
		}
	}
}
