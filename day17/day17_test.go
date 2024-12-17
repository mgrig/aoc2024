package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	regA, regB, regC := 729, 0, 0
	prog := []int{0, 1, 5, 4, 3, 0}
	output := runProgram(&regA, &regB, &regC, prog)

	assert.Equal(t, []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, output)
}

func Test1(t *testing.T) {
	regA, regB, regC := 0, 0, 9
	prog := []int{2, 6}
	runProgram(&regA, &regB, &regC, prog)

	assert.Equal(t, 1, regB)
}

func Test2(t *testing.T) {
	regA, regB, regC := 10, 0, 0
	prog := []int{5, 0, 5, 1, 5, 4}
	output := runProgram(&regA, &regB, &regC, prog)

	assert.Equal(t, []int{0, 1, 2}, output)
}

func Test3(t *testing.T) {
	regA, regB, regC := 2024, 0, 0
	prog := []int{0, 1, 5, 4, 3, 0}
	output := runProgram(&regA, &regB, &regC, prog)

	assert.Equal(t, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, output)
	assert.Equal(t, 0, regA)
}

func Test4(t *testing.T) {
	regA, regB, regC := 0, 29, 0
	prog := []int{1, 7}
	runProgram(&regA, &regB, &regC, prog)

	assert.Equal(t, 26, regB)
}

func Test5(t *testing.T) {
	regA, regB, regC := 0, 2024, 43690
	prog := []int{4, 0}
	runProgram(&regA, &regB, &regC, prog)

	assert.Equal(t, 44354, regB)
}
