package main

import (
	"fmt"
	"strconv"
)

type Disk struct {
	Blocks []Block
}

type MegaBlock struct {
	StartIndex int
	Length     int
}

func (d *Disk) Compact() {
	leftPtr := 0
	rightPtr := len(d.Blocks) - 1

	for rightPtr > leftPtr {
		for d.Blocks[leftPtr].FileId != nil {
			leftPtr++
		}

		for d.Blocks[rightPtr].FileId == nil {
			rightPtr--
		}

		if rightPtr < leftPtr {
			break
		}

		d.Blocks[leftPtr].ReplaceContentWith(&d.Blocks[rightPtr])
	}
}

func (d *Disk) Render() string {
	buf := ""

	for _, block := range d.Blocks {
		if block.HasContent() {
			buf += strconv.FormatUint(*block.FileId, 10)
			continue
		}
		buf += "."
	}

	buf += "\n"

	return buf
}

func (d *Disk) GetChecksum() uint64 {
	var checksum uint64 = 0

	currentBlock := 0
	for currentBlock < len(d.Blocks) {
		if d.Blocks[currentBlock].FileId == nil {
			currentBlock++
			continue
		}
		checksum += uint64(currentBlock) * (*d.Blocks[currentBlock].FileId)

		currentBlock++
	}

	return checksum
}

func (d *Disk) Defrag() {
	rightPtr := len(d.Blocks) - 1
	for rightPtr >= 0 {
		file := d.FindFile(rightPtr)
		freeSpace := d.FindFreeSpace(file.Length, rightPtr)

		rightPtr = file.StartIndex - 1

		if freeSpace == nil {
			continue
		}

		d.SwapFileWithFreeSpace(file, freeSpace)
	}
}

func (d *Disk) FindFile(fromRight int) *MegaBlock {
	megaBlock := &MegaBlock{}

	ptr := fromRight

	for !d.Blocks[ptr].HasContent() && ptr > 0 {
		ptr--
	}

	// No file was found - impossible for AoC, but let's handle it
	if !d.Blocks[ptr].HasContent() {
		return nil
	}

	endIndex := ptr
	fileId := *d.Blocks[ptr].FileId

	for ptr > 1 && d.Blocks[ptr-1].FileId != nil && *d.Blocks[ptr-1].FileId == fileId {
		ptr--
	}

	megaBlock.StartIndex = ptr
	megaBlock.Length = endIndex - megaBlock.StartIndex + 1

	return megaBlock
}

func (d *Disk) FindFreeSpace(minSize int, maxIndex int) *MegaBlock {
	megaBlock := &MegaBlock{}

	ptr := 0
	lastIndex := len(d.Blocks) - 1
	for {
		if ptr >= lastIndex || ptr > maxIndex {
			return nil
		}
		for d.Blocks[ptr].HasContent() {
			ptr++
			if ptr > lastIndex || ptr > maxIndex {
				return nil
			}
		}

		megaBlock.StartIndex = ptr

		if ptr >= lastIndex || ptr > maxIndex {
			return nil
		}
		for !d.Blocks[ptr].HasContent() {
			ptr++
			if ptr > lastIndex || ptr > maxIndex {
				break
			}
		}

		endIndex := ptr - 1
		megaBlock.Length = endIndex - megaBlock.StartIndex + 1

		if megaBlock.Length >= minSize {
			return megaBlock
		}
	}
}

func (d *Disk) SwapFileWithFreeSpace(file *MegaBlock, freeSpace *MegaBlock) {
	for i := 0; i < file.Length; i++ {
		d.Blocks[file.StartIndex+i].ReplaceContentWith(&d.Blocks[freeSpace.StartIndex+i])
	}
}

type Block struct {
	FileId *uint64
}

func (b *Block) HasContent() bool {
	return b.FileId != nil
}

func (b *Block) ReplaceContentWith(other *Block) {
	b.FileId, other.FileId = other.FileId, b.FileId
}

func main() {
	filePath := "./challenge_data.txt"
	part1(filePath)
	part2(filePath)
}

func part1(filePath string) {
	disk, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	disk.Compact()

	checksum := disk.GetChecksum()

	fmt.Printf("checksum: %v\n", checksum)
}

func part2(filePath string) {
	disk, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	disk.Defrag()

	checksum := disk.GetChecksum()

	fmt.Printf("checksum: %v\n", checksum)
}
