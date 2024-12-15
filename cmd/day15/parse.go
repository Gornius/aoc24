package main

import (
	"math"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func ParseFile(filepath string) (*Warehouse, *Robot, []Direction, error) {
	lines, err := fileutils.FileToArrayOfStrings(filepath)
	if err != nil {
		return nil, nil, nil, err
	}

	fileParts := [][]string{}
	currentDelimPos := 0
	for i, line := range lines {
		if line == "" {
			fileParts = append(fileParts, lines[currentDelimPos:i])
			currentDelimPos = i
		}
	}
	boardPart, movesPart := fileParts[0], fileParts[1]

	robot := Robot{}
	warehouse := &Warehouse{
		Objects: map[Coords]WarehouseObjectType{},
	}
	for y, line := range boardPart {
		for x, char := range line {
			if char == '@' {
				warehouse.Objects[Coords{x, y}] = EmptySpace
				robot.Pos = Coords{x, y}
				continue
			}
			warehouse.Objects[Coords{x, y}] = WarehouseObjectType(char)
			warehouse.Size.X = int(math.Max(float64(x), float64(warehouse.Size.X)))
			warehouse.Size.Y = int(math.Max(float64(y), float64(warehouse.Size.Y)))
		}
	}

	warehouse.Size.X++
	warehouse.Size.Y++

	moves := []Direction{}
	for _, line := range movesPart {
		for _, char := range line {
			moves = append(moves, Direction(char))
		}
	}

	return warehouse, &robot, moves, nil
}
