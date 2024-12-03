package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gornius/aoc24/pkg/fileutils"
)

type Memory struct {
	Content string
}

type Mul struct {
	A int
	B int
}

func main() {
	memory, err := parseFile("./test_data_2.txt")
	if err != nil {
		panic(err)
	}

	// ========================

	muls, err := memory.getMuls()
	if err != nil {
		panic(err)
	}

	sumFirstPart := 0
	for _, mul := range muls {
		sumFirstPart += mul.A * mul.B
	}
	fmt.Printf("sumFirstPart: %v\n", sumFirstPart)

	// ========================

	muls2, err := memory.getMulsWithControlInstructions()
	if err != nil {
		panic(err)
	}

	sumSecondPart := 0
	for _, mul := range muls2 {
		sumSecondPart += mul.A * mul.B
	}

	fmt.Printf("sumSecondPart: %v\n", sumSecondPart)
}

func (m *Memory) getMuls() ([]Mul, error) {
	muls := []Mul{}

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := regex.FindAllStringSubmatch(m.Content, -1)

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}
		a, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, err
		}
		muls = append(muls, Mul{a, b})
	}

	return muls, nil
}

func (m *Memory) getMulsWithControlInstructions() ([]Mul, error) {
	muls := []Mul{}

	getSubsequences := func() []string {
		regex := regexp.MustCompile(`do\(\)(.*?)don't\(\)`)

		matches := regex.FindAllString("do()"+m.Content+"don't()", -1)
		return matches
	}

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, subsequence := range getSubsequences() {
		matches := regex.FindAllStringSubmatch(subsequence, -1)
		for _, match := range matches {
			if len(match) < 3 {
				continue
			}
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, err
			}
			muls = append(muls, Mul{a, b})
		}
	}

	return muls, nil
}

func parseFile(filePath string) (*Memory, error) {
	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	line := strings.Join(lines, "")

	return &Memory{Content: line}, nil
}
