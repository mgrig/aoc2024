package day16

import (
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

//func Part1(lines []string) int {
//	unvisited := make(map[Node]struct{})
//	distances := make(map[Node]int)
//
//	grid := NewGrid(len(lines))
//	var startCoord, endCoord Coord
//	for r, line := range lines {
//		for c, cell := range line {
//			coord := NewCoord(r, c)
//			for dir := 0; dir < 4; dir++ {
//				unvisited[NewNode(coord, dir)] = struct{}{}
//			}
//
//			if int(cell) == START {
//				startCoord = coord
//				distances[NewNode(coord, RIGHT)] = 0 // set the starting node
//			} else if int(cell) == END {
//				endCoord = coord
//			}
//			grid.grid[r][c] = int(cell)
//		}
//	}
//
//	node := NewNode(startCoord, RIGHT)
//	for node.coord != endCoord {
//		// try to continue in the current direction
//		nextCoord := node.coord.GetCoordInDir(node.orientation)
//		nextNode := NewNode(nextCoord, node.orientation)
//		if isUnvisited(&unvisited, nextNode) && (grid.ValueAt(nextCoord) == EMPTY || grid.ValueAt(nextCoord) == END || grid.ValueAt(nextCoord) == START) {
//			distances[nextNode] = common.IntMin(getDistance(&distances, nextNode), getDistance(&distances, node)+1)
//		}
//
//		nextNode = node.TurnRight()
//		if isUnvisited(&unvisited, nextNode) {
//			distances[nextNode] = common.IntMin(getDistance(&distances, nextNode), getDistance(&distances, node)+1000)
//		}
//
//		nextNode = node.TurnLeft()
//		if isUnvisited(&unvisited, nextNode) {
//			distances[nextNode] = common.IntMin(getDistance(&distances, nextNode), getDistance(&distances, node)+1000)
//		}
//
//		delete(unvisited, node)
//		node = getUnvisitedNodeWithMinDistance(&unvisited, &distances)
//		if len(unvisited)%1000 == 0 {
//			fmt.Println(len(unvisited))
//		}
//	}
//
//	return distances[node]
//}

func Part2(lines []string) int {
	unvisited := make(map[Node]struct{})
	distances := make(map[Node]int)
	parents := make(map[Node][]Node) // node > list of parent nodes that lead to the current shortest distance

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

	node := NewNode(startCoord, RIGHT)
	var found bool
	var endNode Node
	totalMinDist := math.MaxInt
	for len(unvisited) > 0 {
		currentDistToHere := getDistance(&distances, node)
		if currentDistToHere > totalMinDist {
			delete(unvisited, node)
			node, found = getUnvisitedNodeWithMinDistance(&unvisited, &distances)
			if !found {
				break
			}
			if len(unvisited)%100 == 0 {
				fmt.Println(len(unvisited))
			}
			continue
		}

		if node.coord == endCoord {
			totalMinDist = currentDistToHere
			endNode = node
			delete(unvisited, node)
			node, found = getUnvisitedNodeWithMinDistance(&unvisited, &distances)
			if !found {
				break
			}
			continue
		}

		// try to continue in the current direction
		nextCoord := node.coord.GetCoordInDir(node.orientation)
		nextNode := NewNode(nextCoord, node.orientation)
		if isUnvisited(&unvisited, nextNode) && (grid.ValueAt(nextCoord) == EMPTY || grid.ValueAt(nextCoord) == END || grid.ValueAt(nextCoord) == START) {
			currentDistNextNode := getDistance(&distances, nextNode)
			newDistThroughThisNode := getDistance(&distances, node) + 1
			updateState(currentDistNextNode, newDistThroughThisNode, node, nextNode, &distances, &parents)
		}

		nextNode = node.TurnRight()
		if isUnvisited(&unvisited, nextNode) {
			currentDistNextNode := getDistance(&distances, nextNode)
			newDistThroughThisNode := getDistance(&distances, node) + 1000
			updateState(currentDistNextNode, newDistThroughThisNode, node, nextNode, &distances, &parents)
		}

		nextNode = node.TurnLeft()
		if isUnvisited(&unvisited, nextNode) {
			currentDistNextNode := getDistance(&distances, nextNode)
			newDistThroughThisNode := getDistance(&distances, node) + 1000
			updateState(currentDistNextNode, newDistThroughThisNode, node, nextNode, &distances, &parents)
		}

		delete(unvisited, node)
		node, found = getUnvisitedNodeWithMinDistance(&unvisited, &distances)
		if !found {
			break
		}
		if len(unvisited)%1000 == 0 {
			fmt.Println(len(unvisited))
		}

	}

	fmt.Println("shortest dist:", distances[endNode])

	// walk the parents backwards and collect visited coords
	visitedCoords := make(map[Coord]struct{})
	toVisit := make([]Node, 0)
	toVisit = append(toVisit, endNode)
	for len(toVisit) > 0 {
		current := toVisit[0]
		rest := toVisit[1:]
		toVisit = rest

		visitedCoords[current.coord] = struct{}{}

		pars, exists := parents[current]
		if exists {
			toVisit = append(toVisit, pars...)
		}
	}

	// display final map, with cells on any shortest path marked with "O"
	ret := ""
	for r, row := range grid.grid {
		for c, val := range row {
			if _, exists := visitedCoords[NewCoord(r, c)]; exists {
				ret += "O"
			} else {
				ret += fmt.Sprintf("%c", val)
			}
		}
		ret += "\n"
	}
	fmt.Println(ret)

	return len(visitedCoords)
}

func updateState(currentDistNextNode int, newDistThroughThisNode int, node Node, nextNode Node, distances *map[Node]int, parents *map[Node][]Node) {
	if newDistThroughThisNode < currentDistNextNode {
		// new min dist. drop previous parents.
		(*distances)[nextNode] = newDistThroughThisNode
		(*parents)[nextNode] = make([]Node, 0)
		(*parents)[nextNode] = append((*parents)[nextNode], node)
	} else if newDistThroughThisNode == currentDistNextNode {
		// same min dist via a new path. add parent
		if _, exists := (*parents)[nextNode]; !exists {
			(*parents)[nextNode] = make([]Node, 0)
		}
		(*parents)[nextNode] = append((*parents)[nextNode], node)
	} else {
		// noop
	}
}

func getUnvisitedNodeWithMinDistance(unvisited *map[Node]struct{}, distances *map[Node]int) (node Node, found bool) {
	dist := math.MaxInt
	for unv := range *unvisited {
		d := getDistance(distances, unv)
		if d < dist {
			dist = d
			node = unv
		}
	}
	if dist == math.MaxInt {
		// nothing found
		return node, false
	}
	return node, true
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
