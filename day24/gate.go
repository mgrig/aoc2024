package day24

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
