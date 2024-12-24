package day24

import "fmt"

const (
	AND int = iota
	OR
	XOR
)

var TypMap map[string]int = map[string]int{
	"AND": AND,
	"OR":  OR,
	"XOR": XOR,
}

func TypToString(typ int) string {
	switch typ {
	case AND:
		return "AND"
	case OR:
		return "OR"
	case XOR:
		return "XOR"
	default:
		panic("Invalid typ")
	}
}

type Gate struct {
	in1, in2 string
	typ      int
	out      string
}

func NewGate(in1, in2 string, typ int, out string) Gate {
	return Gate{
		in1: in1,
		in2: in2,
		typ: typ,
		out: out,
	}
}

func (g Gate) Clone() Gate {
	return Gate{
		in1: g.in1,
		in2: g.in2,
		typ: g.typ,
		out: g.out,
	}
}

func (g Gate) String() string {
	return fmt.Sprintf("%s %s %s -> %s", g.in1, TypToString(g.typ), g.in2, g.out)
}
