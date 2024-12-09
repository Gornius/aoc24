package main

import (
	"strconv"

	"github.com/gornius/aoc24/pkg/fileutils"
)

func parseFile(filePath string) (*Disk, error) {
	isFile := true
	var fileId uint64 = 0

	disk := Disk{}
	disk.Blocks = []Block{}

	lines, err := fileutils.FileToArrayOfStrings(filePath)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		for _, char := range line {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			for i := 0; i < number; i++ {
				block := Block{}
				if isFile {
					id := fileId
					block.FileId = &id
				}
				disk.Blocks = append(disk.Blocks, block)
			}
			if isFile {
				fileId++
			}
			isFile = !isFile
		}
	}

	return &disk, nil
}
