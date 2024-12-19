package day19

import (
	"fmt"
	"sort"
	"strings"
)

func Part1(lines []string) int {
	towels := strings.Split(lines[0], ", ")

	sort.Slice(towels, func(i, j int) bool {
		if len(towels[i]) != len(towels[j]) {
			// Sort by length descending
			return len(towels[i]) > len(towels[j])
		}
		// Sort alphabetically if lengths are the same
		return towels[i] < towels[j]
	})
	fmt.Println(towels)

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
