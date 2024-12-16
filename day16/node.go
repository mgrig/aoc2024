package day16

import "fmt"

type Node struct {
	coord       Coord
	orientation int
}

func NewNode(coord Coord, orientation int) Node {
	if orientation < 0 || orientation > 3 {
		panic(fmt.Sprintf("invalid orientation: %d", orientation))
	}
	return Node{coord, orientation}
}

func (n Node) String() string {
	return fmt.Sprintf("Node(%s,%d)", n.coord, n.orientation)
}

func (n Node) TurnRight() Node {
	newDir := (n.orientation + 1) % 4
	return NewNode(n.coord, newDir)
}

func (n Node) TurnLeft() Node {
	newDir := (n.orientation - 1 + 4) % 4
	return NewNode(n.coord, newDir)
}
