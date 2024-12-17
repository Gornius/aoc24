package main

import (
	"fmt"
	"math"

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

	graph := BuildGraph(*maze)
	shortestDistance := maze.FindShortestDistance(graph)

	fmt.Printf("shortestDistance: %v\n", shortestDistance)
}

type GraphNode struct {
	Direction        *Direction
	Visited          bool
	ShortestDistance int
	PreviousNode     *GraphNode
	Nodes            []*GraphNode
}

func BuildGraph(maze Maze) map[Coords]*GraphNode {
	graph := map[Coords]*GraphNode{}

	for coords, block := range maze.Blocks {
		if block.Type == MazeBlockWall {
			continue
		}
		graph[coords] = &GraphNode{
			Direction:        nil,
			Visited:          false,
			ShortestDistance: math.MaxInt,
			PreviousNode:     nil,
			Nodes:            []*GraphNode{},
		}
	}

	for coords, node := range graph {
		for _, dir := range AvailableDirections {
			neighborCoords := Coords{coords.X + dir.ToVector().X, coords.Y + dir.ToVector().Y}
			neighbor, ok := graph[neighborCoords]
			if !ok {
				continue
			}
			node.Nodes = append(node.Nodes, neighbor)
		}
	}

	return graph
}

func (m Maze) FindShortestDistance(graph map[Coords]*GraphNode) int {
	graph[m.Start].ShortestDistance = 0
	direction := DirectionRight
	graph[m.Start].Direction = &direction

	for {
		min := math.MaxInt
		var currentNode *GraphNode = nil
		var currentCoords *Coords = nil
		for coords, node := range graph {
			if !node.Visited && node.ShortestDistance < min {
				min = node.ShortestDistance
				currentNode = node
				currentCoords = &coords
			}
		}

		for _, direction := range AvailableDirections {
			nextCoords := Coords{currentCoords.X + direction.ToVector().X, currentCoords.Y + direction.ToVector().Y}
			nextNode, ok := graph[nextCoords]
			if !ok {
				continue
			}
			nextDirection := nextCoords.GetDirectionFrom(*currentCoords)
			distance := nextDirection.CalculateCostToTurn(*currentNode.Direction) + 1
			if currentNode.ShortestDistance+distance < nextNode.ShortestDistance {
				graph[nextCoords].Direction = &nextDirection
				graph[nextCoords].PreviousNode = currentNode
				graph[nextCoords].ShortestDistance = distance + currentNode.ShortestDistance
			}
			if nextCoords == m.End {
				return nextNode.ShortestDistance
			}
		}
		graph[*currentCoords].Visited = true
	}
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

var AvailableDirections = []Direction{
	DirectionLeft,
	DirectionUp,
	DirectionRight,
	DirectionDown,
}

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
