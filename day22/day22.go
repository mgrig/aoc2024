package day22

import (
	"aoc2024/common"
	"fmt"
)

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += compute(common.StringToInt(line), 2000)
	}

	return sum
}

func compute(n int, times int) int {
	for i := 0; i < times; i++ {
		n = ((n << 6) ^ n) & 0xffffff
		n = ((n >> 5) ^ n) & 0xffffff
		n = ((n << 11) ^ n) & 0xffffff
	}
	return n
}

func Part2(lines []string) int {
	combinedMatches := make(map[[4]int]int)
	for _, line := range lines {
		matches := computeWithDiff(common.StringToInt(line), 2001)
		//fmt.Println(matches)
		mergeMapsInPlace(&combinedMatches, &matches)
	}

	key, value := findKeyWithLargestValue(&combinedMatches)

	fmt.Printf("key: %d, value: %d\n", key, value)

	return value
}

// Merge two maps and store the result in the first map
func mergeMapsInPlace(map1, map2 *map[[4]int]int) {
	// Iterate over the second map
	for key, value := range *map2 {
		// If the key already exists in the first map, add the values
		if existingValue, exists := (*map1)[key]; exists {
			(*map1)[key] = existingValue + value
		} else {
			// Otherwise, add the new key-value pair to the first map
			(*map1)[key] = value
		}
	}
}

func computeWithDiff(n int, times int) map[[4]int]int {
	prev := 0
	diffs := [4]int{}
	matches := make(map[[4]int]int)
	for i := 0; i < times; i++ {
		n = ((n << 6) ^ n) & 0xffffff
		n = ((n >> 5) ^ n) & 0xffffff
		n = ((n << 11) ^ n) & 0xffffff

		m := n % 10
		if i == 0 {
			prev = m
			continue
		}

		diff := m - prev
		prev = m

		if i <= 3 {
			diffs[i] = diff
		} else {
			diffs[0] = diffs[1]
			diffs[1] = diffs[2]
			diffs[2] = diffs[3]
			diffs[3] = diff
		}

		if i >= 3 {
			_, exists := matches[diffs]
			if !exists {
				matches[diffs] = m
			}
		}
	}
	return matches
}

// FindKeyWithLargestValue finds the key corresponding to the largest value in the map
func findKeyWithLargestValue(m *map[[4]int]int) ([4]int, int) {
	var maxKey [4]int
	maxValue := -1 // Initialize to a value below the possible range (0 to 9)

	for key, value := range *m {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}

	return maxKey, maxValue
}
