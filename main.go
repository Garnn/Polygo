package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/garnn/Polygo/pkg/equation"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for a polynomial, try to read it, and panic if it can't be read
	fmt.Print("Enter a polynomial: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Text reader encountered an error: %v", err)
	}

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")

	// Parse the polynomial into a more workable format thanks to Gucio's Parser, panic if it fails somehow
	result, err := equation.Parse(text)
	if err != nil {
		log.Fatalf("Equation parser encountered an error: %v", err)
	}

	// Check if polynomials can convert back into strings
	fmt.Println(result)
}
