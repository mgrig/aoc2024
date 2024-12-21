package day21

import (
	"aoc2024/common"
	"fmt"
	"math"
)

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		g1 := NewController()
		g2 := NewController()
		gk := NewNumpad()
		grids := []*Grid{g1, g2, gk}

		r1 := NewRobotOnGrid(g1)
		r2 := NewRobotOnGrid(g2)
		rk := NewRobotOnGrid(gk)
		robots := []*RobotOnGrid{r1, r2, rk}

		shortest := rec(&grids, &robots, 2, line)
		length := len(shortest[0])
		nr := common.StringToInt(line[:len(line)-1])
		fmt.Println(length * nr)
		sum += length * nr
	}

	return sum
}

func rec(grids *[]*Grid, robots *[]*RobotOnGrid, level int, targetOutput string) []string {
	grid := (*grids)[level]
	robot := (*robots)[level]

	strings := make([]string, 0)
	for _, ch32 := range targetOutput {
		ch := int(ch32)
		from := robot.pos
		to := grid.GetCoord(ch)
		shortestPaths := grid.ShortestPaths(from, to)
		inputs := pathsToInput(shortestPaths)
		//fmt.Printf("ch: %c, inputs: %v\n", ch, inputs)

		var newStrings []string
		if len(strings) == 0 {
			newStrings = inputs
		} else {
			newStrings = make([]string, len(strings)*len(inputs))
			for iString, s := range strings {
				for iInput, input := range inputs {
					newStrings[iString*len(inputs)+iInput] = s + input
				}
			}
		}
		strings = newStrings

		(*robots)[level].pos = to
	}

	shortestStrings := keepShortest(strings)
	//fmt.Println(level, shortestStrings)

	if level == 0 {
		return shortestStrings
	}
	candidates := make([]string, 0)
	for _, candidateInput := range shortestStrings {
		candidates = append(candidates, rec(grids, robots, level-1, candidateInput)...)
	}
	candidates = keepShortest(candidates)
	return candidates
}

func pathsToInput(paths []Path) []string {
	ret := make([]string, len(paths))
	for i, path := range paths {
		ret[i] = path.ToInputString()
	}
	return ret
}

func keepShortest(strings []string) []string {
	shortestStrings := make([]string, 0)
	shortest := math.MaxInt
	for _, str := range strings {
		if len(str) < shortest {
			shortest = len(str)
			shortestStrings = []string{str}
		} else if len(str) == shortest {
			shortestStrings = append(shortestStrings, str)
		}
	}
	return shortestStrings
}
