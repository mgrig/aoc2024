package day08

func Part1(lines []string) int {
	n := len(lines)
	antennas := make(map[int32][]Coord) // freq -> list of antenna positions

	for r, line := range lines {
		for c, cell := range line {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Coord{r, c})
			}
		}
	}

	antiNodes := make(map[Coord]struct{})
	for _, coords := range antennas {
		for i := 0; i < len(coords)-1; i++ {
			coord1 := coords[i]
			for j := i + 1; j < len(coords); j++ {
				coord2 := coords[j]

				delta_row := coord1.r - coord2.r
				delta_col := coord1.c - coord2.c

				anti1 := NewCoord(coord1.r+delta_row, coord1.c+delta_col)
				anti2 := NewCoord(coord2.r-delta_row, coord2.c-delta_col)

				if isInside(anti1, n) {
					antiNodes[anti1] = struct{}{}
				}
				if isInside(anti2, n) {
					antiNodes[anti2] = struct{}{}
				}
			}
		}
	}

	return len(antiNodes)
}

func isInside(coord Coord, n int) bool {
	return coord.r >= 0 && coord.r < n && coord.c >= 0 && coord.c < n
}
