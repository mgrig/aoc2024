package day10

import (
	"aoc2024/common"
)

const (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

func Part12(lines []string) (int, int) {
	n := len(lines)
	grid := NewGrid(n)

	trailheads := make([]Coord, 0)
	tops := make([]Coord, 0)

	for r, line := range lines {
		for c, cell := range line {
			grid.grid[r][c] = common.StringToInt(string(cell))

			if grid.grid[r][c] == 0 {
				trailheads = append(trailheads, NewCoord(r, c))
			} else if grid.grid[r][c] == 9 {
				tops = append(tops, NewCoord(r, c))
			}
		}
	}

	reachableTops := make(map[Coord]map[Coord]struct{}) // coord -> set of reachable tops

	// tops are reachable from the top :)
	for _, top := range tops {
		reachableTops[top] = make(map[Coord]struct{})
		reachableTops[top][top] = struct{}{}
	}

	// BFS
	coordsToProcess := tops
	for len(coordsToProcess) > 0 {
		currentPos, remaining := coordsToProcess[0], coordsToProcess[1:]
		coordsToProcess = remaining

		for dir := 0; dir < 4; dir++ {
			nextPos := coordInDir(currentPos, dir)
			if !grid.IsInside(nextPos) || grid.ValueAt(nextPos) != grid.ValueAt(currentPos)-1 {
				continue
			}

			// path can continue
			_, exists := reachableTops[nextPos]
			if !exists {
				reachableTops[nextPos] = make(map[Coord]struct{})
			}
			for reachableTop := range reachableTops[currentPos] {
				reachableTops[nextPos][reachableTop] = struct{}{}
			}

			coordsToProcess = append(coordsToProcess, nextPos)
		}
	}

	part1 := 0
	for _, trailhead := range trailheads {
		part1 += len(reachableTops[trailhead])
	}

	countWaysToTop := NewGrid(n)
	coordsToProcess = tops
	for len(coordsToProcess) > 0 {
		currentPos, remaining := coordsToProcess[0], coordsToProcess[1:]
		coordsToProcess = remaining
		countWaysToTop.Increment(currentPos)

		//fmt.Println(countWaysToTop)

		for dir := 0; dir < 4; dir++ {
			nextPos := coordInDir(currentPos, dir)
			if !grid.IsInside(nextPos) || grid.ValueAt(nextPos) != grid.ValueAt(currentPos)-1 {
				continue
			}

			// path can continue
			coordsToProcess = append(coordsToProcess, nextPos)
		}
	}
	part2 := 0
	for _, trailhead := range trailheads {
		part2 += countWaysToTop.ValueAt(trailhead)
	}

	return part1, part2
}

// result can be outside the grid
func coordInDir(start Coord, dir int) Coord {
	switch dir {
	case UP:
		return NewCoord(start.r-1, start.c)
	case RIGHT:
		return NewCoord(start.r, start.c+1)
	case DOWN:
		return NewCoord(start.r+1, start.c)
	case LEFT:
		return NewCoord(start.r, start.c-1)
	default:
		panic("wrong dir")
	}
}
