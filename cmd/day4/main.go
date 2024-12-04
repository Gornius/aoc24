package main

import (
	"errors"
	"fmt"

	"github.com/gornius/aoc24/pkg/fileutils"
)

type Board struct {
	Chars [][]rune
}

var (
	ErrOutOfBoardRange   = errors.New("out of board range")
	ErrBoardIsEmpty      = errors.New("board is empty")
	ErrAnchorMustBeXRune = errors.New("anchor must be 'X' rune")
	ErrAnchorMustBeARune = errors.New("anchor must be 'A' rune")
)

func (b *Board) getCharAt(x int, y int) (rune, error) {
	if len(b.Chars) == 0 {
		return '\u0000', ErrBoardIsEmpty
	}
	if x < 0 || y < 0 {
		return '\u0000', ErrOutOfBoardRange
	}
	if y > len(b.Chars)-1 {
		return '\u0000', ErrOutOfBoardRange
	}
	if x > len(b.Chars[0])-1 {
		return '\u0000', ErrOutOfBoardRange
	}
	return b.Chars[y][x], nil
}

func (b *Board) getAnchors(firstLetterOfWord rune) []Anchor {
	anchors := []Anchor{}

	for y, row := range b.Chars {
		for x, char := range row {
			if char == firstLetterOfWord {
				anchors = append(anchors, Anchor{x, y})
			}
		}
	}

	return anchors
}

type Anchor struct {
	X, Y int
}

func (a *Anchor) getXmasCount(b *Board) (int, error) {
	anchorRune, err := b.getCharAt(a.X, a.Y)
	if err != nil {
		return 0, err
	}
	if anchorRune != 'X' {
		return 0, ErrAnchorMustBeXRune
	}
	directions := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	xmas := []rune{'X', 'M', 'A', 'S'}
	count := 0

	for _, direction := range directions {
		for i := 1; i < len(xmas); i++ {
			char, err := b.getCharAt(a.X+i*direction[0], a.Y+i*direction[1])
			if err != nil {
				if errors.Is(err, ErrOutOfBoardRange) {
					break
				}
			}
			if char != xmas[i] {
				break
			}
			if i == len(xmas)-1 {
				count++
			}
		}

	}

	return count, nil
}

func (a *Anchor) isMas(b *Board) (bool, error) {
	anchorRune, err := b.getCharAt(a.X, a.Y)
	if err != nil {
		return false, err
	}
	if anchorRune != 'A' {
		return false, ErrAnchorMustBeARune
	}
	tl, err := b.getCharAt(a.X-1, a.Y+1)
	if err != nil {
		return false, nil
	}
	tr, err := b.getCharAt(a.X+1, a.Y+1)
	if err != nil {
		return false, nil
	}
	bl, err := b.getCharAt(a.X-1, a.Y-1)
	if err != nil {
		return false, nil
	}
	br, err := b.getCharAt(a.X+1, a.Y-1)
	if err != nil {
		return false, nil
	}

	firstDiagonal := string(tr) + string(bl)
	secondDiagonal := string(tl) + string(br)

	isDiagonalMas := func(str string) bool {
		return str == "MS" || str == "SM"
	}

	return isDiagonalMas(firstDiagonal) && isDiagonalMas(secondDiagonal), nil
}

func main() {
	board, err := parseFile("./challenge_data.txt")
	if err != nil {
		panic(err)
	}

	totalXmasCount := 0

	anchors := board.getAnchors('X')
	for _, anchor := range anchors {
		xmasCount, err := anchor.getXmasCount(board)
		if err != nil {
			panic(err)
		}
		totalXmasCount += xmasCount
	}

	fmt.Printf("totalXmasCount: %v\n", totalXmasCount)

	// ===================

	totalCrossMasCount := 0

	masAnchors := board.getAnchors('A')
	for _, masAnchor := range masAnchors {
		isMas, err := masAnchor.isMas(board)
		if err != nil {
			panic(err)
		}
		if isMas {
			totalCrossMasCount++
		}
	}

	fmt.Printf("totalCrossMasCount: %v\n", totalCrossMasCount)
}

func parseFile(filePath string) (*Board, error) {
	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}
	if len(lines) == 0 {
		return nil, ErrBoardIsEmpty
	}

	runes := [][]rune{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		runes = append(runes, []rune(line))
	}

	board := &Board{
		Chars: runes,
	}

	return board, nil
}
