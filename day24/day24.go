package day24

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
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

	return computeNumber(sortedKeysStartingWith(knowns, "z"), knowns)
}

func Part2(lines []string) string {
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

	x := computeNumber(sortedKeysStartingWith(knowns, "x"), knowns)
	y := computeNumber(sortedKeysStartingWith(knowns, "y"), knowns)
	z := Part1(lines)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Printf("x + y: %d\n%b\n", x+y, x+y)
	fmt.Printf("z: %d\n%b\n", z, z)

	fmt.Println("GATES:")
	gates := make([]Gate, 0)
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
		gates = append(gates, gate)
		unknowns[out] = gate
	}

	//outs := make([]string, 0)
	//for i := 0; i <= 11; i++ {
	//	outs = append(outs, fmt.Sprintf("z%2d", i))
	//}
	//outs = append(outs, "jmb")

	//filtered := filterByOut(&unknowns, "z29", "z28")
	////filtered := filterByOut(&unknowns, outs...)
	//fmt.Println(filtered)
	//toDotFile(knowns, gates, filtered)

	//carry := make([]string, 46)
	//carry[0] = "rnv"
	//for i := 1; i <= 45; i++ {
	//	xi := fmt.Sprintf("x%02d", i)
	//	yi := fmt.Sprintf("y%02d", i)
	//	zi := fmt.Sprintf("z%02d", i)
	//	carry_i_1 := carry[i-1]
	//
	//	xi_xor_yi, found := getGateByInputsAndTyp(&gates, xi, yi, XOR)
	//	if !found {
	//		panic("gate not found")
	//	}
	//
	//	xi_xor_yi_xor_carry_i_1, found := getGateByInputsAndTyp(&gates, xi_xor_yi.out, carry_i_1, XOR)
	//	if !found {
	//		fmt.Printf("i=%d, missing gate XOR(%s, %s)\n", i, xi_xor_yi.out, carry_i_1)
	//		if i == 28 {
	//			badGate, _ := getGateByInputsAndTyp(&gates, "x28", "y28", XOR)
	//			newGate := badGate.Clone()
	//			newGate.out = "pvb" // instead of previous qdq
	//			replaceGate(&gates, badGate, newGate)
	//
	//			badGate, _ = getGateByInputsAndTyp(&gates, "x28", "y28", AND)
	//			newGate = badGate.Clone()
	//			newGate.out = "qdq" // instead of previous pvb
	//			replaceGate(&gates, badGate, newGate)
	//		}
	//		i--
	//		continue
	//	}
	//	if xi_xor_yi_xor_carry_i_1.out != zi {
	//		//fmt.Printf("i=%d, correct value goes to wrong z** output: wanted:%s, got:%s\n", i, zi, xi_xor_yi_xor_carry_i_1.out)
	//
	//		//FIX IN PLACE
	//		//replaceGate(&gates, xi_xor_yi_xor_carry_i_1, NewGate(xi_xor_yi_xor_carry_i_1.in1, xi_xor_yi_xor_carry_i_1.in2, xi_xor_yi_xor_carry_i_1.typ, zi))
	//		if i == 11 {
	//			// z11 and vkq are wrong!
	//			badGate, found := getGateByOutput(&gates, "z11")
	//			if !found {
	//				panic("gate not found")
	//			}
	//
	//			newGate := xi_xor_yi_xor_carry_i_1.Clone()
	//			newGate.out = "z11"
	//			replaceGate(&gates, xi_xor_yi_xor_carry_i_1, newGate)
	//
	//			newGate = badGate
	//			newGate.out = "vkq"
	//			replaceGate(&gates, badGate, newGate)
	//		}
	//
	//		if i == 24 {
	//			badGate, found := getGateByOutput(&gates, "z24")
	//			if !found {
	//				panic("gate not found")
	//			}
	//			newGate := xi_xor_yi_xor_carry_i_1.Clone()
	//			newGate.out = "z24"
	//			replaceGate(&gates, xi_xor_yi_xor_carry_i_1, newGate)
	//
	//			newGate = badGate
	//			newGate.out = "mmk"
	//			replaceGate(&gates, badGate, newGate)
	//		}
	//
	//		i-- // redo this i-step
	//		continue
	//	}
	//
	//	xi_and_yi, found := getGateByInputsAndTyp(&gates, xi, yi, AND)
	//	if !found {
	//		panic("gate not found")
	//	}
	//
	//	xi_xor_yi_and_carry_i_1, found := getGateByInputsAndTyp(&gates, xi_xor_yi.out, carry_i_1, AND)
	//	if !found {
	//		panic("gate not found")
	//	}
	//
	//	carry_i, found := getGateByInputsAndTyp(&gates, xi_and_yi.out, xi_xor_yi_and_carry_i_1.out, OR)
	//	if !found {
	//		// one of the inputs points wrong
	//		//TODO manual fix for now:
	//		if xi_and_yi.out == "z11" {
	//			fmt.Printf("WRONG: z11\n")
	//			replaceGate(&gates, xi_and_yi, NewGate(xi_and_yi.in1, xi_and_yi.in2, xi_and_yi.typ, xi_xor_yi_and_carry_i_1.in1))
	//
	//			goodTarget, found := getGateByInputOutTyp(&gates, xi_xor_yi_and_carry_i_1.out, "z11", OR)
	//			if !found {
	//				panic("gate not found")
	//			}
	//			fmt.Printf("find good target gate %s", goodTarget)
	//		}
	//		i--
	//		continue
	//	}
	//	carry[i] = carry_i.out
	//}

	// manually found pairs, using the debugging code above!
	pairs := []string{"z11", "vkq", "z24", "mmk", "pvb", "qdq", "z38", "hqh"}
	sort.Strings(pairs)

	return strings.Join(pairs, ",")
}

