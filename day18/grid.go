package day18

import "fmt"

type Grid struct {
	grid [][]int
}

func NewGrid(n int) *Grid {
	g := make([][]int, n)
	for x := range g {
		g[x] = make([]int, n)
	}
	return &Grid{grid: g}
}

func (g *Grid) IsInside(coord Coord) bool {
	n := len(g.grid)
	return coord.x >= 0 && coord.x < n && coord.y >= 0 && coord.y < n
}

func (g *Grid) Fill(value int) *Grid {
	for x := range g.grid {
		for y := range g.grid[x] {
			g.grid[x][y] = value
		}
	}
	return g
}

func (g *Grid) SetValueAt(coord Coord, value int) {
	g.grid[coord.x][coord.y] = value
}

func (g *Grid) ValueAt(coord Coord) int {
	return g.grid[coord.x][coord.y]
}

func (g *Grid) String() string {
	ret := ""
	ny := len(g.grid)
	nx := len(g.grid[0])
	for iy := 0; iy < ny; iy++ {
		for ix := 0; ix < nx; ix++ {
			ret += fmt.Sprintf("%c", g.grid[ix][iy])
		}
		ret += "\n"
	}
	return ret
}
