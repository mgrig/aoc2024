package day15

const (
	BOX   int = 'O'
	ROBOT int = '@'
	WALL  int = '#'
	EMPTY int = '.'

	UP    int = '^'
	RIGHT int = '>'
	DOWN  int = 'v'
	LEFT  int = '<'
)

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

	return grid.EncodeBoxCoords()
}

func ParseInput(lines []string) (*Grid, Coord, string) {
	var robot Coord
	grid := NewGrid(len(lines[0]))

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
