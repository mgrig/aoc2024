package day20

const (
	START int = 'S'
	END   int = 'E'
	WALL  int = '#'
	EMPTY int = '.'
)

func Part1(lines []string) int {
	grid := NewGrid(len(lines))
	var start, end Coord
	for r, line := range lines {
		for c, v32 := range line {
			v := int(v32)
			if v == START {
				start = NewCoord(r, c)
				grid.SetValueAt(start, EMPTY)
			} else if v == END {
				end = NewCoord(r, c)
				grid.SetValueAt(end, EMPTY)
			} else {
				grid.grid[r][c] = v
			}
		}
	}

	path := findSinglePath(grid, start, end)

	distFromStart := make(map[Coord]int, len(path))
	for i, coord := range path {
		distFromStart[coord] = i
	}

	savesMap := make(map[int]int) // number of saves -> number of cheats
	for startToHere, coord := range path {

		// find possible cheats from here
		cheatCoords := []Coord{
			coord.GetCoordInDir(RIGHT).GetCoordInDir(RIGHT),
			coord.GetCoordInDir(RIGHT).GetCoordInDir(UP),
			coord.GetCoordInDir(RIGHT).GetCoordInDir(DOWN),
			coord.GetCoordInDir(DOWN).GetCoordInDir(DOWN),
			coord.GetCoordInDir(LEFT).GetCoordInDir(LEFT),
			coord.GetCoordInDir(LEFT).GetCoordInDir(DOWN),
			coord.GetCoordInDir(LEFT).GetCoordInDir(UP),
			coord.GetCoordInDir(UP).GetCoordInDir(UP),
		}
		for _, cheatTo := range cheatCoords {
			d, exists := distFromStart[cheatTo]
			if exists && startToHere+2 < d {
				saved := d - (startToHere + 2)
				_, exists = savesMap[saved]
				if !exists {
					savesMap[saved] = 0
				}
				savesMap[saved]++
			}
		}
	}
	//fmt.Println(savesMap)

	atLeast100 := 0
	for nrSaves, nrCheats := range savesMap {
		if nrSaves >= 100 {
			atLeast100 += nrCheats
		}
	}

	return atLeast100
}

func findSinglePath(grid *Grid, start, end Coord) []Coord {
	path := make([]Coord, 0)

	visited := make(map[Coord]struct{}) // same elements as in path, but faster to check for .contains()
	coord := start
	for {
		path = append(path, coord)
		visited[coord] = struct{}{}

		if coord == end {
			break
		}

		var next Coord
		found := false
		for dir := 0; dir < 4; dir++ {
			next = coord.GetCoordInDir(dir)
			if grid.ValueAt(next) == EMPTY && !contains(&visited, next) {
				found = true
				break
			}
		}
		if !found {
			panic("next not found")
		}
		coord = next
	}

	return path
}

func contains(visited *map[Coord]struct{}, coord Coord) bool {
	_, exists := (*visited)[coord]
	return exists
}