package main

import (
	"fmt"
	"iter"

	"github.com/gornius/aoc24/pkg/clone"
)

func main() {
	board, guard, err := parseFile("./test_data.txt")
	if err != nil {
		panic(err)
	}
	visitedFields := part1(board, guard)

	fmt.Printf("visitedFields: %v\n", visitedFields)

	//====================

	board, guard, err = parseFile("./challenge_data.txt")
	if err != nil {
		panic(err)
	}
	allInfiniteLoopVariations := part2(board, guard)

	fmt.Printf("allInfiniteLoopVariations: %v\n", allInfiniteLoopVariations)
}

func part1(board *Board, guard *Guard) int {
	for !guard.IsNextStepOutOfRange() {
		guard.TakeAStep()
	}

	visitedFields := 0
	for _, row := range board.Fields {
		for _, field := range row {
			if field.IsVisited {
				visitedFields++
			}
		}
	}

	return visitedFields
}

func part2(board *Board, guard *Guard) int {
	allVariations := board.GetAllVariationsWithOneAdditionalObstacle()

	infiniteLoopVariations := 0
	checkedVariations := 0
	outOfRangeVariations := 0

	for variation := range allVariations {
		fmt.Printf("checkedVariations: %v\n", checkedVariations)
		fmt.Printf("outOfRangeVariations: %v\n", outOfRangeVariations)
		fmt.Printf("infiniteLoopVariations: %v\n", infiniteLoopVariations)
		fmt.Println("--------")
		thisVariationGuard := *guard
		thisVariationGuard.Board = &variation
		thisVariationGuard.StepsTaken = []Step{}
		for !thisVariationGuard.IsNextStepOutOfRange() && !thisVariationGuard.IsInInfiniteLoop() {
			thisVariationGuard.TakeAStep()
		}
		checkedVariations++
		if thisVariationGuard.IsNextStepOutOfRange() {
			outOfRangeVariations++
			continue
		}
		infiniteLoopVariations++
	}

	return infiniteLoopVariations
}

type Direction int

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

func DirectionToVector(direction Direction) (int, int) {
	switch direction {
	case DirectionUp:
		return 0, -1
	case DirectionRight:
		return 1, 0
	case DirectionDown:
		return 0, 1
	case DirectionLeft:
		return -1, 0
	}

	panic("bad direction")
}

type Pos struct {
	X, Y int
}

type Board struct {
	Fields [][]BoardField
}

func (b *Board) GetField(x, y int) *BoardField {
	return &b.Fields[y][x]
}

func (b *Board) GetDimensions() (int, int) {
	height := len(b.Fields)
	width := len(b.Fields[0])

	return width, height
}

func (b *Board) GetAllVariationsWithOneAdditionalObstacle() iter.Seq[Board] {
	return func(yield func(Board) bool) {
		for y, row := range b.Fields {
			for x, field := range row {
				if field.IsObstacle {
					continue
				}
				if field.IsVisited {
					continue
				}
				newBoard, err := clone.GobDeepClone(*b)
				if err != nil {
					panic(err)
				}
				newBoard.GetField(x, y).IsObstacle = true
				if !yield(*newBoard) {
					return
				}
			}
		}
	}
}

func (b *Board) print() {
	for _, row := range b.Fields {
		for _, field := range row {
			if field.IsObstacle {
				fmt.Printf("#")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Printf("\n")
	}
}

type BoardField struct {
	Position   Pos
	IsObstacle bool
	IsVisited  bool
}
type Guard struct {
	Position   Pos
	Direction  Direction
	Board      *Board
	StepsTaken []Step
}

func (g *Guard) ChangeDirection() {
	g.Direction = (g.Direction + 1) % 4
}

func (g *Guard) TakeAStep() {
	for g.GetFacingField().IsObstacle {
		g.ChangeDirection()
	}

	g.Position = g.GetFacingField().Position

	g.StepsTaken = append(g.StepsTaken, Step{
		Position:  g.Position,
		Direction: g.Direction,
	})

	g.GetOccupiedField().IsVisited = true
}

func (g *Guard) GetOccupiedField() *BoardField {
	return g.Board.GetField(g.Position.X, g.Position.Y)
}

func (g *Guard) GetFacingField() *BoardField {
	nextX, nextY := DirectionToVector(g.Direction)
	return g.Board.GetField(g.Position.X+nextX, g.Position.Y+nextY)
}

func (g *Guard) IsNextStepOutOfRange() bool {
	stepX, stepY := DirectionToVector(g.Direction)
	nextX := g.Position.X + stepX
	nextY := g.Position.Y + stepY

	if nextX < 0 || nextY < 0 {
		return true
	}

	boardWidth, boardHeight := g.Board.GetDimensions()
	if nextX >= boardWidth || nextY >= boardHeight {
		return true
	}

	return false
}

func (g *Guard) IsInInfiniteLoop() bool {
	if len(g.StepsTaken) == 0 {
		return false
	}
	for _, stepTaken := range g.StepsTaken[:len(g.StepsTaken)-1] {
		if stepTaken.Position.X == g.Position.X && stepTaken.Position.Y == g.Position.Y && stepTaken.Direction == g.Direction {
			return true
		}
	}
	return false
}

type Step struct {
	Position  Pos
	Direction Direction
}
