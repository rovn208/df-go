package main

import (
	"fmt"
	counter2 "github.com/rovn208/df-go/ex04/counter"
)

func main() {
	word := "a123bc34d8ef34"
	counter := counter2.NewCounter(word)

	fmt.Println(counter.GetDifferentIntegers())
}
