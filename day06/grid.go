package day06

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

func ParseGridFromLines(lines []string) (*Grid, Guard) {
	n := len(lines)
	g := NewGrid(n)

	var guard Guard
	for r, line := range lines {
		for c, val := range line {
			g.grid[r][c] = int(val)
			if int(val) == GUARD_N {
				guard = NewGuard(NewCoord(r, c), UP)
			}
		}
	}
	return g, guard
}

func (g *Grid) Clone() (other *Grid) {
	m := len(g.grid)
	n := len(g.grid[0])
	otherGrid := make([][]int, m)
	for r := range otherGrid {
		otherGrid[r] = make([]int, n)
		copy(otherGrid[r], g.grid[r])
	}
	return &Grid{
		grid: otherGrid,
	}
}

func (g *Grid) getN() int {
	return len(g.grid)
}

func (g *Grid) IsInside(coord Coord) bool {
	n := g.getN()
	return coord.r >= 0 && coord.r < n && coord.c >= 0 && coord.c < n
}

func (g *Grid) IsObstacle(coord Coord) bool {
	if !g.IsInside(coord) {
		return false
	}
	return g.grid[coord.r][coord.c] == OBSTACLE
}
