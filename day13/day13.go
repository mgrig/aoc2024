package day13

import (
	"aoc2024/common"
	"regexp"
)

func Part1(lines []string) int {
	regA := regexp.MustCompile(`Button A: X\+(.+?), Y\+(.+)`)
	regB := regexp.MustCompile(`Button B: X\+(.+?), Y\+(.+)`)
	regT := regexp.MustCompile(`Prize: X=(.+?), Y=(.+)`)

	sum := 0
	for i := 0; i < len(lines); i += 3 {
		matchesA := regA.FindStringSubmatch(lines[i])
		matchesB := regB.FindStringSubmatch(lines[i+1])
		matchesT := regT.FindStringSubmatch(lines[i+2])
		//fmt.Println(matchesA[1], matchesA[2], matchesB[1], matchesB[2], matchesT[1], matchesT[2])

		machine := NewMachine(
			NewPoint(common.StringToInt(matchesA[1]), common.StringToInt(matchesA[2])),
			NewPoint(common.StringToInt(matchesB[1]), common.StringToInt(matchesB[2])),
			NewPoint(common.StringToInt(matchesT[1]), common.StringToInt(matchesT[2])),
		)
		//fmt.Println(machine.Solve())
		na, nb, solved := machine.Solve()
		if solved && na <= 100 && nb <= 100 {
			sum += 3*na + nb
		}
	}

	return sum
}

func Part2(lines []string) int {
	regA := regexp.MustCompile(`Button A: X\+(.+?), Y\+(.+)`)
	regB := regexp.MustCompile(`Button B: X\+(.+?), Y\+(.+)`)
	regT := regexp.MustCompile(`Prize: X=(.+?), Y=(.+)`)

	sum := 0
	for i := 0; i < len(lines); i += 3 {
		matchesA := regA.FindStringSubmatch(lines[i])
		matchesB := regB.FindStringSubmatch(lines[i+1])
		matchesT := regT.FindStringSubmatch(lines[i+2])
		//fmt.Println(matchesA[1], matchesA[2], matchesB[1], matchesB[2], matchesT[1], matchesT[2])

		machine := NewMachine(
			NewPoint(common.StringToInt(matchesA[1]), common.StringToInt(matchesA[2])),
			NewPoint(common.StringToInt(matchesB[1]), common.StringToInt(matchesB[2])),
			NewPoint(common.StringToInt(matchesT[1])+10000000000000, common.StringToInt(matchesT[2])+10000000000000),
		)
		//fmt.Println(machine.Solve())
		na, nb, solved := machine.Solve()
		if solved {
			sum += 3*na + nb
		}
	}

	return sum
}
