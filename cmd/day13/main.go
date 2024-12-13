package main

import (
	"fmt"
	"math"
)

func main() {
	filePath := "./challenge_data.txt"

	fmt.Println("==== PART 1 ====")
	part1(filePath)

	fmt.Println("==== PART 2 ====")
	part2(filePath)
}

func part1(filePath string) {
	machines, err := ParseFile(filePath)
	if err != nil {
		panic(err)
	}

	var totalTokensUsed uint64 = 0

	for i, machine := range machines {
		fmt.Println("===== MACHINE ", i, " ====")

		A := machine.GetAButtonPresses()
		B := machine.GetBButtonPresses()

		if A != math.Trunc(A) || B != math.Trunc(B) {
			fmt.Println("No answers")
			continue
		}

		fmt.Printf("Answer: A=%g, B=%g\n", A, B)
		usedTokens := uint64(A)*3 + uint64(B)*1
		totalTokensUsed += usedTokens
		fmt.Printf("Used tokens: %d, total tokens used: %d\n", usedTokens, totalTokensUsed)
	}
}

func part2(filePath string) {
	machines, err := ParseFile(filePath)
	if err != nil {
		panic(err)
	}

	var totalTokensUsed uint64 = 0

	for i, machine := range machines {
		machine.P.X = machine.P.X + 10_000_000_000_000
		machine.P.Y = machine.P.Y + 10_000_000_000_000
		fmt.Println("===== MACHINE ", i, " ====")

		A := machine.GetAButtonPresses()
		B := machine.GetBButtonPresses()

		if A != math.Trunc(A) || B != math.Trunc(B) {
			fmt.Println("No answers")
			continue
		}

		fmt.Printf("Answer: A=%g, B=%g\n", A, B)
		usedTokens := uint64(A)*3 + uint64(B)*1
		totalTokensUsed += usedTokens
		fmt.Printf("Used tokens: %d, total tokens used: %d\n", usedTokens, totalTokensUsed)
	}
}

type Machine struct {
	A, B, P Params
}

// U want sum maths? (Cramer's rule)
func (m *Machine) GetAButtonPresses() float64 {
	return (float64(m.P.X)*float64(m.B.Y) - float64(m.B.X)*float64(m.P.Y)) /
		(float64(m.A.X)*float64(m.B.Y) - float64(m.B.X)*float64(m.A.Y))
}

// U want sum maths? (Cramer's rule)
func (m *Machine) GetBButtonPresses() float64 {
	return (float64(m.A.X)*float64(m.P.Y) - float64(m.P.X)*float64(m.A.Y)) /
		(float64(m.A.X)*float64(m.B.Y) - float64(m.B.X)*float64(m.A.Y))
}

type Params struct {
	X, Y uint64
}
