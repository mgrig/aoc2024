package day21

import (
	"aoc2024/common"
	"math"
	"strings"
)

func Part2(lines []string, n int) int {
	// map from desired input string (ending target A, max length ???A) + level >> min length
	cache := NewCache()

	sum := 0
	for _, line := range lines {
		grids := make([]*Grid, n)
		robots := make([]*RobotOnGrid, n)
		for i := 0; i < n; i++ {
			g := NewController()
			grids[i] = g
			robots[i] = NewRobotOnGrid(g)
		}
		gk := NewNumpad()
		grids = append(grids, gk)
		rk := NewRobotOnGrid(gk)
		robots = append(robots, rk)

		length := rec2(&grids, &robots, len(grids)-1, line, cache)
		nr := common.StringToInt(line[:len(line)-1])
		//fmt.Printf("%d * %d\n", length, nr)
		sum += length * nr
	}

	return sum
}

func rec2(grids *[]*Grid, robots *[]*RobotOnGrid, level int, targetOutput string, cache *Cache) int {
	grid := (*grids)[level]
	robot := (*robots)[level]

	minLength := 0
	for _, ch32 := range targetOutput {
		ch := int(ch32)
		from := robot.pos
		to := grid.GetCoord(ch)
		shortestPaths := grid.ShortestPaths(from, to) // can also be cached!
		inputs := pathsToInput(shortestPaths)
		//fmt.Printf("ch: %c, inputs: %v\n", ch, inputs)
		robot.pos = to

		if level == 0 {
			minLength += len(inputs[0])
			continue
		}

		minInput := math.MaxInt
		for _, input := range inputs {
			tokens := splitWithA(input)
			sumTokens := 0
			for _, token := range tokens {
				key := NewKey(token, level)
				sumTokens += cache.GetOrCompute(key, func(key Key) int {
					return rec2(grids, robots, key.level-1, key.target, cache)
				})
			}
			if sumTokens < minInput {
				minInput = sumTokens
			}
		}
		minLength += minInput
	}

	return minLength
}

func splitWithA(input string) []string {
	// Split the input by the delimiter
	parts := strings.Split(input, "A")
	tokens := make([]string, 0, len(parts)-1)

	// Re-add the delimiter to each token (except the last empty part)
	for _, part := range parts[:len(parts)-1] {
		tokens = append(tokens, part+"A")
	}

	return tokens
}

func pathsToInput(paths []Path) []string {
	ret := make([]string, len(paths))
	for i, path := range paths {
		ret[i] = path.ToInputString()
	}
	return ret
}
