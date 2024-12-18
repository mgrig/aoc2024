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

func Part12(lines []string, n int, drops int) (int, Coord) {
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

	part1, cellsOnPath, _ := dijkstra(grid)

	// part 2
	// keep dropping obstacles one by one
	var found bool
	var blocker Coord
	for step := drops; step < len(coords); step++ {
		grid.SetValueAt(coords[step], OBSTACLE)
		// iff current drop blocks the currently known path, search for a new path (full new search)
		if _, exists := cellsOnPath[coords[step]]; exists {
			_, cellsOnPath, found = dijkstra(grid)
			if !found {
				// current step leads to no path
				blocker = coords[step]
			}
		}
	}

	return part1, blocker
}

func dijkstra(grid *Grid) (minDist int, cellsOnPath map[Coord]struct{}, found bool) {
	n := len(grid.grid)

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

	parents := make(map[Coord]Coord) // keep track of one shortest path

	node := start
	for node != end && len(unvisited) > 0 {
		currentDist := distances[node]
		for dir := 0; dir < 4; dir++ {
			next := node.GetCoordInDir(dir)
			if _, exists := unvisited[next]; exists { // implies grid.IsInside()
				d := getDistance(next, &distances)
				if currentDist+1 < d {
					distances[next] = currentDist + 1
					parents[next] = node
				}
			}
		}

		delete(unvisited, node)
		node, found = getUnvisitedCoordSmallestDistance(&unvisited, &distances)
		if !found {
			break
		}
		//fmt.Println(len(unvisited))
	}
	if node != end {
		// did not find a path
		return -1, nil, false
	}

	// pack parents into a set of coordinates
	node = end
	cellsOnPath = make(map[Coord]struct{})
	for node != start {
		cellsOnPath[node] = struct{}{}
		node = parents[node]
	}

	return getDistance(end, &distances), cellsOnPath, true
}

func Part2(lines []string, n int, drops int) (ret Coord) {
	return
}

func getUnvisitedCoordSmallestDistance(unvisited *map[Coord]struct{}, distances *map[Coord]int) (ret Coord, found bool) {
	minDist := math.MaxInt
	for coord := range *unvisited {
		d := getDistance(coord, distances)
		if d < minDist {
			minDist = d
			ret = coord
		}
	}
	if minDist == math.MaxInt {
		return ret, false
	}
	return ret, true
}

func getDistance(coord Coord, distances *map[Coord]int) int {
	d, exists := (*distances)[coord]
	if !exists {
		return math.MaxInt
	}
	return d
}
