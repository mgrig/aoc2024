package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func Part12(lines []string) (int, int) {
	operations, _ := parseLines(lines)

	sum1 := 0
	sum2 := 0
	for _, operation := range operations {
		if operation.CanBeMadeTrue() {
			sum1 += operation.result
			sum2 += operation.result
		} else if operation.CanBeMadeTrue3() {
			sum2 += operation.result
		}
	}

	return sum1, sum2
}

// parseLines parses each line in the input slice and returns a slice of Rows.
// Each line should be in the format: number: space-separated numbers
func parseLines(lines []string) ([]Operation, error) {
	var operations []Operation

	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line into key and values based on the colon delimiter
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format on line %d: %s", lineNumber+1, line)
		}

		// Parse the result number
		resultStr := strings.TrimSpace(parts[0])
		result, err := strconv.Atoi(resultStr)
		if err != nil {
			return nil, fmt.Errorf("invalid key on line %d: %s", lineNumber+1, resultStr)
		}

		// Parse the list of numbers
		valuesStr := strings.TrimSpace(parts[1])
		var values []int
		if valuesStr != "" {
			// Split the values by spaces
			valueParts := strings.Fields(valuesStr)
			values = make([]int, 0, len(valueParts))
			for _, valStr := range valueParts {
				val, err := strconv.Atoi(valStr)
				if err != nil {
					return nil, fmt.Errorf("invalid value on line %d: %s", lineNumber+1, valStr)
				}
				values = append(values, val)
			}
		}

		// Create a Row struct and append it to the slice
		operation := Operation{
			result:   result,
			operands: values,
		}
		operations = append(operations, operation)
	}

	return operations, nil
}
