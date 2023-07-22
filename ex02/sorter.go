package main

import (
	"github.com/rovn208/df-go/ex02/cmd"
	"log"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		log.Fatal(err)
	}
}
