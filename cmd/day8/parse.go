package main

import "github.com/gornius/aoc24/pkg/fileutils"

func parseFile(filePath string) (*Board, error) {
	board := &Board{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	for y, line := range lines {
		row := []BoardField{}
		if line == "" {
			continue
		}

		for x, char := range line {
			row = append(row, BoardField{
				X:          x,
				Y:          y,
				IsAntinode: false,
				Frequency:  char,
			})
		}

		board.Fields = append(board.Fields, row)
	}

	return board, nil

}
