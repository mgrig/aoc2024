package day23

import (
	"sort"
	"strings"
)

// Graph represents an undirected graph using an adjacency list with string nodes
type Graph struct {
	adjacency map[string][]string
}

// NewGraph creates and initializes a new graph
func NewGraph() *Graph {
	return &Graph{
		adjacency: make(map[string][]string),
	}
}

// AddEdge adds an edge between two vertices
func (g *Graph) AddEdge(u, v string) {
	g.adjacency[u] = append(g.adjacency[u], v)
	g.adjacency[v] = append(g.adjacency[v], u)
}

// GetNeighbors returns the neighbors of a given vertex
func (g *Graph) GetNeighbors(v string) []string {
	return g.adjacency[v]
}

// FindMaximalClique finds one maximal clique of the largest size
func (g *Graph) FindMaximalClique() []string {
	var maxClique []string
	allNodes := []string{}
	for node := range g.adjacency {
		allNodes = append(allNodes, node)
	}

	g.bronKerbosch([]string{}, allNodes, []string{}, &maxClique)
	return maxClique
}

// bronKerbosch is a recursive implementation of the Bron–Kerbosch algorithm
func (g *Graph) bronKerbosch(r []string, p []string, x []string, maxClique *[]string) {
	if len(p) == 0 && len(x) == 0 {
		// Base case: R is a maximal clique
		if len(r) > len(*maxClique) {
			*maxClique = append([]string{}, r...) // Update maxClique
		}
		return
	}

	for i := 0; i < len(p); i++ {
		v := p[i]
		rNew := append(r, v)
		pNew := g.intersect(p, g.GetNeighbors(v))
		xNew := g.intersect(x, g.GetNeighbors(v))
		g.bronKerbosch(rNew, pNew, xNew, maxClique)
		p = append(p[:i], p[i+1:]...)
		x = append(x, v)
		i-- // Adjust index after modifying p
	}
}

// intersect returns the intersection of two slices
func (g *Graph) intersect(a, b []string) []string {
	m := make(map[string]bool)
	for _, v := range b {
		m[v] = true
	}

	var result []string
	for _, v := range a {
		if m[v] {
			result = append(result, v)
		}
	}
	return result
}

func Part2(lines []string) string {
	// Create a graph
	g := NewGraph()

	for _, line := range lines {
		tokens := strings.Split(line, "-")
		g.AddEdge(tokens[0], tokens[1])
	}

	// Find one maximal clique of the largest size
	maximalClique := g.FindMaximalClique()
	sort.Strings(maximalClique)

	return strings.Join(maximalClique, ",")
}
