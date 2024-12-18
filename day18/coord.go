package day18

import "fmt"

type Coord struct {
	x, y int // 0-based
}

func NewCoord(x, y int) Coord {
	return Coord{x: x, y: y}
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

func (coord Coord) GetCoordInDir(dir int) Coord {
	switch dir {
	case UP:
		return NewCoord(coord.x, coord.y-1)
	case DOWN:
		return NewCoord(coord.x, coord.y+1)
	case LEFT:
		return NewCoord(coord.x-1, coord.y)
	case RIGHT:
		return NewCoord(coord.x+1, coord.y)
	}
	panic("unreachable")
}
