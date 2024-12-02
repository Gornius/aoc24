package fileutils

import (
	"os"
	"strings"
)

func FileToArrayOfStrings(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	for _, line := range strings.Split(string(file), "\n") {
		lines = append(lines, line)
	}

	return lines, nil
}
