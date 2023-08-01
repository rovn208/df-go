package counter

import (
	"strconv"
)

type SimpleCounter struct {
	word string
}

type Counter interface {
	GetDifferentIntegers() (error, int)
}

// NewCounter creates a new counter
func NewCounter(word string) *SimpleCounter {
	return &SimpleCounter{word: word}
}

// GetDifferentIntegers returns the number of different integers
func (c *SimpleCounter) GetDifferentIntegers() int {
	word := c.word
	dict := make(map[string]string)
	num := ""
	isFloat := false

	for _, ch := range word {
		char := string(ch)
		_, err := strconv.Atoi(char)
		if err != nil {
			// Float number
			if char == "." {
				isFloat = true
				continue
			}
			if len(num) > 0 {
				numInt, _ := strconv.Atoi(num)
				num = strconv.Itoa(numInt)
				if _, ok := dict[num]; !ok && len(num) > 0 {
					if !isFloat {
						dict[num] = num
						num = ""
						continue
					}
					isFloat = false
				}
				num = ""
				continue
			}
		} else {
			num += char
		}

	}
	return len(dict)
}
