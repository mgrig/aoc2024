package day19

import (
	"sort"
	"strings"
)

func Part1(lines []string) int {
	towels := getTowels(lines)

	cache := make(map[string]bool)

	count := 0
	for i := 2; i < len(lines); i++ {
		ok := canCompose(towels, lines[i], &cache)
		//fmt.Println(ok, lines[i])
		if ok {
			count++
		}
	}

	return count
}

func Part2(lines []string) int {
	towels := getTowels(lines)

	cache := make(map[string]int)

	count := 0
	for i := 2; i < len(lines); i++ {
		arrangements := CountArrangements(towels, lines[i], &cache)
		count += arrangements
	}

	return count
}

func CountArrangements(towels []string, target string, cache *map[string]int) int {
	count := 0
	for _, towel := range towels {
		if towel == target {
			count++
		} else if strings.HasPrefix(target, towel) {
			rest := target[len(towel):]
			subCounts, exists := (*cache)[rest]
			if !exists {
				subCounts = CountArrangements(towels, rest, cache)
				(*cache)[rest] = subCounts
			}
			count += subCounts
		}
	}
	return count
}

func canCompose(towels []string, target string, cache *map[string]bool) bool {
	for _, towel := range towels {
		if towel == target {
			return true
		}
		if strings.HasPrefix(target, towel) {
			can, exists := (*cache)[target]
			if !exists {
				rest := target[len(towel):]
				can = canCompose(towels, rest, cache)
				(*cache)[rest] = can
			}
			if can {
				return true
			}
		}
	}
	return false
}

func getTowels(lines []string) []string {
	towels := strings.Split(lines[0], ", ")

	sort.Slice(towels, func(i, j int) bool {
		if len(towels[i]) != len(towels[j]) {
			// Sort by length descending
			return len(towels[i]) > len(towels[j])
		}
		// Sort alphabetically if lengths are the same
		return towels[i] < towels[j]
	})

	return towels
}
