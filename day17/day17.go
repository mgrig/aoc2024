package day17

import (
	"aoc2024/common"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Part1(lines []string) string {
	regexA := regexp.MustCompile(`Register A: (\d+)`)
	regexB := regexp.MustCompile(`Register B: (\d+)`)
	regexC := regexp.MustCompile(`Register C: (\d+)`)
	regexProg := regexp.MustCompile(`Program: (.*)`)

	matches := regexA.FindStringSubmatch(lines[0])
	regA := common.StringToInt(matches[1])

	matches = regexB.FindStringSubmatch(lines[1])
	regB := common.StringToInt(matches[1])

	matches = regexC.FindStringSubmatch(lines[2])
	regC := common.StringToInt(matches[1])

	matches = regexProg.FindStringSubmatch(lines[4])
	progStr := strings.Split(matches[1], ",")
	prog := make([]int, len(progStr))
	for i, str := range progStr {
		prog[i] = common.StringToInt(str)
	}

	fmt.Println("regs:", regA, regB, regC)
	fmt.Println("prog:", prog)

	output := runProgram(&regA, &regB, &regC, prog)
	fmt.Println("output:", output)

	ret := fmt.Sprintf("%d", output[0])
	for i := 1; i < len(output); i++ {
		ret += fmt.Sprintf(",%d", output[i])
	}

	return ret
}

func runProgram(regA, regB, regC *int, prog []int) (output []int) {
	ip := 0
	output = make([]int, 0)
	var pOut *int
	for {
		if ip < 0 || ip >= len(prog) {
			//halt
			break
		}
		opcode := prog[ip]

		ip, pOut = processOp(opcode, regA, regB, regC, prog, ip)
		if pOut != nil {
			output = append(output, *pOut)
		}
	}
	return output
}

func processOp(opcode int, regA, regB, regC *int, prog []int, ip int) (newIp int, pOut *int) {
	out := -1
	switch opcode {
	case 0: // adv
		operand := readComboOperand(prog[ip+1], *regA, *regB, *regC)
		result := *regA / int2Pow(operand)
		*regA = result
		ip += 2
	case 1: // bxl
		operand := readLiteralOperand(prog[ip+1])
		result := *regB ^ operand
		*regB = result
		ip += 2
	case 2: // bst
		operand := readComboOperand(prog[ip+1], *regA, *regB, *regC)
		result := operand % 8
		*regB = result
		ip += 2
	case 3: // jnz
		if *regA == 0 {
			ip += 2
		} else {
			operand := readLiteralOperand(prog[ip+1])
			ip = operand
		}
	case 4: // bxc
		result := *regB ^ *regC
		*regB = result
		ip += 2
	case 5: // out
		operand := readComboOperand(prog[ip+1], *regA, *regB, *regC)
		result := operand % 8
		out = result
		ip += 2
	case 6: // bdv
		operand := readComboOperand(prog[ip+1], *regA, *regB, *regC)
		result := *regA / int2Pow(operand)
		*regB = result
		ip += 2
	case 7: // cdv
		operand := readComboOperand(prog[ip+1], *regA, *regB, *regC)
		result := *regA / int2Pow(operand)
		*regC = result
		ip += 2
	default:
		panic(fmt.Sprintf("wrong opcode %d at ip %d", opcode, ip))
	}
	if out != -1 {
		return ip, &out
	}
	return ip, nil
}

func readLiteralOperand(in int) (out int) {
	return in
}

func readComboOperand(in int, regA, regB, regC int) (out int) {
	if in >= 0 && in <= 3 {
		return in
	}
	if in == 4 {
		return regA
	}
	if in == 5 {
		return regB
	}
	if in == 6 {
		return regC
	}
	// 7 or other is error
	panic("wrong in operand")
}

func int2Pow(power int) int {
	return int(math.Pow(2.0, float64(power)))
}
