package day03

import (
	"aoc2024/common"
	"regexp"
)

func Part1(lines []string) int {
	reMul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	sum := 0
	for _, line := range lines {
		matches := reMul.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sum += (common.StringToInt(match[1]) * common.StringToInt(match[2]))
		}
	}

	return sum
}
