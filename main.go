package main

import (
	"aoc2024/common"
	"aoc2024/day10"
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

	//// day 7
	//lines := common.GetLinesFromFile("day07/07.txt", true, true)
	//part1, part2 := day07.Part12(lines)
	//fmt.Println(part1)
	//fmt.Println(part2)

	//// day 8
	//lines := common.GetLinesFromFile("day08/08.txt", true, true)
	//part1 := day08.Part1(lines)
	//part2 := day08.Part2(lines)
	//fmt.Println(part1)
	//fmt.Println(part2)

	//// day 9
	//lines := common.GetLinesFromFile("day09/09.txt", true, true)
	//part1 := day09.Part1(lines)
	//part2 := day09.Part2(lines)
	//fmt.Println(part1)
	//fmt.Println(part2)

	// day 10
	lines := common.GetLinesFromFile("day10/10.txt", true, true)
	part1, part2 := day10.Part12(lines)
	fmt.Println(part1, part2)
}
