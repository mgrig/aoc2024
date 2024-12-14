package day14

import "fmt"

type Grid struct {
	grid [][]int
}

func NewGrid(nx, ny int) *Grid {
	g := make([][]int, nx)
	for x := range g {
		g[x] = make([]int, ny)
	}
	return &Grid{grid: g}
}

func (g *Grid) SetValueAt(p Point, value int) {
	g.grid[p.x][p.y] = value
}

func (g *Grid) ValueAt(p Point) int {
	return g.grid[p.x][p.y]
}

func (g *Grid) Increment(p Point) {
	g.grid[p.x][p.y]++
}

func (g *Grid) IsSymmetric() bool {
	nx := len(g.grid)
	ny := len(g.grid[0])
	midx := (nx - 1) / 2
	for x := 0; x < midx; x++ {
		for y := 0; y < ny; y++ {
			if g.grid[x][y] != g.grid[nx-x-1][y] {
				return false
			}
		}
	}
	return true
}

func (g *Grid) ContainsSegment(length int) bool {
	nx := len(g.grid)
	ny := len(g.grid[0])
	for x := 0; x < nx-length; x++ {
		for y := 0; y < ny; y++ {
			all := true
			for k := 0; k < length; k++ {
				if g.grid[x+k][y] == 0 {
					all = false
					break
				}
			}
			if all {
				return true
			}
		}
	}
	return false
}

func (g *Grid) String() string {
	ret := ""
	for y := 0; y < len(g.grid[0]); y++ {
		for x := 0; x < len(g.grid); x++ {
			val := g.grid[x][y]
			if val == 0 {
				ret += "."
			} else {
				ret += "#"
			}
		}
		ret += fmt.Sprintln()
	}
	return ret
}
