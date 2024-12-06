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
