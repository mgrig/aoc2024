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
		nodes[tokens[0]] = struct{}{}
		nodes[tokens[1]] = struct{}{}
	}

	// nodes to sorted array
	sortedNodes := make([]string, 0)
	for node := range nodes {
		sortedNodes = append(sortedNodes, node)
	}
	sort.Strings(sortedNodes)

	count := 0
	for i := 0; i < len(sortedNodes); i++ {
		iNode := sortedNodes[i]
		if !strings.HasPrefix(iNode, "t") {
			continue
		}
		for j := 0; j < len(sortedNodes); j++ {
			if i == j {
				continue
			}
			jNode := sortedNodes[j]
			if !containsEdge(&edges, NewEdge(iNode, jNode)) {
				continue
			}
			if strings.HasPrefix(jNode, "t") && j < i {
				continue
			}

			for k := j + 1; k < len(sortedNodes); k++ {
				if k == i {
					continue
				}
				kNode := sortedNodes[k]
				if !containsEdge(&edges, NewEdge(iNode, kNode)) {
					continue
				}

				if containsEdge(&edges, NewEdge(jNode, kNode)) {
					fmt.Printf("%s-%s-%s\n", iNode, jNode, kNode)
					count++
				}
			}
		}
	}

	return count
}

func containsEdge(edges *map[Edge]struct{}, edge Edge) bool {
	_, exists := (*edges)[edge]
	if exists {
		return true
	}
	_, exists = (*edges)[NewEdge(edge.b, edge.a)]
	return exists
}
