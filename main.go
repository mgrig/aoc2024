package main

import (
	"aoc2024/common"
	"aoc2024/day03"
	"fmt"
)

func main() {
	// day 1 - octave
	// day 2 - octave

	// day 3
	lines := common.GetLinesFromFile("resources/03.txt", true, true)
	part1 := day03.Part1(lines)
	fmt.Println(part1)
}
