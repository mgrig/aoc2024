package main

import (
	"aoc2024/common"
	"aoc2024/day01"
	"fmt"
)

func main() {
	// day 1
	lines := common.GetLinesFromFile("resources/01_test.txt", true, true)
	part1 := day01.Part1(lines)
	fmt.Println(part1)
}
