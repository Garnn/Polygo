package main

import (
	"fmt"

	"github.com/garnn/Polygo/pkg/equation"
)

func main() {
	var text string

	fmt.Print("Enter a polynomial: ")
	fmt.Scanln(&text)

	//text = strings.ReplaceAll(text, "\r", "")
	//text = strings.ReplaceAll(text, "\n", "")

	//num, err := strconv.Atoi(text)
	//if err != nil {
	//	fmt.Printf("Atoi encountered an error: %v", err)
	//}

	//fmt.Println(mathelpers.BruteGetDivisors(num))
	res, err := equation.Parse(text)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
