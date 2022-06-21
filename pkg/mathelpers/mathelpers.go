package mathelpers

func BruteGetDivisors(i int) (divisors []int) {
	for divisor := 1; divisor <= i/2; divisor += 1 {
		if i%divisor == 0 {
			divisors = append(divisors, divisor)
		}
	}

	return divisors
}
