package main

import (
	"strconv"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func parseFile(filePath string) (*Board, error) {
	board := &Board{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	for y, line := range lines {
		if line == "" {
			continue
		}
		row := []Field{}
		for x, char := range line {
			intVal, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row = append(row, Field{
				X:      x,
				Y:      y,
				Height: intVal,
			})
		}
		board.Fields = append(board.Fields, row)
	}

	return board, nil
}
