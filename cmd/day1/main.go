package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/gornius/aoc24/pkg/fileutils"
	"github.com/gornius/aoc24/pkg/mathutils"
)

type Location struct {
	LocationID int
	Position   int
}

func main() {
	locationList1, locationList2, err := parseFile("challenge_data.txt")
	cmp := func(first, second Location) int {
		return first.LocationID - second.LocationID
	}

	slices.SortFunc(locationList1, cmp)
	slices.SortFunc(locationList2, cmp)

	if err != nil {
		panic(err.Error())
	}

	sumOfDistances := 0
	similarityScore := 0

	for i := range len(locationList1) {
		sumOfDistances += mathutils.Abs(locationList1[i].LocationID - locationList2[i].LocationID)
		similarityScore += locationList1[i].LocationID * countAppearances(locationList2, locationList1[i].LocationID)
	}

	fmt.Printf("sumOfDistances: %v\n", sumOfDistances)
	fmt.Printf("similarityScore: %v\n", similarityScore)
}

func countAppearances(locations []Location, locationID int) int {
	appearances := 0
	for _, location := range locations {
		if location.LocationID == locationID {
			appearances += 1
		}
	}
	return appearances
}

func parseFile(filePath string) ([]Location, []Location, error) {
	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, nil, err
	}

	locationList1 := []Location{}
	locationList2 := []Location{}

	regex := regexp.MustCompile(`\s+`)
	for i, line := range lines {
		if line == "" {
			continue
		}
		parts := regex.Split(line, -1)

		locationId1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}

		locationId2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		locationList1 = append(locationList1, Location{
			LocationID: locationId1,
			Position:   i,
		})

		locationList2 = append(locationList2, Location{
			LocationID: locationId2,
			Position:   i,
		})
	}
	return locationList1, locationList2, nil
}
