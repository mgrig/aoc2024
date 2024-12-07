package main

import (
	"aoc2024/common"
	"aoc2024/day07"
	"fmt"
)

func main() {
	// day 1 - octave
	// day 2 - octave

	//// day 3
	//lines := common.GetLinesFromFile("resources/03.txt", true, true)
	//part1 := day03.Part1(lines)
	//part2 := day03.Part2(lines)
	//fmt.Println(part1, part2)

	// // day 6
	// lines := common.GetLinesFromFile("day06/06.txt", true, true)
	// part1 := day06.Part1(lines)
	// part2 := day06.Part2(lines)
	// fmt.Println(part1)
	// fmt.Println(part2)

	// day 7
	lines := common.GetLinesFromFile("day07/07.txt", true, true)
	part1, part2 := day07.Part12(lines)
	fmt.Println(part1)
	fmt.Println(part2)
}
