package main

import "github.com/gornius/aoc24/pkg/fileutils"

func parseFile(filePath string) (*Garden, error) {
	garden := &Garden{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	for y, line := range lines {
		if line == "" {
			continue
		}
		row := []Plot{}
		for x, char := range line {
			row = append(row, Plot{
				Garden: garden,
				Coords: Coords{x, y},
				Plant:  char,
			})
		}
		garden.Plots = append(garden.Plots, row)
	}

	return garden, nil
}
