package day13

import "fmt"

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
