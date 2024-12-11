package day11

import (
	"aoc2024/common"
	"fmt"
	"strings"
)

func Part1(lines []string) int {
	tokens := strings.Split(lines[0], " ")

	stones := make([]int, len(tokens))
	for i, token := range tokens {
		stones[i] = common.StringToInt(token)
	}

	// Naive solution to part 1. Works in a few seconds, but is no-go for part 2.
	//blinks := 25
	//for blink := 0; blink < blinks; blink++ {
	//	fmt.Println(blink, len(stones))
	//	for i := 0; i < len(stones); {
	//		value := stones[i]
	//		if value == 0 {
	//			stones[i] = 1
	//			i++
	//		} else {
	//			strValue := fmt.Sprintf("%d", stones[i])
	//			if len(strValue)%2 == 0 {
	//				half := len(strValue) / 2
	//				strLeft := strValue[:half]
	//				strRight := strValue[half:]
	//				stones[i] = common.StringToInt(strLeft)
	//
	//				stones = append(stones[:i+1], append([]int{common.StringToInt(strRight)}, stones[i+1:]...)...)
	//				i += 2
	//			} else {
	//				stones[i] *= 2024
	//				i++
	//			}
	//		}
	//	}
	//}
	//return len(stones)

	cache := make(map[key]int)

	// Update value of blinks for part 1 (25)  or part 2 (75)
	blinks := 75

	sum := 0
	for _, stone := range stones {
		sum += rec(stone, blinks, &cache)
	}

	return sum
}

// rec Computes the number of resulting stones for given 'value'
// after given 'steps'
// The used cache is mutable!
func rec(value int, steps int, cache *map[key]int) int {
	if steps == 0 {
		return 1
	}

	if value == 0 {
		return withCache(1, steps-1, cache)
	}

	strValue := fmt.Sprintf("%d", value)
	if len(strValue)%2 == 0 {
		half := len(strValue) / 2
		strLeft := strValue[:half]
		strRight := strValue[half:]

		return withCache(common.StringToInt(strLeft), steps-1, cache) + withCache(common.StringToInt(strRight), steps-1, cache)
	}

	return withCache(value*2024, steps-1, cache)
}

func withCache(value int, steps int, cache *map[key]int) int {
	//if steps == 0 {
	//	fmt.Println(value)
	//}

	k := NewKey(value, steps-1)
	count, exists := (*cache)[k]
	if !exists {
		count = rec(value, steps, cache)
		(*cache)[k] = count
	}
	return count
}

type key struct {
	value, steps int
}

func NewKey(value, steps int) key {
	return key{value: value, steps: steps}
}
