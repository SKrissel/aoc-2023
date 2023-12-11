// This one was doable and pretty easy to do in go.
// Expanding the columns was a bit difficult, but I found a nice way to do it by
// transposing the grid, expanding the lines, and transposing back.
// My maths teacher would be proud.
// Calculating the distance between all pairs of stars was a matter of calculating
// the absolute difference between the line and column of the two stars. Nice.
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
	dat, err := os.ReadFile("./example.txt")
	handleError(err, "reading input file")
	str := string(dat)
	return strings.Trim(str, "\n \t")
}

type Point struct {
	line int
	col  int
}

// Shamelessly stolen and adapted from https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func transpose(slice [][]rune) [][]rune {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]rune, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

// Expand lines by adding a line if there is no star in the line
func expandLines(lines [][]rune) [][]rune {
	// expand lines
	var result [][]rune
	for _, line := range lines {
		result = append(result, line)
		// check if line contains a star
		for i, char := range line {
			if char == '#' {
				// line contains a star, break
				break
			}

			if i == len(line)-1 {
				// line does not contain a star, expand
				result = append(result, line)
			}
		}

	}

	return result
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

	grid = expandLines(grid)
	// Since expanding lines is easier than expanding columns, we transpose the grid, expand lines, and transpose back
	grid = transpose(grid)
	grid = expandLines(grid)
	grid = transpose(grid)

	// The grid is now cosmically expanded
	for _, line := range grid {
		fmt.Println(string(line))
	}

	stars := findStars(grid)

	result := 0
	// Find the (shortest) distance between all pairs of stars
	for i, star1 := range stars {
		for j := i + 1; j < len(stars); j++ {
			star2 := stars[j]
			result += int(math.Abs(float64(star1.line-star2.line)) + math.Abs(float64(star1.col-star2.col)))
		}
	}

	fmt.Printf("Result: %d\n", result)

	fmt.Printf("Stars: %d\n", stars)

}
