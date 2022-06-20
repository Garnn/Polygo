package main

import (
	"fmt"
	"strconv"
	"strings"
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

	fmt.Println(bruteGetDivisors(num))
}

func bruteGetDivisors(i int) (divisors []int) {
	for divisor := 1; divisor <= i/2; divisor += 1 {
		if i%divisor == 0 {
			divisors = append(divisors, divisor)
		}
	}

	return divisors
}
