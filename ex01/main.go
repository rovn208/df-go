package main

import (
	"fmt"
	"github.com/rovn208/df-go/ex01/util"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Invalid input: Missing arguments\nUsage: go run main.go [full name] [language]\nExample: Go run main.go Ro Ngoc Vo VN")
		os.Exit(1)
	}
	countryCode, ok := util.GetCountryCode(args)
	if !ok {
		log.Fatal("Error when getting countryCode")
	}

	fmt.Println(util.ReOrderName(countryCode, args))
}
