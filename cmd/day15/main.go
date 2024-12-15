package main

import (
	"fmt"
)

func main() {
	filePath := "./challenge_data.txt"

	Part1(filePath)
}

func Part1(filePath string) {
	warehouse, robot, moves, err := ParseFile(filePath)
	if err != nil {
		panic(err)
	}

	for _, move := range moves {
		PerformMove(robot, warehouse, move)
	}

	sumOfGpsCoords := 0
	for pos, objectType := range warehouse.Objects {
		if objectType == Box {
			sumOfGpsCoords += 100*pos.Y + pos.X
		}
	}

	PrintBoard(*warehouse, *robot)
	fmt.Printf("sumOfGpsCoords: %v\n", sumOfGpsCoords)
}

func PerformMove(robot *Robot, warehouse *Warehouse, direction Direction) {
	nextWarehouseObjectCoords := Coords{robot.Pos.X + direction.ToVector().X, robot.Pos.Y + direction.ToVector().Y}
	nextWarehouseObject := warehouse.Objects[nextWarehouseObjectCoords]

	if nextWarehouseObject == Wall {
		return
	}

	if nextWarehouseObject == EmptySpace {
		robot.Pos = nextWarehouseObjectCoords
		return
	}

	var tryMoveCrate func(crateCoords Coords, direction Direction) bool
	tryMoveCrate = func(crateCoords Coords, direction Direction) bool {
		nextWarehouseObjectCoords := Coords{crateCoords.X + direction.ToVector().X, crateCoords.Y + direction.ToVector().Y}
		nextWarehouseObject := warehouse.Objects[nextWarehouseObjectCoords]
		switch nextWarehouseObject {
		case Wall:
			return false
		case EmptySpace:
			warehouse.Objects[crateCoords] = EmptySpace
			warehouse.Objects[nextWarehouseObjectCoords] = Box
			return true
		case Box:
			if tryMoveCrate(nextWarehouseObjectCoords, direction) {
				warehouse.Objects[crateCoords] = EmptySpace
				warehouse.Objects[nextWarehouseObjectCoords] = Box
				return true
			}
			return false
		}
		return false
	}

	nextTileCoords := Coords{robot.Pos.X + direction.ToVector().X, robot.Pos.Y + direction.ToVector().Y}
	if tryMoveCrate(nextTileCoords, direction) {
		robot.Pos = nextTileCoords
	}
}

func PrintBoard(warehouse Warehouse, robot Robot) {
	buf := string(make([]byte, 1024))

	for y := range warehouse.Size.Y {
		for x := range warehouse.Size.X {
			obj, ok := warehouse.Objects[Coords{x, y}]
			if !ok {
				buf += " "
				continue
			}
			if robot.Pos.X == x && robot.Pos.Y == y {
				buf += "@"
				continue
			}
			buf += string(obj)
		}
		buf += "\n"
	}

	fmt.Println(buf)
}

type Coords struct {
	X, Y int
}

type Warehouse struct {
	Objects map[Coords]WarehouseObjectType
	Size    Coords
}

type WarehouseObjectType rune

const (
	EmptySpace WarehouseObjectType = '.'
	Wall       WarehouseObjectType = '#'
	Box        WarehouseObjectType = 'O'
)

type Robot struct {
	Pos Coords
}

type Direction rune

const (
	DirectionUp    Direction = '^'
	DirectionLeft  Direction = '<'
	DirectionRight Direction = '>'
	DirectionDown  Direction = 'v'
)

func (d Direction) ToVector() Coords {
	switch d {
	case DirectionUp:
		return Coords{0, -1}
	case DirectionLeft:
		return Coords{-1, 0}
	case DirectionRight:
		return Coords{1, 0}
	case DirectionDown:
		return Coords{0, 1}
	}
	panic("Wrong conversion")
}
