package equation

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// UnexpectedEquationError is returned by Parse if there is something wrong with equation string given
// you cann reffer to this error by calling errors.Is(err, UnexpectedEquationError) in order to detect it.
var UnexpectedEquationError = errors.New("unexpected equation")

type Equation struct {
	Monomials []*Monomial
}

func Parse(equation string) (result *Equation, err error) {
	result = &Equation{}
	// remove all spaces - they are user-friendly but they are not necessary
	equation = strings.ReplaceAll(equation, " ", "")

	// split to right ane left sides
	splited := strings.Split(equation, "=")

	var leftSideStr, rightSideStr string

	switch len(splited) {
	// only left side - x^2 + 2x = 0
	case 1:
		leftSideStr = splited[0]
	// standard equation x^2 + 2x = 8x + 2
	case 2:
		leftSideStr, rightSideStr = splited[0], splited[1]
	// wtf???
	default:
		return nil, UnexpectedEquationError
	}

	leftSide, rightSide := parseStr(leftSideStr), parseStr(rightSideStr)

	// move eveerything from right to left side

	for _, x := range rightSide {
		x.ChangeSide()
		leftSide = append(leftSide, x)
	}

	result.Monomials = leftSide

	return result, nil
}

func parseStr(s string) []*Monomial {
	result := make([]*Monomial, 0)
	if len(s) == 0 {
		return result
	}

	if s[0] != '-' {
		result = append(result, newMonomial())
	}

	for i := 0; i < len(s); i++ {
		currentLetter := s[i]
		switch {
		case currentLetter == '-':
			m := newMonomial()
			m.Positive = true
			result = append(result, m)
		case currentLetter == '+':
			result = append(result, newMonomial())
		case currentLetter == '^':
			last := result[len(result)-1]
			p := "0"
			for j := i + 1; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
				p += string(s[j])
			}

			power, err := strconv.Atoi(p)
			if err != nil {
				log.Panicf("Unexpected error occured - this couldn't have happend: %v", err)
			}

			last.Power = power

			// -1 for base 0 in p value
			i += len(p) - 1
		case currentLetter >= '0' && currentLetter <= '9':
			last := result[len(result)-1]
			c := "0"
			for j := i; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
				c += string(s[j])
			}

			coefficient, err := strconv.Atoi(c)
			if err != nil {
				log.Panicf("Unexpected error occured - this couldn't have happend: %v", err)
			}

			last.Coefficient = coefficient

			// - 1 for base 0 value of c and -1 for i++ in for assignment
			i += len(c) - 2
		default:
			last := result[len(result)-1]
			last.Power = 1
			last.Letter = rune(s[i])
		}
	}

	return result
}

// String implements fmt.Stringer
func (e *Equation) String() string {
	result := ""
	for _, i := range e.Monomials {
		result += i.String() + " "
	}

	return fmt.Sprintf("%v = 0", result)
}

type Monomial struct {
	Positive    bool
	Coefficient int
	Letter      rune
	Power       int
}

func newMonomial() *Monomial {
	return &Monomial{
		Positive:    true,
		Coefficient: 1,
		Letter:      'x',
		Power:       0,
	}
}

// String implements fmt.Stringer interface.
// it is like ToString in C# - makes fmt.Print and familiar to print
// this as human-readable
func (m *Monomial) String() string {
	positiveStr := ""
	if m.Positive {
		positiveStr = "+"
	} else {
		positiveStr = "-"
	}

	var letter rune
	if m.Power > 0 {
		letter = m.Letter
	}

	return fmt.Sprintf("%s%d%c^%d", positiveStr, m.Coefficient, letter, m.Power)
}

func (s *Monomial) ChangeSide() {
	s.Positive = !s.Positive
}
