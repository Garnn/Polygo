package polynomials

import (
	"github.com/garnn/Polygo/pkg/mathelpers"
)

func findPotentialSolutions(equation Equation) (solutions []int) {
	divisors_1 := mathelpers.BruteGetDivisors(equation.Monomials[len(equation.Monomials)-1].Coefficient)
	divisors_2 := mathelpers.BruteGetDivisors(equation.Monomials[0].Coefficient)
	solutions = divisors_1
	for _, div1 := range divisors_1 {
		for _, div2 := range divisors_2 {
			solutions = append(solutions, div1/div2)
		}
	}
	return solutions
}
