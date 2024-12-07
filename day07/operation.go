package day07

import "math"

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

func (oper Operation) CanBeMadeTrue3() bool {
	inter := oper.operands[0]
	opers := oper.operands[1:]

	return rec3(inter, opers, oper.result)
}

func rec3(inter int, opers []int, final int) bool {
	if inter > final {
		return false
	}

	if len(opers) == 1 {
		if inter+opers[0] == final || inter*opers[0] == final || concatenate(inter, opers[0]) == final {
			return true
		}
		return false
	}

	first := opers[0]
	rest := opers[1:]
	if rec3(inter+first, rest, final) {
		return true
	}
	if rec3(inter*first, rest, final) {
		return true
	}
	if rec3(concatenate(inter, first), rest, final) {
		return true
	}

	return false
}

// concatenate takes two positive integers a and b,
// concatenates them, and returns the resulting integer.
// For example, concatenate(12, 34) returns 1234.
func concatenate(a, b int) int {
	// Edge case: If b is 0, simply return a * 10 + 0 = a * 10
	if b == 0 {
		return a * 10
	}

	// Determine the number of digits in b
	digits := 0
	temp := b
	for temp > 0 {
		temp /= 10
		digits++
	}

	// Calculate the factor as 10^digits
	factor := int(math.Pow(10, float64(digits)))

	// Concatenate a and b
	return a*factor + b
}