func replaceGate(gates *[]Gate, targetGate Gate, newGate Gate) {
	for i, gate := range *gates {
		if gate == targetGate {
			(*gates)[i] = newGate
			return
		}
	}
	panic("gate not found")
}

func getGateByInputOutTyp(gates *[]Gate, in, out string, typ int) (gate Gate, found bool) {
	for _, gate := range *gates {
		if gate.typ != typ || gate.out != out {
			continue
		}
		if gate.in1 == in || gate.in2 == in {
			return gate, true
		}
	}
	return gate, false
}

func getGateByOutput(gates *[]Gate, out string) (gate Gate, found bool) {
	for _, gate := range *gates {
		if gate.out == out {
			return gate, true
		}
	}
	return gate, false
}

func getGateByInputsAndTyp(gates *[]Gate, in1, in2 string, typ int) (gate Gate, found bool) {
	for _, gate := range *gates {
		if gate.typ != typ {
			continue
		}
		if (gate.in1 == in1 && gate.in2 == in2) || (gate.in1 == in2 && gate.in2 == in1) {
			return gate, true
		}
	}
	return gate, false
}

func filterByOut(unknowns *map[string]Gate, outs ...string) map[string]struct{} {
	ret := make(map[string]struct{})

	toVisit := outs
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		ret[current] = struct{}{}

		gate, exists := (*unknowns)[current]
		if exists {
			_, exists = ret[gate.in1]
			if !exists {
				toVisit = append(toVisit, gate.in1)
			}
			_, exists = ret[gate.in2]
			if !exists {
				toVisit = append(toVisit, gate.in2)
			}
		}
	}

	return ret
}

func filterContains(nodeFilter *map[string]struct{}, node string) bool {
	_, exists := (*nodeFilter)[node]
	return exists
}

func toDotFile(knowns map[string]bool, gates []Gate, nodeFilter map[string]struct{}) {
	fmt.Println("digraph G {")

	//fmt.Println(" subgraph cluster_x {")
	//for _, x := range sortedKeysStartingWith(knowns, "x") {
	//	if filterContains(&nodeFilter, x) {
	//		fmt.Printf("    %s\n", x)
	//	}
	//}
	//fmt.Println("  }")
	//
	//fmt.Println("\n subgraph cluster_y {")
	//for _, y := range sortedKeysStartingWith(knowns, "y") {
	//	if filterContains(&nodeFilter, y) {
	//		fmt.Printf("    %s\n", y)
	//	}
	//}
	//fmt.Println("  }")

	for i := 0; i < len(gates); i++ {
		gate := gates[i]
		gateName := fmt.Sprintf("gate_%d", i)
		if filterContains(&nodeFilter, gate.out) {
			fmt.Printf("  %s [label=\"%s\"]\n", gateName, TypToString(gate.typ))
			fmt.Printf("    %s -> %s\n", gate.in1, gateName)
			fmt.Printf("    %s -> %s\n", gate.in2, gateName)
			fmt.Printf("    %s -> %s\n", gateName, gate.out)
		}
	}

	//fmt.Println("\n subgraph cluster_gates {")
	//for i, gate := range gates {
	//	if filterContains(&nodeFilter, gate.out) {
	//		gateName := fmt.Sprintf("gate_%d", i)
	//		fmt.Printf("    %s\n", gateName)
	//	}
	//}
	//fmt.Println("  }")

	fmt.Println("\n subgraph cluster_z {")
	outs := make([]string, 0)
	for _, gate := range gates {
		if strings.HasPrefix(gate.out, "z") {
			if filterContains(&nodeFilter, gate.out) {
				outs = append(outs, gate.out)
			}
		}
	}
	sort.Strings(outs)
	for _, out := range outs {
		fmt.Printf("    %s\n", out)
	}
	fmt.Println("  }")

	fmt.Println("}")
}

func sortedKeysStartingWith(m map[string]bool, prefix string) []string {
	// Collect keys starting with "z"
	keys := []string{}
	for key := range m {
		if len(key) > 0 && strings.HasPrefix(key, prefix) {
			keys = append(keys, key)
		}
	}

	// Sort keys in decreasing order
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	return keys
}

func computeNumber(keys []string, m map[string]bool) int {
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
