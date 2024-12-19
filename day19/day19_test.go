package day19

import (
	"aoc2024/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	lines := common.GetLinesFromFile("19_test.txt", false, true)
	part1 := Part1(lines)
	assert.Equal(t, 6, part1)
}

func Test2(t *testing.T) {
	towels := getTowels(common.GetLinesFromFile("19_test.txt", false, true))
	target := "brwrr"
	cache := make(map[string]int)
	arrg := CountArrangements(towels, target, &cache)
	assert.Equal(t, 2, arrg)
}

func Test3(t *testing.T) {
	towels := getTowels(common.GetLinesFromFile("19_test.txt", false, true))
	target := "bggr"
	cache := make(map[string]int)
	arrg := CountArrangements(towels, target, &cache)
	assert.Equal(t, 1, arrg)
}

func Test4(t *testing.T) {
	towels := getTowels(common.GetLinesFromFile("19_test.txt", false, true))
	target := "gbbr"
	cache := make(map[string]int)
	arrg := CountArrangements(towels, target, &cache)
	assert.Equal(t, 4, arrg)
}
