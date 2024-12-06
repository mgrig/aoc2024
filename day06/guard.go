package day06

type Guard struct {
	coord Coord
	dir   int // must be UP, DOWN, LEFT or RIGHT
}

func NewGuard(coord Coord, dir int) Guard {
	return Guard{
		coord: coord,
		dir:   dir,
	}
}

func (g Guard) TurnRight() Guard {
	var newDir int
	switch g.dir {
	case UP:
		newDir = RIGHT
	case DOWN:
		newDir = LEFT
	case LEFT:
		newDir = UP
	case RIGHT:
		newDir = DOWN
	}
	return NewGuard(g.coord, newDir)
}
