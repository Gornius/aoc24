package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gornius/aoc24/pkg/fileutils"
	"github.com/gornius/aoc24/pkg/mathutils"
)

const MaxUnsafeLevelDiff = 3

type Report struct {
	Levels []int
}

func main() {
	reports, err := parseData("./challenge_data.txt")
	if err != nil {
		panic(err.Error())
	}

	quantityOfSafeReports := 0
	quantityOfSafeReportsDampened := 0

	for _, report := range reports {
		if checkFirstCondition(report) && checkSecondCondition(report) {
			fmt.Printf("report: %v\n", report)
			quantityOfSafeReports++
		}
	}

	for _, report := range reports {
		for _, reportVariant := range getAllDampenedVariations(report) {
			if checkFirstCondition(reportVariant) && checkSecondCondition(reportVariant) {
				quantityOfSafeReportsDampened++
				break
			}
		}

	}

	fmt.Printf("quantityOfSafeReports: %v\n", quantityOfSafeReports)
	fmt.Printf("quantityOfSafeReportsDampened: %v\n", quantityOfSafeReportsDampened)
}

func getAllDampenedVariations(report Report) []Report {
	reports := []Report{}
	for i := 0; i < len(report.Levels); i++ {
		newLevels := []int{}
		for j, level := range report.Levels {
			if i != j {
				newLevels = append(newLevels, level)
			}
		}
		reports = append(reports, Report{Levels: newLevels})
	}
	return reports
}

func checkFirstCondition(report Report) bool {
	if report.Levels[1]-report.Levels[0] == 0 {
		return false // Edge case condition, where first two pair is equal, so we can't determine if it's decreasing or increasing, but thankfully second condition rules that out
	}
	increasing := report.Levels[1]-report.Levels[0] > 0
	if increasing {
		for i := 1; i < len(report.Levels)-1; i++ {
			if report.Levels[i+1]-report.Levels[i] <= 0 {
				return false
			}
		}
	} else {
		for i := 1; i < len(report.Levels)-1; i++ {
			if report.Levels[i+1]-report.Levels[i] >= 0 {
				return false
			}
		}
	}
	return true
}

func checkSecondCondition(report Report) bool {
	for i := 0; i < len(report.Levels)-1; i++ {
		diff := mathutils.Abs(report.Levels[i+1] - report.Levels[i])
		if diff > MaxUnsafeLevelDiff {
			return false
		}
	}
	return true
}

func parseData(filePath string) ([]Report, error) {
	reports := []Report{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	regex := regexp.MustCompile(`\s+`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		levels := []int{}
		parts := regex.Split(line, -1)
		for _, part := range parts {
			if strings.TrimSpace(part) == "" {
				continue
			}
			intVal, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			levels = append(levels, intVal)
		}
		reports = append(reports, Report{
			Levels: levels,
		})
	}

	return reports, nil
}
