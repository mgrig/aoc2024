package day12

const (
	UP    int = 0
	DOWN  int = 1
	LEFT  int = 2
	RIGHT int = 3
)

func Part1(lines []string) (int, int) {
	n := len(lines)
	grid := NewGrid(n)

	for r, line := range lines {
		for c, cell := range line {
			grid.grid[r][c] = int(cell)
		}
	}

	regions := NewGrid(n)

	totalCost := 0
	totalCostSides := 0
	regionId := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			// if not yet assigned to a region
			if regions.grid[r][c] == 0 {
				toProcess := make([]Coord, 0)
				toProcess = append(toProcess, NewCoord(r, c))

				regionId++

				cellValue := grid.grid[r][c]
				regionCells := make(map[Coord]struct{})
				// parse new region (assign new id)
				for len(toProcess) > 0 {
					pos := toProcess[0]
					toProcess = toProcess[1:]

					if regions.ValueAt(pos) == regionId {
						continue
					}
					regions.SetValueAt(pos, regionId)
					regionCells[pos] = struct{}{}

					for dir := 0; dir < 4; dir++ {
						nextPos := pos.GetCoordInDir(dir)
						if !grid.IsInside(nextPos) || grid.ValueAt(nextPos) != cellValue || regions.ValueAt(nextPos) > 0 {
							continue
						}
						toProcess = append(toProcess, nextPos)
					}
				}

				area := len(regionCells)

				edges := make(map[*Edge]struct{})
				for cellPos := range regionCells {
					for dir := 0; dir < 4; dir++ {
						nextPos := cellPos.GetCoordInDir(dir)
						if _, exists := regionCells[nextPos]; !exists {
							var newEdge *Edge
							switch dir {
							case UP:
								newEdge = NewEdgeOriginDir(cellPos, RIGHT)
							case DOWN:
								newEdge = NewEdgeOriginDir(NewCoord(cellPos.r+1, cellPos.c), RIGHT)
							case LEFT:
								newEdge = NewEdgeOriginDir(cellPos, DOWN)
							case RIGHT:
								newEdge = NewEdgeOriginDir(NewCoord(cellPos.r, cellPos.c+1), DOWN)
							default:
								panic("wrong dir")
							}
							edges[newEdge] = struct{}{}
						}
					}
				}
				perimeter := len(edges)

				// compute sides from edges
				for true {
					dirty := false
					for edge := range edges {
						otherEdge, count := getOtherEdges(&edges, edge, edge.c1)
						if count > 1 {
							continue
						}
						if Aligned(edge, otherEdge) {
							// merge into 1 edge
							delete(edges, edge)
							delete(edges, otherEdge)
							merged := NewEdge(edge.c2, otherEdge.TheOtherEnd(edge.c1))
							edges[merged] = struct{}{}
							dirty = true
							break
						}

						otherEdge, count = getOtherEdges(&edges, edge, edge.c2)
						if count > 1 {
							continue
						}
						if Aligned(edge, otherEdge) {
							// merge into 1 edge
							delete(edges, edge)
							delete(edges, otherEdge)
							edges[NewEdge(edge.c1, otherEdge.TheOtherEnd(edge.c2))] = struct{}{}
							dirty = true
							break
						}
					}
					if !dirty {
						break
					}
				}
				sides := len(edges)

				//fmt.Println(regionId, "area:", area, "perimeter:", perimeter, "sides:", sides)
				totalCost += area * perimeter
				totalCostSides += area * sides
			}
		}
	}

	//fmt.Println(regions.String())

	return totalCost, totalCostSides
}

func Aligned(edge1, edge2 *Edge) bool {
	return (edge1.IsHorizontal() && edge2.IsHorizontal()) || (!edge1.IsHorizontal() && !edge2.IsHorizontal())
}

func getOtherEdges(edges *map[*Edge]struct{}, currentEdge *Edge, commonCoord Coord) (ret *Edge, count int) {
	for e := range *edges {
		if e.Contains(commonCoord) && e != currentEdge {
			count++
			ret = e
		}
	}
	if count > 1 {
		return nil, count
	}
	return
}
