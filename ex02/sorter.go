package main

import (
	"fmt"
	"github.com/rovn208/df-go/ex02/cmd"
	"log"
)

func main() {
	sortedString, err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sortedString)
}
