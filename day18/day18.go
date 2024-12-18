package day18

import (
	"aoc2024/common"
	"math"
	"strings"
)

const (
	OBSTACLE int = '#'
	EMPTY    int = '.'

	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

func Part1(lines []string, n int, drops int) int {
	coords := make([]Coord, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, ",")
		coords[i] = NewCoord(common.StringToInt(tokens[0]), common.StringToInt(tokens[1]))
	}

	grid := NewGrid(n).Fill(EMPTY)
	for i := 0; i < drops; i++ {
		grid.SetValueAt(coords[i], OBSTACLE)
	}
	//fmt.Println(grid)

	// Dijkstra from top left to bottom right
	unvisited := make(map[Coord]struct{})
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if grid.grid[x][y] == EMPTY {
				unvisited[NewCoord(x, y)] = struct{}{}
			}
		}
	}

	start := NewCoord(0, 0)
	end := NewCoord(n-1, n-1)
	distances := make(map[Coord]int)
	distances[start] = 0

	node := start
	for node != end {
		currentDist := distances[node]
		for dir := 0; dir < 4; dir++ {
			next := node.GetCoordInDir(dir)
			if _, exists := unvisited[next]; exists { // implies grid.IsInside()
				d := getDistance(next, &distances)
				if currentDist+1 < d {
					distances[next] = currentDist + 1
				}
			}
		}

		delete(unvisited, node)
		node = getUnvisitedCoordSmallestDistance(&unvisited, &distances)
		//fmt.Println(len(unvisited))
	}

	return getDistance(end, &distances)
}

func getUnvisitedCoordSmallestDistance(unvisited *map[Coord]struct{}, distances *map[Coord]int) (ret Coord) {
	minDist := math.MaxInt
	for coord := range *unvisited {
		d := getDistance(coord, distances)
		if d < minDist {
			minDist = d
			ret = coord
		}
	}
	if minDist == math.MaxInt {
		panic("node not found")
	}
	return ret
}

func getDistance(coord Coord, distances *map[Coord]int) int {
	d, exists := (*distances)[coord]
	if !exists {
		return math.MaxInt
	}
	return d
}
