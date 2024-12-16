package main

import (
	"math"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func ParseFile(filePath string) (*Maze, error) {
	maze := &Maze{
		Blocks: map[Coords]MazeBlock{},
	}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	for y, line := range lines {
		for x, char := range line {
			currentCoords := Coords{x, y}
			switch true {
			case char == rune(MazeBlockFloor) || char == 'S' || char == 'E':
				maze.Blocks[currentCoords] = MazeBlock{
					Type:          MazeBlockFloor,
					CheapestPrice: math.MaxInt,
				}
			case char == rune(MazeBlockWall):
				maze.Blocks[currentCoords] = MazeBlock{
					Type: MazeBlockWall,
				}
			}
			if char == 'S' {
				maze.Start = currentCoords
			}
			if char == 'E' {
				maze.End = currentCoords
			}
		}
	}
	return maze, nil
}
