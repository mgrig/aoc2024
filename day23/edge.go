package day23

type Edge struct {
	a, b string
}

func NewEdge(a, b string) Edge {
	return Edge{a, b}
}

func (e Edge) String() string {
	return e.a + "-" + e.b
}
