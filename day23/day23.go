package day23

import (
	"fmt"
	"sort"
	"strings"
)

func Part1(lines []string) int {
	edges := map[Edge]struct{}{}
	nodes := map[string]struct{}{}

	for _, line := range lines {
		tokens := strings.Split(line, "-")
		edges[NewEdge(tokens[0], tokens[1])] = struct{}{}
		edges[NewEdge(tokens[1], tokens[0])] = struct{}{}
		nodes[tokens[0]] = struct{}{}
		nodes[tokens[1]] = struct{}{}
	}

	// nodes to sorted array
	sortedNodes := make([]string, 0)
	for node := range nodes {
		sortedNodes = append(sortedNodes, node)
	}
	sort.Strings(sortedNodes)

	results := make(map[string]struct{}, 0)
	for t := 0; t < len(sortedNodes); t++ {
		tNode := sortedNodes[t]
		if !strings.HasPrefix(tNode, "t") {
			continue
		}
		for i := 0; i < len(sortedNodes); i++ {
			if t == i {
				continue
			}
			iNode := sortedNodes[i]
			if !containsEdge(&edges, NewEdge(tNode, iNode)) {
				continue
			}

			for j := i + 1; j < len(sortedNodes); j++ {
				if j == t {
					continue
				}
				jNode := sortedNodes[j]
				if !containsEdge(&edges, NewEdge(tNode, jNode)) {
					continue
				}

				if containsEdge(&edges, NewEdge(iNode, jNode)) {
					three := []int{t, i, j}
					sort.Ints(three)
					results[fmt.Sprintf("%s-%s-%s", sortedNodes[three[0]], sortedNodes[three[1]], sortedNodes[three[2]])] = struct{}{}
				}
			}
		}
	}

	return len(results)
}

func containsEdge(edges *map[Edge]struct{}, edge Edge) bool {
	_, exists := (*edges)[edge]
	return exists
}
