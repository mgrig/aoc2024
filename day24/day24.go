package day24

import (
	"fmt"
	"regexp"
	"sort"
)

func Part1(lines []string) int {
	knowns := make(map[string]bool)

	regInput := regexp.MustCompile(`(.*?): (.*)`)
	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		tokens := regInput.FindStringSubmatch(line)
		if tokens[2] == "0" {
			knowns[tokens[1]] = false
		} else {
			knowns[tokens[1]] = true
		}
	}

	gates := make(map[Gate]struct{})
	unknowns := make(map[string]Gate)
	regGate := regexp.MustCompile(`(.*?) (AND|XOR|OR) (.*?) -> (.*)`)
	for i = i + 1; i < len(lines); i++ {
		line := lines[i]
		tokens := regGate.FindStringSubmatch(line)
		in1 := tokens[1]
		typ := TypMap[tokens[2]]
		in2 := tokens[3]
		out := tokens[4]

		gate := NewGate(in1, in2, typ, out)
		gates[gate] = struct{}{}
		unknowns[out] = gate
	}

	for len(unknowns) > 0 {
		// go through all unknowns and compute what is possible
		becameKnown := make([]string, 0)
		for _, gate := range unknowns {
			if isKnown(&knowns, gate.in1) && isKnown(&knowns, gate.in2) {
				var out bool
				switch gate.typ {
				case AND:
					out = knowns[gate.in1] && knowns[gate.in2]
				case OR:
					out = knowns[gate.in1] || knowns[gate.in2]
				case XOR:
					out = knowns[gate.in1] != knowns[gate.in2]
				default:
					panic(fmt.Sprintf("Unknown type: %s", gate.typ))
				}
				knowns[gate.out] = out
				becameKnown = append(becameKnown, gate.out)
			}
		}
		for _, unknown := range becameKnown {
			delete(unknowns, unknown)
		}
	}

	return computeNumber(knowns)
}

func computeNumber(m map[string]bool) int {
	// Collect keys starting with "z"
	keys := []string{}
	for key := range m {
		if len(key) > 0 && key[0] == 'z' { // Check if the key starts with 'z'
			keys = append(keys, key)
		}
	}

	// Sort keys in decreasing order
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	// Convert boolean values to bits and form the integer
	result := 0
	for _, key := range keys {
		result = (result << 1) | btoi(m[key])
	}

	return result
}

// Helper function to convert bool to int (0 or 1)
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func isKnown(knowns *map[string]bool, key string) bool {
	_, exists := (*knowns)[key]
	return exists
}
