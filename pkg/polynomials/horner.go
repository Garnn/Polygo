package polynomials

import (
	"math"

	"github.com/garnn/Polygo/pkg/mathelpers"
)

func convIntArrayToFloat32Array(array []int) (output []float64) {
	for i := range array {
		output = append(output, float64(array[i]))
	}
	return output
}

//This function takes an Equation and attempts to find all possible solutions,
//some solutions may not be valid and must be weeded out with the checkValidSolutions helper function
func findPotentialSolutions(equation *Equation) (solutions []float64) {
	divisors_1int := mathelpers.BruteGetDivisors(equation.Monomials[len(equation.Monomials)-1].Coefficient)
	divisors_2int := mathelpers.BruteGetDivisors(equation.Monomials[0].Coefficient)
	divisors_1 := convIntArrayToFloat32Array(divisors_1int)
	divisors_2 := convIntArrayToFloat32Array(divisors_2int)
	solutionsUnprocessed := divisors_1
	//Calculate rest of possible solutions
	for _, div1 := range divisors_1 {
		for _, div2 := range divisors_2 {
			solutionsUnprocessed = append(solutionsUnprocessed, div1/div2)
		}
	}
	//Mark duplicates
	for i_1, sol := range solutionsUnprocessed {
		if sol != 0 {
			for i_2 := i_1 + 1; i_2 < len(solutionsUnprocessed); i_2++ {
				if solutionsUnprocessed[i_2] == sol {
					solutionsUnprocessed[i_2] = 0
				}
			}
		}
	}
	//Erase duplicates
	for i := range solutionsUnprocessed {
		if solutionsUnprocessed[i] != 0 {
			solutions = append(solutions, solutionsUnprocessed[i])
		}
	}
	return solutions
}

//This function takes an Equation, and a list of possible solutions, and weeds out the wrong ones
func checkValidSolutions(equation *Equation, posSolutions []float64) (solutions []float64) {
	var evaluatedAnswer float64
	for _, sol := range posSolutions {
		evaluatedAnswer = 0
		for _, elem := range equation.Monomials {
			evaluatedAnswer += float64(elem.Coefficient) * math.Pow(sol, float64(elem.Power))
		}
		if evaluatedAnswer == 0.0 {
			solutions = append(solutions, sol)
		}
	}
	return solutions
}
func hornerize(equation *Equation, solution1 float64) (newEquation Equation, solution2 float64) {
	curr := solution1
	for i, elem := range equation.Monomials {
		newMono := newMonomial()
		newMono.Coefficient = int(curr)

	}
}

func SolveEquation(equation *Equation) (finalSolutions []float64) {
	dirtySolutions := findPotentialSolutions(equation)
	cleanSolutions := checkValidSolutions(equation, dirtySolutions)
	currEquation := *equation

}
