package day21

import (
	"aoc2024/common"
	"math"
)

const (
	_A int = 'A'
	_0 int = '0'
	_1 int = '1'
	_2 int = '2'
	_3 int = '3'
	_4 int = '4'
	_5 int = '5'
	_6 int = '6'
	_7 int = '7'
	_8 int = '8'
	_9 int = '9'

	_UP    int = '^'
	_RIGHT int = '>'
	_DOWN  int = 'v'
	_LEFT  int = '<'
)

type Grid struct {
	cells map[Coord]int
	start Coord
}

func NewGrid(start Coord) *Grid {
	g := &Grid{
		cells: make(map[Coord]int),
		start: start,
	}
	g.SetValue(start, _A)
	return g
}

func (g *Grid) Inside(pos Coord) bool {
	_, exists := g.cells[pos]
	return exists
}

func (g *Grid) At(pos Coord) int {
	val, exists := g.cells[pos]
	if !exists {
		panic("no cell")
	}
	return val
}

func (g *Grid) SetValue(pos Coord, value int) {
	g.cells[pos] = value
}

func (g *Grid) GetCoord(value int) Coord {
	for pos, val := range g.cells {
		if val == value {
			return pos
		}
	}
	panic("no cell for value")
}

func (g *Grid) GetStart() Coord {
	return g.start
}

func (g *Grid) GetManhattanDistance(from, to Coord) int {
	return common.IntAbs(from.r-to.r) + common.IntAbs(from.c-to.c)
}

func (g *Grid) ShortestPaths(from, to Coord) []Path {
	return g.dijkstra(from, to)
}

func (g *Grid) dijkstra(from, to Coord) []Path {
	// create set of unvisited nodes and
	// keep track of parents (which way we reached a node on shortest path)
	// assign distance from start to every node (Inf, expect start)
	unvisited := make(map[Coord]struct{})
	parents := make(map[Coord][]Coord)
	distances := make(map[Coord]int)
	for pos := range g.cells {
		unvisited[pos] = struct{}{}
		parents[pos] = make([]Coord, 0)
		if pos == from {
			distances[pos] = 0
		} else {
			distances[pos] = math.MaxInt
		}
	}

	for {
		// select next node: unvisited with shortest distance from start
		pos, found := getClosestUnvisited(&unvisited, &distances)
		if !found {
			break
		}

		// for current node: visit all neighbors and update their distances (through current node)
		distToCurrent := distances[pos]
		for dir := 0; dir < 4; dir++ {
			next := pos.GetCoordInDir(dir)
			if g.Inside(next) {
				d, _ := distances[next]
				if distToCurrent+1 < d {
					distances[next] = distToCurrent + 1
					parents[next] = []Coord{pos}
				} else if distToCurrent+1 == d {
					parents[next] = append(parents[next], pos)
				}
			}
		}

		// remove current node from unvisited
		delete(unvisited, pos)
	}

	//find all shortest paths
	paths := allPathsTo(&parents, to)

	return paths
}

// allPathsTo returns all paths that end at `to` in forward order.
// We pass *map[Coord][]Coord to show how to use a pointer to a map,
// although maps in Go are already reference types.
func allPathsTo(parents *map[Coord][]Coord, to Coord) []Path {
	// Check if `to` has any parents in the map.
	ps, ok := (*parents)[to]

	// If it doesn't exist in the map or the slice is empty,
	// then `to` itself must be a root (no parents).
	if !ok || len(ps) == 0 {
		// Return a single path that contains only [to].
		return []Path{
			{coords: []Coord{to}},
		}
	}

	// If `to` does have parents, we collect all paths ending in each parent,
	// then append `to` to each of those paths.
	var result []Path
	for _, p := range ps {
		subPaths := allPathsTo(parents, p)
		// Each subPath is forward-ordered: [root, ..., p]
		// so we extend it by appending `to`, making [root, ..., p, to].
		for _, sp := range subPaths {
			newCoords := make([]Coord, len(sp.coords), len(sp.coords)+1)
			copy(newCoords, sp.coords)
			newCoords = append(newCoords, to)

			result = append(result, Path{coords: newCoords})
		}
	}
	return result
}

func isUnvisited(unvisited *map[Coord]struct{}, pos Coord) bool {
	_, exists := (*unvisited)[pos]
	return exists
}

func getClosestUnvisited(unvisited *map[Coord]struct{}, distances *map[Coord]int) (closest Coord, found bool) {
	minDist := math.MaxInt
	for pos := range *unvisited {
		d, exists := (*distances)[pos]
		if !exists {
			panic("oops")
		}
		if d < minDist {
			minDist = d
			closest = pos
		}
	}
	if minDist == math.MaxInt {
		return closest, false
	}
	return closest, true
}

func NewNumpad() *Grid {
	start := Coord{3, 2}
	g := NewGrid(start)
	g.SetValue(Coord{0, 0}, _7)
	g.SetValue(Coord{0, 1}, _8)
	g.SetValue(Coord{0, 2}, _9)
	g.SetValue(Coord{1, 0}, _4)
	g.SetValue(Coord{1, 1}, _5)
	g.SetValue(Coord{1, 2}, _6)
	g.SetValue(Coord{2, 0}, _1)
	g.SetValue(Coord{2, 1}, _2)
	g.SetValue(Coord{2, 2}, _3)
	g.SetValue(Coord{3, 1}, _0)
	return g
}

func NewController() *Grid {
	start := Coord{0, 2}
	g := NewGrid(start)
	g.SetValue(Coord{0, 1}, _UP)
	g.SetValue(Coord{1, 0}, _LEFT)
	g.SetValue(Coord{1, 1}, _DOWN)
	g.SetValue(Coord{1, 2}, _RIGHT)
	return g
}
