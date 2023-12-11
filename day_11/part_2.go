// What a struggle. Generating the full grid was easy but impossible to use, because
// the grid would be over 72 TB in size. So I had to find a way to calculate the
// distance between all pairs of stars without generating the full grid.
// My first idea was to calculate the distance step by step and add 1000000 to the
// distance if the line or column is empty. This gave wrong results.
// So I decided to add 1000000 to the coordinates of all stars that are below or to
// the right of an empty line or column. This gave the same wrong results.
// Now I knew that the problem had to lie somewhere else.
// I then noticed that I should have added 999999 instead of 1000000, because the line
// is already counted in the distance. This gave the correct result.
// So, in the end all 3 ideas were correct, but I had to find the right number to add.
package main

import (
	"fmt"
	"math"
	"os"
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

type Point struct {
	line int
	col  int
}

// Check if the line at 'line' is empty
func isEmptyLine(lines [][]rune, line int) bool {
	for _, char := range lines[line] {
		if char == '#' {
			return false
		}
	}
	return true
}

// Check if the column at 'col' is empty
func isEmptyColumn(lines [][]rune, col int) bool {
	for _, line := range lines {
		if line[col] == '#' {
			return false
		}
	}
	return true
}

// Find all stars in the grid marked with '#'
func findStars(lines [][]rune) []Point {
	var result []Point
	for l, line := range lines {
		for c, char := range line {
			if char == '#' {
				result = append(result, Point{
					line: l,
					col:  c,
				})
			}
		}
	}

	return result
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	stars := findStars(grid)
	fmt.Printf("Found %d stars\n", len(stars))

	// Find all empty lines
	var emptyLines []int
	for l := range grid {
		if isEmptyLine(grid, l) {
			emptyLines = append(emptyLines, l)
		}
	}

	// Find all empty columns
	var emptyColumns []int
	for c := range grid[0] {
		if isEmptyColumn(grid, c) {
			emptyColumns = append(emptyColumns, c)
		}
	}

	// add 1 to the length coordinates of all stars that are below or to the right of an empty line or column
	for i, star := range stars {
		// Lines
		addLines := 0
		for _, line := range emptyLines {
			if star.line >= line {
				addLines += 999999
			}
		}
		star.line += addLines

		// Columns
		addCols := 0
		for _, col := range emptyColumns {
			if star.col >= col {
				addCols += 999999
			}
		}
		star.col += addCols

		// Write back
		stars[i] = star
	}

	result := 0
	// Find the (shortest) distance between all pairs of stars
	for i, star1 := range stars {
		for j := i + 1; j < len(stars); j++ {
			star2 := stars[j]
			result += int(math.Abs(float64(star1.line-star2.line)) + math.Abs(float64(star1.col-star2.col)))
		}
	}

	fmt.Printf("Result: %d\n", result)

}
