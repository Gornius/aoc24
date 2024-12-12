package main

import (
	"fmt"
	"strconv"
)

func main() {
	filePath := "./challenge_data.txt"

	fmt.Println("==== PART 1 ====")
	part1(filePath)
	fmt.Println("==== PART 2 ====")
	part2(filePath)
}

func part1(filePath string) {
	board, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	trails := board.findTrails(true)
	score := calculateTrailScore(trails)

	fmt.Printf("score: %v\n", score)
}

func part2(filePath string) {
	board, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	trails := board.findTrails(false)
	score := calculateTrailScore(trails)

	fmt.Printf("score: %v\n", score)
}

type Board struct {
	Fields [][]Field
}

func (b *Board) toString() string {
	buf := ""
	for _, row := range b.Fields {
		for _, field := range row {
			buf += strconv.Itoa(field.Height)
		}
		buf += "\n"
	}
	return buf
}

func (b *Board) getZeroHeightFields() []Field {
	fields := []Field{}

	for _, row := range b.Fields {
		for _, field := range row {
			if field.Height == 0 {
				fields = append(fields, field)
			}
		}
	}

	return fields
}

func (b *Board) getHeight() int {
	return len(b.Fields)
}

func (b *Board) getWidth() int {
	return len(b.Fields[0])
}

func (b *Board) findTrails(oneTrailPerNine bool) [][]Field {
	trails := [][]Field{}

	boardHeight := b.getHeight()
	boardWidth := b.getWidth()

	var findPathToNineField func(currentPath []Field, trails *[][]Field)
	findPathToNineField = func(currentPath []Field, trails *[][]Field) {
		currentField := currentPath[len(currentPath)-1]

		isAlreadyReached := func(fieldToCheck Field) bool {
			for _, trail := range *trails {
				for _, field := range trail {
					if field.X == fieldToCheck.X && field.Y == fieldToCheck.Y {
						return true
					}
				}
			}
			return false
		}

		if currentField.Height == 9 {
			if isAlreadyReached(currentField) && oneTrailPerNine {
				return
			}
			pathToCopy := []Field{}
			pathToCopy = append(pathToCopy, currentPath...)
			*trails = append(*trails, pathToCopy)
		}

		positionsToCheck := [][]int{
			{currentField.Y + 1, currentField.X},
			{currentField.Y - 1, currentField.X},
			{currentField.Y, currentField.X + 1},
			{currentField.Y, currentField.X - 1},
		}

		for _, positionToCheck := range positionsToCheck {
			if positionToCheck[0] < 0 || positionToCheck[1] < 0 {
				continue
			}
			if positionToCheck[0] >= boardHeight || positionToCheck[1] >= boardWidth {
				continue
			}
			nextField := b.Fields[positionToCheck[0]][positionToCheck[1]]
			if nextField.Height != currentField.Height+1 {
				continue
			}
			currentPath = append(currentPath, nextField)
			findPathToNineField(currentPath, trails)
		}
	}

	for _, field := range b.getZeroHeightFields() {
		foundTrails := [][]Field{}
		findPathToNineField([]Field{field}, &foundTrails)
		trails = append(trails, foundTrails...)
	}

	return trails
}

type Field struct {
	X, Y   int
	Height int
}

func calculateTrailScore(trails [][]Field) int {
	return len(trails)
}
