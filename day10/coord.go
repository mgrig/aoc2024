package day10

import "fmt"

type Coord struct {
	r, c int // 0-based
}

func NewCoord(r, c int) Coord {
	return Coord{r: r, c: c}
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.r, c.c)
}
