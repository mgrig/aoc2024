package day06

const (
	EMPTY    int = int('.')
	OBSTACLE int = int('#')
	GUARD_N  int = int('^')

	UP    int = 0
	DOWN  int = 1
	LEFT  int = 2
	RIGHT int = 3
)

func Part1(lines []string) int {
	n := len(lines)
	g := NewGrid(n)

	var guard Guard
	for r, line := range lines {
		for c, val := range line {
			g.grid[r][c] = int(val)
			if int(val) == GUARD_N {
				guard = NewGuard(NewCoord(r, c), UP)
			}
		}
	}

	walkedCoords := make(map[Coord]struct{})
	for g.IsInside(guard.coord) {
		walkedCoords[guard.coord] = struct{}{}
		guard = MoveGuard(g, guard)
	}

	return len(walkedCoords)
}

func MoveGuard(g *Grid, guard Guard) (newGuard Guard) {
	tryNewPos := guard.coord.GetCoordInDir(guard.dir)
	if g.IsObstacle(tryNewPos) {
		// rotate right
		return guard.TurnRight()
	}
	return NewGuard(tryNewPos, guard.dir)
}
