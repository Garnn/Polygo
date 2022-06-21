package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"main/pkg/equation"
)

func main() {
	fmt.Print("Enter your equation: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading from input: %v", err)
	}

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	result, err := equation.Parse(text)
	if err != nil {
		log.Fatalf("error parsing equation: %v", err)
	}

	result.Simplify()
	fmt.Println(result)
}