package day15

import (
	"fmt"
)

const (
	BOX           int = 'O'
	ROBOT         int = '@'
	WALL          int = '#'
	EMPTY         int = '.'
	BIG_BOX_LEFT  int = '['
	BIG_BOX_RIGHT int = ']'

	UP    int = '^'
	RIGHT int = '>'
	DOWN  int = 'v'
	LEFT  int = '<'
)

func dirAsString(dir int) string {
	switch dir {
	case UP:
		return "^"
	case LEFT:
		return "<"
	case RIGHT:
		return ">"
	case DOWN:
		return "v"
	default:
		panic("invalid direction")
	}
}

func Part1(lines []string) int {
	grid, robot, directions := ParseInput(lines)

	for _, dir32 := range directions {
		dir := int(dir32)

		nextCoord := robot.GetCoordInDir(dir)
		switch grid.ValueAt(nextCoord) {
		case WALL:
			continue
		case EMPTY:
			robot = nextCoord
		case BOX:
			pushed := grid.Push(nextCoord, dir)
			if pushed {
				robot = nextCoord
			}
		}
	}

	return grid.EncodeBoxCoords(BOX)
}

func Part2(lines []string) int {
	grid, robot, directions := ParseInput(lines)

	// widen grid
	n := len(grid.grid)
	wideGrid := NewGrid(n, 2*n)
	for r, row := range grid.grid {
		for c, value := range row {
			switch value {
			case WALL:
				wideGrid.grid[r][2*c] = WALL
				wideGrid.grid[r][2*c+1] = WALL
			case BOX:
				wideGrid.grid[r][2*c] = BIG_BOX_LEFT
				wideGrid.grid[r][2*c+1] = BIG_BOX_RIGHT
			case EMPTY:
				wideGrid.grid[r][2*c] = EMPTY
				wideGrid.grid[r][2*c+1] = EMPTY
			default:
				panic("oops")
			}
		}
	}
	robot = NewCoord(robot.r, robot.c*2)

	// start moving around
	//showGridAndRobot(wideGrid, robot)
	for _, dir32 := range directions {
		dir := int(dir32)
		//fmt.Printf("%d: Move %s\n", i, dirAsString(dir))

		nextCoord := robot.GetCoordInDir(dir)
		switch wideGrid.ValueAt(nextCoord) {
		case WALL: // continue
		case EMPTY:
			robot = nextCoord
		case BIG_BOX_LEFT, BIG_BOX_RIGHT:
			pushed := wideGrid.PushBigBox(nextCoord, dir)
			if pushed {
				//fmt.Println("PUSHED")
				robot = nextCoord
			}
		default:
			panic("oops")
		}

		//showGridAndRobot(wideGrid, robot)
	}

	return wideGrid.EncodeBoxCoords(BIG_BOX_LEFT)
}

func showGridAndRobot(g *Grid, robot Coord) {
	ret := ""
	for r, row := range g.grid {
		for c, val := range row {
			if robot.r == r && robot.c == c {
				ret += "@"
			} else {
				ret += fmt.Sprintf("%c", val)
			}
		}
		ret += "\n"
	}
	fmt.Println(ret)
}

func ParseInput(lines []string) (*Grid, Coord, string) {
	var robot Coord
	n := len(lines[0])
	grid := NewGrid(n, n)

	readingMap := true
	directions := ""
	for r, line := range lines {
		if line == "" {
			readingMap = false
			continue
		}

		if readingMap {
			for c, value := range line {
				if int(value) == ROBOT {
					robot = NewCoord(r, c)
					grid.grid[r][c] = EMPTY
					continue
				}
				grid.grid[r][c] = int(value)
			}
		} else {
			directions += line
		}
	}

	return grid, robot, directions
}
