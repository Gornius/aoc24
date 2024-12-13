package main

import (
	"regexp"
	"strconv"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func ParseFile(filePath string) ([]Machine, error) {
	machines := []Machine{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	buttonRegex := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	parseRegex := func(line string, regex *regexp.Regexp) (uint64, uint64, error) {
		parts := regex.FindStringSubmatch(line)
		a, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		b, err := strconv.ParseUint(parts[2], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		return a, b, nil
	}

	machinesAmount := len(lines) / 4
	for i := range machinesAmount {
		aX, aY, err := parseRegex(lines[i*4], buttonRegex)
		if err != nil {
			return nil, err
		}
		bX, bY, err := parseRegex(lines[i*4+1], buttonRegex)
		if err != nil {
			return nil, err
		}
		pX, pY, err := parseRegex(lines[i*4+2], prizeRegex)
		if err != nil {
			return nil, err
		}
		machines = append(machines, Machine{
			A: Params{aX, aY},
			B: Params{bX, bY},
			P: Params{pX, pY},
		})
	}

	return machines, nil
}
