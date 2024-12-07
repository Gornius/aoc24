package main

import (
	"fmt"
	"strconv"
)

type Operator string

const (
	OperatorAdd      Operator = "+"
	OperatorMultiply Operator = "*"
	OperatorConcat   Operator = "||"
)

type Equation struct {
	TestValue uint64
	Operands  []int
}

func main() {
	equations, err := parseFile("./challenge_data.txt")
	if err != nil {
		panic(err)
	}

	compliantEquations := 0
	var possibleEquationsSumOfTestValues uint64 = 0

	for _, equation := range equations {
		possiblePermutations := equation.GetAllOperatorPermutations()
		for _, permutation := range possiblePermutations {
			if equation.IsEquationPossible(permutation) {
				compliantEquations++
				possibleEquationsSumOfTestValues += equation.TestValue
				break
			}
		}
	}

	fmt.Printf("====================== PART 1 =======================\n")
	fmt.Printf("compliantEquations: %v\n", compliantEquations)
	fmt.Printf("possibleEquationsSumOfTestValues: %v\n", possibleEquationsSumOfTestValues)

	//=====================

	compliantEquations = 0
	possibleEquationsSumOfTestValues = 0

	for _, equation := range equations {
		possiblePermutations := equation.GetAllOperatorPermutationsWithConcat()
		for _, permutation := range possiblePermutations {
			if equation.IsEquationPossible(permutation) {
				compliantEquations++
				possibleEquationsSumOfTestValues += equation.TestValue
				break
			}
		}
	}

	fmt.Printf("======================= PART 2 =======================\n")
	fmt.Printf("compliantEquations: %v\n", compliantEquations)
	fmt.Printf("possibleEquationsSumOfTestValues: %v\n", possibleEquationsSumOfTestValues)
}

func (e *Equation) IsEquationPossible(operators []Operator) bool {
	acc := uint64(e.Operands[0])

	for i, operator := range operators {
		secondOperand := e.Operands[i+1]
		switch operator {
		case OperatorAdd:
			acc += uint64(secondOperand)
		case OperatorMultiply:
			acc *= uint64(secondOperand)
		case OperatorConcat:
			acc = concat(acc, uint64(secondOperand))
		}
	}

	return acc == e.TestValue
}

func (e *Equation) GetAllOperatorPermutations() [][]Operator {
	n := len(e.Operands) - 1
	allowedOperators := []Operator{"+", "*"}

	permutations := [][]Operator{}

	generatePermutations(allowedOperators, n, []Operator{}, &permutations)

	return permutations
}

func (e *Equation) GetAllOperatorPermutationsWithConcat() [][]Operator {
	n := len(e.Operands) - 1
	allowedOperators := []Operator{"+", "*", "||"}

	permutations := [][]Operator{}

	generatePermutations(allowedOperators, n, []Operator{}, &permutations)

	return permutations
}

func generatePermutations(allowedOperators []Operator, n int, current []Operator, results *[][]Operator) {
	if len(current) == n {
		result := []Operator{}
		result = append(result, current...)
		*results = append(*results, result)
		return
	}

	for _, operator := range allowedOperators {
		generatePermutations(allowedOperators, n, append(current, operator), results)
	}
}

func concat(a uint64, b uint64) uint64 {
	aStr := strconv.FormatUint(a, 10)
	bStr := strconv.FormatUint(b, 10)

	concatInt, err := strconv.ParseUint(aStr+bStr, 10, 64)
	if err != nil {
		panic(err)
	}

	return concatInt
}
