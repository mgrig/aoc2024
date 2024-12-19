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
