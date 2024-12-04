package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// rotateMatrix90 rotates a 2D rune slice by 90 degrees clockwise
func rotateMatrix90(matrix [][]rune) [][]rune {
	if len(matrix) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])
	rotated := make([][]rune, cols)

	for i := 0; i < cols; i++ {
		rotated[i] = make([]rune, rows)
		for j := 0; j < rows; j++ {
			rotated[i][j] = matrix[rows-1-j][i]
		}
	}

	return rotated
}

// rotateMatrix45 rotates a 2D rune slice by 45 degrees clockwise
func rotateMatrix45(matrix [][]rune) []string {
	if len(matrix) == 0 {
		return []string{}
	}

	rows := len(matrix)
	cols := len(matrix[0])
	maxSum := rows + cols - 2

	rotatedLines := make([]string, 0, rows+cols-1)

	// Define the sum order: start from the center and alternate outwards
	kCenter := maxSum / 2
	sumOrder := []int{}

	for d := 0; d <= kCenter; d++ {
		leftSum := kCenter - d
		rightSum := kCenter + d
		if leftSum >= 0 {
			sumOrder = append(sumOrder, leftSum)
		}
		if rightSum <= maxSum && d != 0 {
			sumOrder = append(sumOrder, rightSum)
		}
	}

	// Collect characters based on the defined sum order
	for _, sum := range sumOrder {
		var lineChars []rune
		for i := 0; i < rows; i++ {
			j := cols - 1 - sum + i
			if j >= 0 && j < cols {
				lineChars = append(lineChars, matrix[i][j])
			}
		}
		if len(lineChars) > 0 {
			rotatedLines = append(rotatedLines, string(lineChars))
		}
	}

	return rotatedLines
}

// readMatrix reads the input file and returns a 2D rune slice along with the maximum number of columns
func readMatrix(inputFile string) ([][]rune, int, error) {
	infile, err := os.Open(inputFile)
	if err != nil {
		return nil, 0, fmt.Errorf("error opening input file: %v", err)
	}
	defer infile.Close()

	scanner := bufio.NewScanner(infile)
	var matrix [][]rune
	maxCols := 0

	// Read the input file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Convert the line to a slice of runes (characters)
		runes := []rune(line)
		matrix = append(matrix, runes)
		if len(runes) > maxCols {
			maxCols = len(runes)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, fmt.Errorf("error reading input file: %v", err)
	}

	// Pad shorter rows with spaces to ensure all lines have the same length
	for i := range matrix {
		if len(matrix[i]) < maxCols {
			padding := make([]rune, maxCols-len(matrix[i]))
			for j := range padding {
				padding[j] = ' '
			}
			matrix[i] = append(matrix[i], padding...)
		}
	}

	return matrix, maxCols, nil
}

// writeMatrix writes the rotated matrix to the output file with proper EOL handling
func writeMatrix(rotated interface{}, outputFile string) error {
	outfile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file '%s': %v", outputFile, err)
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)

	switch v := rotated.(type) {
	case [][]rune:
		for _, row := range v {
			line := string(row)
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				return fmt.Errorf("error writing to output file '%s': %v", outputFile, err)
			}
		}
	case []string:
		for _, line := range v {
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				return fmt.Errorf("error writing to output file '%s': %v", outputFile, err)
			}
		}
	default:
		return fmt.Errorf("unsupported type for rotated matrix")
	}

	// Flush the buffer to ensure all data is written
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing to output file '%s': %v", outputFile, err)
	}

	return nil
}

func WriteToFile(matrix interface{}, dirName, nameOnly, suffix, ext string) {
	// Create the output filename
	outputFileName := fmt.Sprintf("%s_%s%s", nameOnly, suffix, ext)
	outputFilePath := filepath.Join(dirName, outputFileName)

	// Write the rotated matrix to the output file
	err := writeMatrix(matrix, outputFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Check for correct number of arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: rotate90 <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Check if the input file exists and is a regular file
	fileInfo, err := os.Stat(inputFile)
	if os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' does not exist.\n", inputFile)
		os.Exit(1)
	}
	if fileInfo.IsDir() {
		fmt.Printf("Error: '%s' is a directory, not a file.\n", inputFile)
		os.Exit(1)
	}

	// Read the input matrix
	matrix, _, err := readMatrix(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Determine the base name and extension
	baseName := filepath.Base(inputFile)
	dirName := filepath.Dir(inputFile)
	ext := filepath.Ext(baseName)
	nameOnly := strings.TrimSuffix(baseName, ext)

	// Initialize current matrix as the original matrix
	currentMatrix := matrix

	rotated45 := rotateMatrix45(currentMatrix)
	WriteToFile(rotated45, dirName, nameOnly, "45", ext)

	rotated90 := rotateMatrix90(currentMatrix)
	WriteToFile(rotated90, dirName, nameOnly, "90", ext)

	rotated45 = rotateMatrix45(rotated90)
	WriteToFile(rotated45, dirName, nameOnly, "135", ext)

	rotated90 = rotateMatrix90(rotated90)
	WriteToFile(rotated90, dirName, nameOnly, "180", ext)

	rotated45 = rotateMatrix45(rotated90)
	WriteToFile(rotated45, dirName, nameOnly, "225", ext)

	rotated90 = rotateMatrix90(rotated90)
	WriteToFile(rotated90, dirName, nameOnly, "270", ext)

	rotated45 = rotateMatrix45(rotated90)
	WriteToFile(rotated45, dirName, nameOnly, "315", ext)

}
