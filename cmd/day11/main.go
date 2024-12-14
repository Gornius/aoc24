package main

import (
	"fmt"
	"strconv"
)

type Stone struct {
	Val uint64
}

const TotalNumberOfSteps int = 75

func main() {
	stones := []Stone{{6563348}, {67}, {395}, {0}, {6}, {4425}, {89567}, {739318}}
	solveProblem(stones)
}

func solveProblem(stones []Stone) {
	var total uint64 = 0
	for _, stone := range stones {
		total += GetQuantityOfStonesProduced(int(stone.Val), TotalNumberOfSteps)
	}

	fmt.Printf("total: %v\n", total)
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

type GetQuantityOfStonesProducedArgs struct {
	StoneVar, NumberOfBlinks int
}

var GetQuantityOfStonesProducedMemos = map[GetQuantityOfStonesProducedArgs]uint64{}

func GetQuantityOfStonesProduced(stoneVar int, numberOfBlinks int) uint64 {
	memo, ok := GetQuantityOfStonesProducedMemos[GetQuantityOfStonesProducedArgs{stoneVar, numberOfBlinks}]
	if !ok {
		memo = GetQuantityOfStonesProducedProcessor(stoneVar, numberOfBlinks)
		GetQuantityOfStonesProducedMemos[GetQuantityOfStonesProducedArgs{stoneVar, numberOfBlinks}] = memo
	}

	return memo
}

func GetQuantityOfStonesProducedProcessor(stoneVar int, numberOfBlinks int) uint64 {
	if numberOfBlinks == 1 {
		stoneVarStr := strconv.Itoa(stoneVar)
		if len(stoneVarStr)%2 == 0 {
			return 2
		}
		return 1
	}

	newStones := ProcessStone(Stone{uint64(stoneVar)})

	var acc uint64 = 0
	for _, newStone := range newStones {
		acc += GetQuantityOfStonesProduced(int(newStone.Val), numberOfBlinks-1)
	}
	return acc
}
