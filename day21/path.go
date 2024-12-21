package day21

type Path struct {
	coords []Coord
}

func NewPath(coords []Coord) Path {
	return Path{coords}
}

func (p Path) Start() Coord {
	return p.coords[0]
}

func (p Path) End() Coord {
	return p.coords[len(p.coords)-1]
}

func (p Path) Length() int {
	return len(p.coords)
}

func (p Path) ToInputString() string {
	if len(p.coords) == 0 {
		panic("empty path")
	}
	if len(p.coords) == 1 {
		return "A"
	}
	ret := ""
	for i := 0; i < len(p.coords)-1; i++ {
		current := p.coords[i]
		next := p.coords[i+1]

		if current.GetCoordInDir(UP) == next {
			ret += "^"
		} else if current.GetCoordInDir(RIGHT) == next {
			ret += ">"
		} else if current.GetCoordInDir(DOWN) == next {
			ret += "v"
		} else if current.GetCoordInDir(LEFT) == next {
			ret += "<"
		} else {
			panic("oops")
		}
	}
	ret += "A"
	return ret
}
