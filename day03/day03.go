package day03

import (
	"aoc2024/common"
	"fmt"
	"regexp"
	"strings"
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

func Part2(lines []string) int {
	patternMul := `mul\((\d{1,3}),(\d{1,3})\)`
	patternDo := `do\(\)`
	patternDont := `don't\(\)`
	patternCombined := fmt.Sprintf(
		`(%s)|(%s)|(%s)`,
		patternMul, patternDo, patternDont,
	)

	reCombined := regexp.MustCompile(patternCombined)

	sum := 0
	do := true
	for _, line := range lines {
		matches := reCombined.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch {
			case strings.HasPrefix(match[0], "mul("):
				if do {
					x := common.StringToInt(match[2])
					y := common.StringToInt(match[3])
					mul := x * y
					sum += mul
				}
			case strings.HasPrefix(match[0], "do("):
				do = true
			case strings.HasPrefix(match[0], "don't("):
				do = false
			default:
				panic(fmt.Sprintf("oops %s, %T", match, match))
			}
		}
	}

	return sum
}
