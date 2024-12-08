package main

import (
	"fmt"

	"github.com/gornius/aoc24/pkg/arrayutils"
)

type Coords struct {
	X, Y int
}

type BoardField struct {
	X, Y       int
	IsAntinode bool
	Frequency  rune
}

func (f *BoardField) MarkAntinodes(other *BoardField, board *Board) {
	diff := Coords{
		X: other.X - f.X,
		Y: other.Y - f.Y,
	}

	firstAntinodeCoords := Coords{
		X: f.X + 2*diff.X,
		Y: f.Y + 2*diff.Y,
	}

	secondAntinodeCoords := Coords{
		X: other.X - 2*diff.X,
		Y: other.Y - 2*diff.Y,
	}

	board.MarkAntinode(firstAntinodeCoords)
	board.MarkAntinode(secondAntinodeCoords)
}

func (f *BoardField) MarkAntinodesWithHarmonics(other *BoardField, board *Board) {
	diff := Coords{
		X: other.X - f.X,
		Y: other.Y - f.Y,
	}

	nextX := f.X
	nextY := f.Y
	for {
		nextX = nextX - diff.X
		nextY = nextY - diff.Y

		success := board.MarkAntinode(Coords{nextX, nextY})
		if !success {
			break
		}
	}

	nextX = f.X
	nextY = f.Y
	for {
		nextX = nextX + diff.X
		nextY = nextY + diff.Y

		success := board.MarkAntinode(Coords{nextX, nextY})
		if !success {
			break
		}
	}

	board.MarkAntinode(Coords{f.X, f.Y})
}

type Board struct {
	Fields [][]BoardField
}

func (b *Board) render() string {
	buf := ""
	for _, row := range b.Fields {
		for _, field := range row {
			if field.IsAntinode && field.Frequency == '.' {
				buf += "#"
			} else {
				buf += (string(field.Frequency))
			}
		}
		buf += "\n"
	}
	return buf
}

func (b *Board) GetAllTowersWithFrequency(frequency rune) []BoardField {
	towers := []BoardField{}

	for _, row := range b.Fields {
		for _, field := range row {
			if field.Frequency == frequency {
				towers = append(towers, field)
			}
		}
	}

	return towers
}

func (b *Board) GetAllCompatibleTowerPairs() ([][]BoardField, error) {
	pairs := [][]BoardField{}

	for _, frequency := range AllPossibleFrequencies {
		towers := b.GetAllTowersWithFrequency(frequency)
		if len(towers) == 0 {
			continue
		}
		freqPairs, err := arrayutils.GenerateCombinations(towers, 2)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, freqPairs...)
	}

	return pairs, nil
}

func (b *Board) MarkAntinode(coords Coords) bool {
	boardHeight := len(b.Fields)
	boardWidth := len(b.Fields[0])

	if coords.X < 0 || coords.X >= boardWidth {
		return false
	}

	if coords.Y < 0 || coords.Y >= boardHeight {
		return false
	}

	b.GetFieldAt(coords.X, coords.Y).IsAntinode = true

	return true
}

func (b *Board) GetFieldAt(x, y int) *BoardField {
	return &b.Fields[y][x]
}

func (b *Board) GetAllAntinodesCount() int {
	antinodes := 0

	for _, row := range b.Fields {
		for _, field := range row {
			if field.IsAntinode {
				antinodes++
			}
		}
	}

	return antinodes
}

const AllPossibleFrequencies = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func part1(filePath string) {
	board, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	compatibleTowerPairs, err := board.GetAllCompatibleTowerPairs()
	if err != nil {
		panic(err)
	}

	for _, pair := range compatibleTowerPairs {
		tower := pair[0]
		otherTower := pair[1]

		tower.MarkAntinodes(&otherTower, board)
	}

	fmt.Print(board.render())

	antinodesCount := board.GetAllAntinodesCount()
	fmt.Printf("antinodesCount: %v\n", antinodesCount)
}

func part2(filePath string) {
	board, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	compatibleTowerPairs, err := board.GetAllCompatibleTowerPairs()
	if err != nil {
		panic(err)
	}

	for _, pair := range compatibleTowerPairs {
		tower := pair[0]
		otherTower := pair[1]

		tower.MarkAntinodesWithHarmonics(&otherTower, board)
	}

	fmt.Print(board.render())

	antinodesCount := board.GetAllAntinodesCount()
	fmt.Printf("antinodesCount: %v\n", antinodesCount)
}

func main() {
	filePath := "./challenge_data.txt"

	fmt.Println("=== PART 1 ===")
	part1(filePath)
	fmt.Println("=== PART 2 ===")
	part2(filePath)
}
