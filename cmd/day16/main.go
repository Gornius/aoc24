package main

import (
	"fmt"
	"slices"

	"github.com/gornius/aoc24/pkg/mathutils"
)

func main() {
	filePath := "./challenge_data.txt"

	part1(filePath)
}

func part1(filePath string) {
	maze, err := ParseFile(filePath)
	if err != nil {
		panic(err)
	}

	shortestPath := maze.FindCheapestPath()
	fmt.Println(len(shortestPath.Coords))
	fmt.Println(CalculatePathCost(shortestPath.Coords))
}

type MazeBlockType rune

const (
	MazeBlockFloor MazeBlockType = '.'
	MazeBlockWall  MazeBlockType = '#'
)

type MazeBlock struct {
	Type          MazeBlockType
	CheapestPrice int
}

type Maze struct {
	Blocks map[Coords]MazeBlock
	Start  Coords
	End    Coords
}

type Coords struct {
	X, Y int
}

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionUp
	DirectionRight
	DirectionDown
)

type Path struct {
	Coords    []Coords
	Cost      int
	Direction Direction
}

func (d Direction) CalculateCostToTurn(previous Direction) int {
	diff := mathutils.Abs(previous - d)
	switch diff {
	case 1, 3:
		return 1000
	case 2:
		return 2000
	}
	return 0
}

func (d *Direction) ToVector() Coords {
	switch *d {
	case DirectionLeft:
		return Coords{-1, 0}
	case DirectionUp:
		return Coords{0, -1}
	case DirectionRight:
		return Coords{1, 0}
	case DirectionDown:
		return Coords{0, 1}
	}

	panic("invalid conversion")
}

func (c Coords) GetDirectionFrom(previous Coords) Direction {
	coords := Coords{c.X - previous.X, c.Y - previous.Y}
	switch coords {
	case Coords{-1, 0}:
		return DirectionLeft
	case Coords{0, -1}:
		return DirectionUp
	case Coords{1, 0}:
		return DirectionRight
	case Coords{0, 1}:
		return DirectionDown
	}

	panic("invalid conversion")
}

func (m Maze) FindCheapestPath() Path {
	paths := make([]Path, 0, 1024)
	paths = append(paths,
		Path{
			Coords:    []Coords{m.Start},
			Cost:      0,
			Direction: DirectionRight,
		})

	availableDirections := []Direction{DirectionLeft, DirectionUp, DirectionRight, DirectionDown}
	skip := 0
	for {
		if skip >= len(paths) {
			return paths[0]
		}
		slices.SortFunc(paths, func(a Path, b Path) int {
			return a.Cost - b.Cost
		})
		currentPath := paths[skip]
		head := currentPath.Coords[len(currentPath.Coords)-1]
		if head == m.End {
			fmt.Println(currentPath.Cost)
		}
		appendedPaths := 0
		impossibleDirections := 0
		subSkip := 0
		for _, direction := range availableDirections {
			nextCoord := Coords{head.X + direction.ToVector().X, head.Y + direction.ToVector().Y}
			nextBlock := m.Blocks[nextCoord]

			if nextBlock.Type == MazeBlockWall {
				impossibleDirections++
				continue
			}

			if slices.Contains(currentPath.Coords, nextCoord) {
				impossibleDirections++
				continue
			}

			if nextBlock.Type == MazeBlockFloor {
				nextDirection := nextCoord.GetDirectionFrom(head)
				turnCost := nextDirection.CalculateCostToTurn(currentPath.Direction)
				totalCost := turnCost + 1
				if totalCost > nextBlock.CheapestPrice {
					subSkip++
					continue
				}
				m.Blocks[nextCoord] = MazeBlock{
					Type:          nextBlock.Type,
					CheapestPrice: totalCost,
				}

				if appendedPaths == 0 {
					paths[skip].Coords = append(currentPath.Coords, nextCoord)
					paths[skip].Direction = nextDirection
					paths[skip].Cost += totalCost
				} else {
					copy := Path{
						Cost:   currentPath.Cost,
						Coords: []Coords{},
					}
					copy.Coords = append(copy.Coords, currentPath.Coords...)
					copy.Coords = append(copy.Coords, nextCoord)
					copy.Direction = nextDirection
					copy.Cost += totalCost
					paths = append(paths, copy)
				}
				appendedPaths++
			}
		}
		if impossibleDirections == 4 {
			skip++
		}
		if subSkip == 4-impossibleDirections {
			skip++
		}
	}
}

func CalculatePathCost(coords []Coords) int {
	cost := 0
	currentDirection := DirectionRight
	for i := 1; i < len(coords); i++ {
		cost++
		newDirection := coords[i].GetDirectionFrom(coords[i-1])
		cost += newDirection.CalculateCostToTurn(currentDirection)
		currentDirection = newDirection
	}

	return cost
}
