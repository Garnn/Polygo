package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		print("ReadString encountered an error: ")
		print(err)
	}
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	num, err := strconv.Atoi(text)
	if err != nil {
		print("Atoi encountered an error: ")
		print(err)
	}
	fmt.Println(bruteGetDivisors(num))
}

func bruteGetDivisors(i int) []int {
	var divisors []int
	for divisor := 1; divisor <= i/2; divisor += 1 {
		if i%divisor == 0 {
			divisors = append(divisors, divisor)
		}
	}
	return divisors
}
