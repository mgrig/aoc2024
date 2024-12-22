package day22

import "aoc2024/common"

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += compute(common.StringToInt(line), 2000)
	}

	return sum
}

func compute(n int, times int) int {
	for i := 0; i < times; i++ {
		n = ((n << 6) ^ n) & 0xffffff
		n = ((n >> 5) ^ n) & 0xffffff
		n = ((n << 11) ^ n) & 0xffffff
	}
	return n
}
