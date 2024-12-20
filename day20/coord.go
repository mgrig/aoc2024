package day20

import "fmt"

const (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

type Coord struct {
	r, c int // 0-based
}

func NewCoord(r, c int) Coord {
	return Coord{r: r, c: c}
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.r, c.c)
}

func (coord Coord) GetCoordInDir(dir int) Coord {
	switch dir {
	case UP:
		return NewCoord(coord.r-1, coord.c)
	case DOWN:
		return NewCoord(coord.r+1, coord.c)
	case LEFT:
		return NewCoord(coord.r, coord.c-1)
	case RIGHT:
		return NewCoord(coord.r, coord.c+1)
	}
	panic("unreachable")
}
