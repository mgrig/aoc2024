package day09

import (
	"fmt"
	"strconv"
)

func Part1(lines []string) int {

	fs := ParseFilesystem(lines)

	//fmt.Println(fs.PrettyPrint())
	fs.Compress()
	//fmt.Println(fs.PrettyPrint())

	return fs.checksum()
}

func Part2(lines []string) int {

	fs := ParseFilesystem(lines)

	//fmt.Println(fs.PrettyPrint())
	fs.CompressFiles()
	//fmt.Println(fs.PrettyPrint())

	return fs.checksum()
}

func ParseFilesystem(lines []string) *Filesystem {
	fs := NewFilesystem()
	nextIsBlock := true
	nextBlockId := 0
	currentPos := 0
	for _, char := range lines[0] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			panic(fmt.Sprintln("Error converting character to digit:", err))
		}

		if nextIsBlock {
			fs.AddBlockUnsafe(NewBlock(nextBlockId, currentPos, digit))
			currentPos += digit
			nextIsBlock = false
			nextBlockId++
		} else {
			currentPos += digit
			nextIsBlock = true
		}
	}
	return fs
}
