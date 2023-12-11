// I made this one a bit harder than it had to be. I thought that I had to
// reject gears tjat are surrounded by more than two numbers. Turns out that
// there are no such gears in the input. Oh well.
// After playing around with the offsets for the gears, I found the solution.
// DEFINETELY harder than part 1.
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

type GearMatch struct {
	col     int
	line    int
	ratio   int
	checked bool
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")
	result := 0

	lines = padLines(lines)

	reNum := regexp.MustCompile(`\d+`)
	reGear := regexp.MustCompile(`\*`)

	gearMatches := make([]GearMatch, 0)

	for i, line := range lines {
		// Find all numbers in the line
		matches := reNum.FindAllStringIndex(line, -1)
		for _, match := range matches {
			start, end := match[0], match[1]
			num, err := strconv.Atoi(line[start:end])
			handleError(err, "converting to int")

			// Check if Gear ('*') surrounds the number
			// If yes, add it to the matches with an offset
			// because the matches are found with the relative position
			// of the number in the line.
			// Above
			gearMatch := reGear.FindStringIndex(lines[i-1][start-1 : end+1])
			if gearMatch != nil {
				gearMatches = append(gearMatches, GearMatch{
					col:   match[0] + gearMatch[0] - 1,
					line:  i - 1,
					ratio: num,
				})
			}

			// Below
			gearMatch = reGear.FindStringIndex(lines[i+1][start-1 : end+1])
			if gearMatch != nil {
				gearMatches = append(gearMatches, GearMatch{
					col:   match[0] + gearMatch[0] - 1,
					line:  i + 1,
					ratio: num,
				})
			}

			// Left
			gearMatch = reGear.FindStringIndex(line[start-1 : end])
			if gearMatch != nil {
				gearMatches = append(gearMatches, GearMatch{
					col:   match[0] - 1,
					line:  i,
					ratio: num,
				})
			}

			// Right
			gearMatch = reGear.FindStringIndex(line[start : end+1])
			if gearMatch != nil {
				gearMatches = append(gearMatches, GearMatch{
					col:   match[1],
					line:  i,
					ratio: num,
				})
			}
		}
	}

	// Go through all matches and check if there are two matches with the same
	// coordinates. If yes, multiply the ratios and add them to the result.
	for i, gearMatch := range gearMatches {
		foundMatches := make([]int, 0)
		foundMatches = append(foundMatches, i)

		for j, gearMatch2 := range gearMatches {
			if i == j || gearMatch.checked || gearMatch2.checked {
				// Skip if the match is the same as the one we are checking
				// or if one of the matches has already been checked.
				// This has to be done via a flag, because we can't remove
				// elements from the slice while iterating over it.
				continue
			}
			if gearMatch.col == gearMatch2.col && gearMatch.line == gearMatch2.line {
				// Both matches refer to the same gear!
				foundMatches = append(foundMatches, j)
			}
		}

		// Check if we found two matches
		if len(foundMatches) == 2 {
			result += gearMatches[foundMatches[0]].ratio * gearMatches[foundMatches[1]].ratio
			fmt.Printf("Found match: %d, %d\n", gearMatches[foundMatches[0]].ratio, gearMatches[foundMatches[1]].ratio)
		}

		// Mark all found matches as checked
		for _, foundMatch := range foundMatches {
			gearMatches[foundMatch].checked = true
		}
	}

	fmt.Printf("Result: %d\n", result)
}
