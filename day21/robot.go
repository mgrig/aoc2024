package day21

type RobotOnGrid struct {
	g   *Grid
	pos Coord
}

func NewRobotOnGrid(g *Grid) *RobotOnGrid {
	return &RobotOnGrid{
		g:   g,
		pos: g.GetStart(),
	}
}
