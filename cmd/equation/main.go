package main

import (
	"fmt"
	"log"
	"main/pkg/equation"
)

func main() {
	fmt.Print("Enter your equation: ")
	var text string
	fmt.Scanln(&text)
	result, err := equation.Parse(text)
	if err != nil {
		log.Fatalf("error parsing equation: %v", err)
	}

	fmt.Println(result)
	fmt.Println(*result.Monomials[0])
}
