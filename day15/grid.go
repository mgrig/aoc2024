package day15

import (
	"aoc2024/common"
	"fmt"
)

type Grid struct {
	grid [][]int
}

func NewGrid(m, n int) *Grid {
	g := make([][]int, m)
	for r := range g {
		g[r] = make([]int, n)
	}
	return &Grid{grid: g}
}

func (g *Grid) Push(coord Coord, dir int) (pushed bool) {
	if g.ValueAt(coord) != BOX {
		panic("invalid coord for box")
	}

	emptyCell, found := g.searchEmptyCoordInGivenDirection(coord, dir)
	if !found {
		return false
	}

	// move current box into the found empty cell
	g.SetValueAt(emptyCell, BOX)
	g.SetValueAt(coord, EMPTY)

	return true
}

func (g *Grid) PushBigBox(coord Coord, dir int) (pushed bool) {
	if g.ValueAt(coord) != BIG_BOX_LEFT && g.ValueAt(coord) != BIG_BOX_RIGHT {
		panic("invalid coord for box")
	}

	if dir == LEFT || dir == RIGHT {
		// similar to simple push: find next empty cell and push each affected box once
		emptyCell, found := g.searchEmptyCoordInGivenDirection(coord, dir)
		if !found {
			return false
		}

		// move all boxes into empty
		if common.IntAbs(emptyCell.c-coord.c) < 2 {
			panic("oops")
		}
		var step int
		if dir == RIGHT {
			step = 1
			g.SetValueAt(emptyCell, BIG_BOX_RIGHT)
		} else { // is LEFT
			step = -1
			g.SetValueAt(emptyCell, BIG_BOX_LEFT)
		}
		g.SetValueAt(coord, EMPTY)

		nrSteps := common.IntAbs(emptyCell.c-coord.c) - 1
		for ik := 0; ik < nrSteps; ik++ {
			k := coord.c + step + ik*step
			if g.grid[coord.r][k] == BIG_BOX_LEFT {
				g.grid[coord.r][k] = BIG_BOX_RIGHT
			} else if g.grid[coord.r][k] == BIG_BOX_RIGHT {
				g.grid[coord.r][k] = BIG_BOX_LEFT
			} else {
				panic(fmt.Sprintf("wrong cell content grid[%d][%d]=%c", coord.r, k, g.grid[coord.r][k]))
			}
		}

		return true
	} else {
		// here we use recursion with a set of boxes that would be pushed
		boxesToPush := make(map[Coord]struct{})
		if g.rec(coord, dir, &boxesToPush) {
			// apply boxesToPush
			// - remove all boxes first...
			for box := range boxesToPush {
				g.SetValueAt(box, EMPTY)
				g.SetValueAt(box.GetCoordInDir(RIGHT), EMPTY)
			}

			// - ... then apply them in new positions
			for box := range boxesToPush {
				g.SetValueAt(box.GetCoordInDir(dir), BIG_BOX_LEFT)
				g.SetValueAt(box.GetCoordInDir(dir).GetCoordInDir(RIGHT), BIG_BOX_RIGHT)
			}

			return true
		} else {
			// push did not work
			return false
		}
	}

	return false
}

func (g *Grid) rec(coord Coord, dir int, boxesToPush *map[Coord]struct{}) (pushed bool) {
	if (g.ValueAt(coord) != BIG_BOX_LEFT && g.ValueAt(coord) != BIG_BOX_RIGHT) || (dir != UP && dir != DOWN) {
		panic("invalid coord for box")
	}

	var boxCoord Coord
	switch g.ValueAt(coord) {
	case BIG_BOX_LEFT:
		boxCoord = coord
	case BIG_BOX_RIGHT:
		boxCoord = coord.GetCoordInDir(LEFT)
	default:
		panic("invalid coord for box")
	}
	(*boxesToPush)[boxCoord] = struct{}{}

	leftCoord := boxCoord.GetCoordInDir(dir)
	rightCoord := leftCoord.GetCoordInDir(RIGHT)
	leftValue, rightValue := g.ValueAt(leftCoord), g.ValueAt(rightCoord)

	if leftValue == EMPTY && rightValue == EMPTY {
		return true
	}
	if leftValue == WALL || rightValue == WALL {
		return false
	}
	if leftValue == BIG_BOX_LEFT { // && rightValue == BIG_BOX_RIGHT
		return g.rec(leftCoord, dir, boxesToPush)
	}

	var canPushLeft, canPushRight bool
	if leftValue == EMPTY {
		canPushLeft = true
	} else if leftValue == BIG_BOX_RIGHT {
		canPushLeft = g.rec(leftCoord, dir, boxesToPush)
	}

	if rightValue == EMPTY {
		canPushRight = true
	} else if rightValue == BIG_BOX_LEFT {
		canPushRight = g.rec(rightCoord, dir, boxesToPush)
	}

	return canPushLeft && canPushRight
}

func (g *Grid) searchEmptyCoordInGivenDirection(coord Coord, dir int) (emptyCoord Coord, found bool) {
	nextCoord := coord
	found = false
	for {
		switch g.ValueAt(nextCoord) {
		case WALL:
			return NewCoord(-1, -1), false
		case EMPTY:
			return nextCoord, true
		}
		nextCoord = nextCoord.GetCoordInDir(dir)
	}
}

func (g *Grid) EncodeBoxCoords(target int) int {
	sum := 0

	for r, row := range g.grid {
		for c, v := range row {
			if v == target {
				sum += 100*r + c
			}
		}
	}

	return sum
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

func (g *Grid) Fill(value int) *Grid {
	for r, row := range g.grid {
		for c := range row {
			g.grid[r][c] = value
		}
	}
	return g
}

func (g *Grid) SetValueAt(coord Coord, value int) {
	g.grid[coord.r][coord.c] = value
}

func (g *Grid) ValueAt(coord Coord) int {
	return g.grid[coord.r][coord.c]
}

func (g *Grid) Increment(coord Coord) {
	g.grid[coord.r][coord.c]++
}

func (g *Grid) String() string {
	ret := ""
	for _, row := range g.grid {
		for _, val := range row {
			ret += fmt.Sprintf("%c ", val)
		}
		ret += "\n"
	}
	return ret
}
