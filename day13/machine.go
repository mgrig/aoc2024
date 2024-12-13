package day13

type Machine struct {
	a, b, target Point
}

func NewMachine(a Point, b Point, target Point) Machine {
	return Machine{a: a, b: b, target: target}
}

func (m Machine) Solve() (na, nb int, solved bool) {
	det := m.a.x*m.b.y - m.a.y*m.b.x
	if det != 0 {
		na, rest := intDivision(m.b.y*m.target.x-m.b.x*m.target.y, det)
		if rest != 0 {
			return -1, -1, false
		}

		nb, rest = intDivision(m.a.x*m.target.y-m.a.y*m.target.x, det)
		if rest != 0 {
			return -1, -1, false
		}

		return na, nb, true
	}
	panic("oops")
}

func intDivision(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
