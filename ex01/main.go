package main

import (
	"fmt"
	"github.com/rovn208/df-go/ex01/util"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Invalid input: Missing arguments\nUsage: go run main.go [full name] [language]\nExample: Go run main.go Ro Ngoc Vo VN")
		os.Exit(1)
	}

	name, ok := util.ReOrderName(args)
	if !ok {
		fmt.Println("Error when reordering name")
		os.Exit(1)
	}
	fmt.Println(name)
}
