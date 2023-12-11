// Ez pz. I was told that this could be implemented as a one-liner, at least
// in python. I don't think that is possible in Go, and I don't think that my
// solution is optimal, but it works.
// On to part 2.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func handleError(err error, desc string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s: %v\n", desc, err)
		os.Exit(1)
	}
}

func getInputText() string {
	dat, err := os.ReadFile("./input.txt")
	handleError(err, "reading input file")
	str := string(dat)
	return strings.Trim(str, "\n \t")
}

// Pad lines with . to prevent out of bounds errors
func padLines(lines []string) []string {
	var result []string
	for _, line := range lines {
		result = append(result, "."+line+".")
	}
	// Add top and bottom row
	emptyLine := strings.Repeat(".", len(result[0]))
	result = append([]string{emptyLine}, result...)
	result = append(result, emptyLine)
	return result
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")
	result := 0

	lines = padLines(lines)

	reNum := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[^\d.]`)

	for i, line := range lines {
		// Find all numbers in the line
		matches := reNum.FindAllStringIndex(line, -1)
		for _, match := range matches {
			start, end := match[0], match[1]
			num, err := strconv.Atoi(line[start:end])
			handleError(err, "converting to int")

			// Check if symbol surrounds the number
			// If yes, add it to the result
			// Above
			if reSymbol.MatchString(lines[i-1][start-1 : end+1]) {
				result += num
				continue
			}
			// Below
			if reSymbol.MatchString(lines[i+1][start-1 : end+1]) {
				result += num
				continue
			}
			// Left
			if reSymbol.MatchString(line[start-1 : end]) {
				result += num
				continue
			}
			// Right
			if reSymbol.MatchString(line[start : end+1]) {
				result += num
				continue
			}
		}
	}

	fmt.Printf("Result: %d\n", result)
}
