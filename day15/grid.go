package day15

import (
	"fmt"
)

type Grid struct {
	grid [][]int
}

func NewGrid(n int) *Grid {
	g := make([][]int, n)
	for r := range g {
		g[r] = make([]int, n)
	}
	return &Grid{grid: g}
}

func (g *Grid) Push(coord Coord, dir int) (pushed bool) {
	if g.ValueAt(coord) != BOX {
		panic("invalid coord for box")
	}

	// search empty cell in given direction
	nextCoord := coord
	found := false
	for !found {
		nextCoord = nextCoord.GetCoordInDir(dir)
		switch g.ValueAt(nextCoord) {
		case WALL:
			return false
		case EMPTY:
			found = true
		case BOX:
			continue
		default:
			panic("invalid cell type")
		}
	}

	// move current box into the found empty cell
	g.SetValueAt(nextCoord, BOX)
	g.SetValueAt(coord, EMPTY)

	return true
}

func (g *Grid) EncodeBoxCoords() int {
	sum := 0

	for r, row := range g.grid {
		for c, v := range row {
			if v == BOX {
				sum += 100*r + c
			}
		}
	}

	return sum
}

func (g *Grid) IsInside(coord Coord) bool {
	return coord.r >= 0 && coord.r < len(g.grid) && coord.c >= 0 && coord.c < len(g.grid[0])
}

func (g *Grid) Fill(value int) *Grid {
	for r, row := range g.grid {
		for c := range row {
			g.grid[r][c] = value
		}
	}
	return g
}

func (g *Grid) SetValueAt(coord Coord, value int) {
	g.grid[coord.r][coord.c] = value
}

func (g *Grid) ValueAt(coord Coord) int {
	return g.grid[coord.r][coord.c]
}

func (g *Grid) Increment(coord Coord) {
	g.grid[coord.r][coord.c]++
}

func (g *Grid) String() string {
	ret := ""
	for _, row := range g.grid {
		for _, val := range row {
			ret += fmt.Sprintf("%c ", val)
		}
		ret += "\n"
	}
	return ret
}
