package common

import (
	"bufio"
	"math"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetLinesFromFile(fileName string, skipEmpty bool, trim bool) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if trim {
			line = strings.Trim(line, " ")
		}
		if skipEmpty && len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func ToIntegerValues(lines []string) []int {
	values := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(strings.Trim(line, " "))
		if err != nil {
			panic(err)
		}

		values[i] = value
	}

	return values
}

func StringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return value
}

// return -1 if value is negative, 0 if 0, 1 if positive
func IntSgn(value int) int {
	if value == 0 {
		return 0
	}
	if value > 0 {
		return 1
	}
	return -1
}

func IntAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func IntMin(values ...int) int {
	ret := math.MaxInt
	for i := range values {
		if values[i] < ret {
			ret = values[i]
		}
	}
	return ret
}

func IntMax(values ...int) int {
	ret := math.MinInt
	for i := range values {
		if values[i] > ret {
			ret = values[i]
		}
	}
	return ret
}

func GetOneRegexGroup(re *regexp.Regexp, line string) string {
	tokens := re.FindStringSubmatch(line)
	if len(tokens) != 2 {
		panic("wrong number of groups in line " + line)
	}

	return tokens[1]
}

func IntToBigInt(value int) *big.Int {
	return big.NewInt(int64(value))
}

func Uint64ToBigInt(value uint64) *big.Int {
	return new(big.Int).SetUint64(value)
}
