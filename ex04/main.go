package main

import (
	"fmt"
	"github.com/rovn208/df-go/ex04/counter"
)

func main() {
	word := "a123bc34d8ef34"
	simpleCounter := counter.NewCounter(word)

	fmt.Println(simpleCounter.GetDifferentIntegers())
}
