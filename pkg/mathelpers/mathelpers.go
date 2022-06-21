package mathelpers

//BruteGetDivisors is a really inefficient algorithm that tries to find all divisors of a number by pure brute force,
//DO NOT overuse it, it's seriously slow, a better algorithm would factorize the number and then get divisors by multiplying
//it's factors
//TODO: Maybe try to do the factorization approach with some fancy, slightly better than brute force algorithm

func BruteGetDivisors(i int) (divisors []int) {
	for divisor := 1; divisor <= i/2; divisor += 1 {
		if i%divisor == 0 {
			divisors = append(divisors, divisor)
		}
	}

	return divisors
}
