package day07

type Operation struct {
	result   int
	operands []int
}

func (oper Operation) CanBeMadeTrue() bool {
	inter := oper.operands[0]
	opers := oper.operands[1:]

	return rec(inter, opers, oper.result)
}

func rec(inter int, opers []int, final int) bool {
	if inter > final {
		return false
	}

	if len(opers) == 1 {
		if inter+opers[0] == final || inter*opers[0] == final {
			return true
		}
		return false
	}

	first := opers[0]
	rest := opers[1:]
	if rec(inter+first, rest, final) {
		return true
	}
	if rec(inter*first, rest, final) {
		return true
	}

	return false
}
