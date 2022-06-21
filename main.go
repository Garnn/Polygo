package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/garnn/Polygo/pkg/mathelpers"
)

func main() {
	var text string

	fmt.Print("Enter a number: ")
	fmt.Scanln(&text)

	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")

	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Printf("Atoi encountered an error: %v", err)
	}

	fmt.Println(mathelpers.BruteGetDivisors(num))
}
