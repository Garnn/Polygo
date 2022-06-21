package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/garnn/Polygo/pkg/equation"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a polynomial: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Text reader encountered an error: %v", err)
	}
	//text = strings.ReplaceAll(text, "\r", "")
	//text = strings.ReplaceAll(text, "\n", "")

	//num, err := strconv.Atoi(text)
	//if err != nil {
	//	fmt.Printf("Atoi encountered an error: %v", err)
	//}

	//fmt.Println(mathelpers.BruteGetDivisors(num))
	result, err := equation.Parse(text)
	if err != nil {
		log.Fatalf("Equation parser encountered an error: %v", err)
	}
	fmt.Println(result)
}
