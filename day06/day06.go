package day06

import "fmt"

const (
	OBSTACLE int = int('#')
	GUARD_N  int = int('^')

	UP    int = 0
	DOWN  int = 1
	LEFT  int = 2
	RIGHT int = 3
)

func Part1(lines []string) int {
	g, guard := ParseGridFromLines(lines)

	walkedCoords := make(map[Coord]struct{})
	for g.IsInside(guard.coord) {
		walkedCoords[guard.coord] = struct{}{}
		guard = MoveGuard(g, guard)
	}

	return len(walkedCoords)
}

func Part2(lines []string) int {
	g, guard := ParseGridFromLines(lines)
	n := g.getN()

	countSolutions := 0
	for r := 0; r < n; r++ {
		fmt.Println(r, countSolutions)
		for c := 0; c < n; c++ {
			if r == guard.coord.r && c == guard.coord.c {
				continue
			}
			// try obstacle here
			newG := g.Clone()
			newG.grid[r][c] = OBSTACLE
			if WalksInLoop(newG, guard) {
				countSolutions++
			}
		}
	}

	return countSolutions
}

func WalksInLoop(g *Grid, guard Guard) bool {
	guardPath := make(map[Guard]struct{})
	guardPath[guard] = struct{}{}

	for g.IsInside(guard.coord) {
		guardPath[guard] = struct{}{}
		guard = MoveGuard(g, guard)
		if _, exists := guardPath[guard]; exists {
			return true
		}
	}

	return false
}

func MoveGuard(g *Grid, guard Guard) (newGuard Guard) {
	tryNewPos := guard.coord.GetCoordInDir(guard.dir)
	if g.IsObstacle(tryNewPos) {
		// rotate right
		return guard.TurnRight()
	}
	return NewGuard(tryNewPos, guard.dir)
}
