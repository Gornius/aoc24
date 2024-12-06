package main

import "github.com/gornius/aoc24/pkg/fileutils"

func parseFile(filePath string) (*Board, *Guard, error) {
	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, nil, err
	}

	var guard Guard
	var board Board
	boardFields := [][]BoardField{}

	for y, line := range lines {
		if line == "" {
			continue
		}
		fieldRow := []BoardField{}
		for x, char := range line {
			if char == '^' {
				guard = Guard{
					Board:     &board,
					Position:  Pos{x, y},
					Direction: DirectionUp,
				}
			}
			fieldRow = append(fieldRow, BoardField{
				Position:   Pos{x, y},
				IsObstacle: char == '#',
				IsVisited:  char == '^',
			})
		}
		boardFields = append(boardFields, fieldRow)
	}

	board.Fields = boardFields

	return &board, &guard, nil
}
