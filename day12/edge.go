package day12

import "fmt"

type Edge struct {
	c1, c2 Coord
}

func NewEdge(coord1, coord2 Coord) *Edge {
	return &Edge{
		c1: coord1, c2: coord2,
	}
}

func NewEdgeOriginDir(origin Coord, dir int) *Edge {
	switch dir {
	case UP:
		return NewEdge(origin, NewCoord(origin.r-1, origin.c))
	case DOWN:
		return NewEdge(origin, NewCoord(origin.r+1, origin.c))
	case LEFT:
		return NewEdge(origin, NewCoord(origin.r, origin.c-1))
	case RIGHT:
		return NewEdge(origin, NewCoord(origin.r, origin.c+1))
	default:
		panic("wrong dir")
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("Edge[%s, %s]", (*e).c1, (*e).c2)
}

func (e *Edge) IsHorizontal() bool {
	return e.c1.r == e.c2.r
}

func (e *Edge) Contains(c Coord) bool {
	return (*e).c1 == c || (*e).c2 == c
}

func (e *Edge) TheOtherEnd(c Coord) Coord {
	if (*e).c1 == c {
		return (*e).c2
	}
	if (*e).c2 == c {
		return (*e).c1
	}
	panic(fmt.Sprintf("other end not found, e:", *e, "c:", c))
}
