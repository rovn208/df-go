package util

// Contains is a helper function to check if a string is in a slice
func Contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}
