package main

import (
	"fmt"
	"strconv"
)

type Stone struct {
	Val uint64
}

func GetMapOfNumberOfStonesWithValAfterNBlinks(val uint64) map[int]int {
	m := map[int]int{}
	stones := []Stone{{val}}

	for i := 1; i <= TotalNumberOfSteps; i++ {
		fmt.Println("-- PRECALCULATING STEP", i, "FOR", val, "--")
		stones = Blink(stones)
		m[i] = len(stones)
	}

	return m
}

var MapOfMadeStonesForStone1 map[int]int = GetMapOfNumberOfStonesWithValAfterNBlinks(1)

const TotalNumberOfSteps int = 25

func main() {
	// stones := []Stone{{125}, {17}}
	stones := []Stone{{6563348}, {67}, {395}, {0}, {6}, {4425}, {89567}, {739318}}
	// part1(stones)
	part1Eager(stones)
}

func part1(stones []Stone) {
	for i := 1; i <= TotalNumberOfSteps; i++ {
		fmt.Println("=== BLINK", i, "===")
		stones = Blink(stones)
		fmt.Printf("len(stones): %v\n", len(stones))
	}

	fmt.Printf("len(stones): %v\n", len(stones))

}

func part1Eager(stones []Stone) {
	var eagerNumberOfStones uint64 = 0

	for i := 1; i <= TotalNumberOfSteps; i++ {
		fmt.Println("=== BLINK", i, "===")
		EagerProcessResults(&stones, &eagerNumberOfStones, i)
		stones = Blink(stones)
		fmt.Printf("len(stones): %v\n", len(stones))
		fmt.Printf("eagerNumberOfStones: %v\n", eagerNumberOfStones)
		fmt.Println("Total:", len(stones)+int(eagerNumberOfStones))
	}
}

func EagerProcessResults(stones *[]Stone, quantityOfStones *uint64, currentStep int) {
	for i, stone := range *stones {
		if stone.Val == 1 {
			remainingNumberOfSteps := TotalNumberOfSteps - currentStep + 1
			*quantityOfStones += uint64(MapOfMadeStonesForStone1[remainingNumberOfSteps])
			*stones = append((*stones)[:i], (*stones)[i+1:]...)
		}
	}
}

func Blink(stones []Stone) []Stone {
	newStones := []Stone{}

	for _, stone := range stones {
		stonesFromStone := ProcessStone(stone)
		newStones = append(newStones, stonesFromStone...)
	}

	return newStones
}

func ProcessStone(stone Stone) []Stone {
	stones := []Stone{}
	stoneStr := strconv.FormatUint(stone.Val, 10)

	switch true {
	case stone.Val == 0:
		stones = append(stones, Stone{1})
	case len(stoneStr)%2 == 0:
		half := len(stoneStr) / 2
		firstStone, _ := strconv.ParseUint(stoneStr[:half], 10, 64)
		secondStone, _ := strconv.ParseUint(stoneStr[half:], 10, 64)
		stones = append(stones, Stone{firstStone}, Stone{secondStone})
	default:
		stones = append(stones, Stone{stone.Val * 2024})
	}

	return stones
}
