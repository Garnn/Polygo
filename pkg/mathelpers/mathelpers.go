package mathelpers

import (
	"fmt"
	"math"

	"github.com/cznic/mathutil"
)

//BruteGetDivisors is a really inefficient algorithm that tries to find all divisors of a number by pure brute force,
//DO NOT overuse it, it's seriously slow, a better algorithm would factorize the number and then get divisors by multiplying
//it's factors
//TODO: Maybe try to do the factorization approach with some fancy, slightly better than brute force algorithm

func BruteGetDivisors(i int) (divisors []int) {
	if i < 0 {
		i *= -1
	}
	for divisor := 1; divisor <= i/2; divisor++ {
		if i%divisor == 0 {
			divisors = append(divisors, divisor)
			divisors = append(divisors, -divisor)
		}
	}
	divisors = append(divisors, i)
	divisors = append(divisors, -i)
	return divisors
}

//This function takes a calculated root, and spits back a human-readable version of it as a string

func MakeRootReadable(i float64) (output string) {
	orig_num := math.Round(i * i)
	factors := mathutil.FactorInt(uint32(orig_num))
	coeff, root := 1, 1
	for i := 0; i < len(factors); i++ {
		if factors[i].Power%2 != 0 {
			root *= int(factors[i].Prime)
		}
		if factors[i].Power/2 != 0 {
			coeff *= int(factors[i].Power/2) * int(factors[i].Prime)
		}
	}
	if coeff == 1 {
		return fmt.Sprintf("√%v", root)
	}
	output = fmt.Sprintf("%v√%v", coeff, root)
	return output
}
