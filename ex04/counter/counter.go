package counter

import "strconv"

type SimpleCounter struct {
	word string
}

type Counter interface {
	GetDifferentIntegers() (error, int)
}

func NewCounter(word string) *SimpleCounter {
	return &SimpleCounter{word: word}
}

func (c *SimpleCounter) GetDifferentIntegers() int {
	word := c.word
	dict := make(map[string]string)
	num := ""

	for _, ch := range word {
		char := string(ch)
		_, err := strconv.Atoi(char)
		if err != nil {
			if _, ok := dict[num]; !ok && len(num) > 0 {
				dict[num] = num
			}
			num = ""
			continue

		} else {
			num += char
		}
	}

	return len(dict)
}
