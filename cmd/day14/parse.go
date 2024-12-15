package main

import (
	"regexp"
	"strconv"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func parseFile(filePath string) ([]Robot, error) {
	robots := []Robot{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	robotRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := robotRegex.FindStringSubmatch(line)
		pX, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}
		pY, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}
		vX, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}
		vY, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}

		robots = append(robots, Robot{
			Pos:      Coords{pX, pY},
			Velocity: Coords{vX, vY},
		})
	}

	return robots, nil
}
