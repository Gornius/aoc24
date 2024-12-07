package main

import (
	"strconv"
	"strings"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func parseFile(filePath string) ([]Equation, error) {
	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	equations := []Equation{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")

		testVal, err := strconv.ParseUint(parts[0], 10, 64)
		if err != nil {
			return nil, err
		}

		operands := []int{}
		operandsString := parts[1]

		for _, operandString := range strings.Split(operandsString, " ") {
			operand, err := strconv.Atoi(operandString)
			if err != nil {
				return nil, err
			}

			operands = append(operands, operand)
		}

		equations = append(equations, Equation{
			TestValue: testVal,
			Operands:  operands,
		})
	}

	return equations, nil
}
