package day16

import (
	"aoc2024/common"
	"fmt"
	"math"
)

const (
	START int = 'S'
	END   int = 'E'
	WALL  int = '#'
	EMPTY int = '.'

	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

func Part1(lines []string) int {
	unvisited := make(map[Node]struct{})
	distances := make(map[Node]int)

	grid := NewGrid(len(lines))
	var startCoord, endCoord Coord
	for r, line := range lines {
		for c, cell := range line {
			coord := NewCoord(r, c)
			for dir := 0; dir < 4; dir++ {
				unvisited[NewNode(coord, dir)] = struct{}{}
			}

			if int(cell) == START {
				startCoord = coord
				distances[NewNode(coord, RIGHT)] = 0 // set the starting node
			} else if int(cell) == END {
				endCoord = coord
			}
			grid.grid[r][c] = int(cell)
		}
	}
	//fmt.Println(grid, startCoord, endCoord)

	node := NewNode(startCoord, RIGHT)
	for node.coord != endCoord {
		// try to continue in the current direction
		nextCoord := node.coord.GetCoordInDir(node.orientation)
		nextNode := NewNode(nextCoord, node.orientation)
		if isUnvisited(&unvisited, nextNode) && (grid.ValueAt(nextCoord) == EMPTY || grid.ValueAt(nextCoord) == END || grid.ValueAt(nextCoord) == START) {
			distances[nextNode] = common.IntMin(getDistance(&distances, nextNode), getDistance(&distances, node)+1)
		}

		nextNode = node.TurnRight()
		if isUnvisited(&unvisited, nextNode) {
			distances[nextNode] = common.IntMin(getDistance(&distances, nextNode), getDistance(&distances, node)+1000)
		}

		nextNode = node.TurnLeft()
		if isUnvisited(&unvisited, nextNode) {
			distances[nextNode] = common.IntMin(getDistance(&distances, nextNode), getDistance(&distances, node)+1000)
		}

		delete(unvisited, node)
		node = getUnvisitedNodeWithMinDistance(&unvisited, &distances)
		if len(unvisited)%1000 == 0 {
			fmt.Println(len(unvisited))
		}
	}

	return distances[node]
}

func getUnvisitedNodeWithMinDistance(unvisited *map[Node]struct{}, distances *map[Node]int) (node Node) {
	dist := math.MaxInt
	for unv := range *unvisited {
		d := getDistance(distances, unv)
		if d < dist {
			dist = d
			node = unv
		}
	}
	return
}

func getDistance(distances *map[Node]int, node Node) int {
	value, exists := (*distances)[node]
	if !exists {
		return math.MaxInt
	}
	return value
}

func isUnvisited(unvisited *map[Node]struct{}, node Node) bool {
	_, exists := (*unvisited)[node]
	return exists
}
