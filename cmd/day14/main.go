package main

import (
	"fmt"
	"strconv"
)

const (
	RoomWidth  = 101
	RoomHeight = 103
)

const (
	ChristmasTreeDetectorSensitivity      = 25
	ChristmasTreeDetectorMaxNumberOfSteps = 100_000
)

func main() {
	filePath := "./challenge_data.txt"

	// Part1(filePath)
	Part2(filePath)
}

func Part1(filePath string) {
	robots, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	for range 100 {
		Simulate(robots)
	}

	PrintRoom(robots)
	fmt.Printf("CalculateSafetyFactor(robots): %v\n", CalculateSafetyFactor(robots))
}

func Part2(filePath string) {
	robots, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	for i := range ChristmasTreeDetectorMaxNumberOfSteps {
		Simulate(robots)
		if PotentiallyHasEasterEgg(robots) {
			fmt.Println("=== POTENTIAL CHRISTMAS TREE AT", i+1, "SECONDS ===")
			PrintRoom(robots)
			break
		}
	}
}

func Simulate(robots []Robot) {
	for i, robot := range robots {
		robot.Pos.X = (robot.Pos.X + robot.Velocity.X + RoomWidth) % RoomWidth
		robot.Pos.Y = (robot.Pos.Y + robot.Velocity.Y + RoomHeight) % RoomHeight
		robots[i] = robot
	}
}

func PrintRoom(robots []Robot) {
	buf := ""

	tiles := map[Coords]int{}

	for _, robot := range robots {
		tiles[robot.Pos]++
	}

	for y := range RoomHeight {
		for x := range RoomWidth {
			count, ok := tiles[Coords{x, y}]
			if !ok {
				buf += "."
				continue
			}
			buf += strconv.Itoa(count)
		}
		buf += "\n"
	}

	fmt.Println(buf)
}

func CalculateSafetyFactor(robots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	middleX := RoomWidth / 2
	middleY := RoomHeight / 2

	for _, robot := range robots {
		switch true {
		case robot.Pos.X < middleX && robot.Pos.Y < middleY:
			q1++
		case robot.Pos.X > middleX && robot.Pos.Y < middleY:
			q2++
		case robot.Pos.X > middleX && robot.Pos.Y > middleY:
			q3++
		case robot.Pos.X < middleX && robot.Pos.Y > middleY:
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func PotentiallyHasEasterEgg(robots []Robot) bool {
	tilesThatHaveAnyRobot := map[Coords]bool{}
	for _, robot := range robots {
		tilesThatHaveAnyRobot[robot.Pos] = true
	}

	checkedTiles := map[Coords]bool{}

	var helper func(groupSize *int, coords Coords)
	helper = func(groupSize *int, coords Coords) {
		if _, ok := checkedTiles[coords]; ok {
			return
		}

		if _, ok := tilesThatHaveAnyRobot[coords]; !ok {
			return
		}

		checkedTiles[coords] = true
		*groupSize++

		coordsToCheck := []Coords{
			{coords.X - 1, coords.Y},
			{coords.X + 1, coords.Y},
			{coords.X, coords.Y - 1},
			{coords.X, coords.Y + 1},
		}

		for _, coordToCheck := range coordsToCheck {
			helper(groupSize, coordToCheck)
		}
	}

	for _, robot := range robots {
		size := 0
		helper(&size, robot.Pos)
		if size >= ChristmasTreeDetectorSensitivity {
			return true
		}
	}

	return false
}

type Robot struct {
	Pos      Coords
	Velocity Coords
}

type Coords struct {
	X, Y int
}
