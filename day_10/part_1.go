// Nice problem. I liked the idea. It took longer than expected to implement
// the solution, but it was still fun. I think the solution could be improved
// by only traversing the path once, because we don't really care abput the
// Second direction. We already know that it should have the same length as
// the first direction.
// Another point of improvement would be the huge switch-statement in the
// step-function. I think there is a better way to implement this, but I
// couldn't think of one.
// All in all, nice part 1 for today.
package main

import (
	"errors"
	"fmt"
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

// Find the starting point in the grid marked with 'S'
// Returns error, Point
func findStart(grid []string) (error, Point) {
	// Go through each line and column and check if the current char is 'S'
	for l, line := range grid {
		for c, col := range line {
			if col == 'S' {
				return nil, Point{
					line: l,
					col:  c,
				}
			}
		}
	}

	return errors.New("No starting Point"), Point{}
}

// Do one step in the simulation. Check if the current pipe allows the direction to continue.
// Direction has to be a Vector with 0 for one value, 1 or -1 for other
// Returns error, new Postion, new Direction
func step(pos Point, direction Point, grid []string) (error, Point, Point) {
	// Check if the new position is out of bounds
	if pos.line < 0 || pos.line >= len(grid) || pos.col < 0 || pos.col >= len(grid[0]) {
		return errors.New("Out of Bounds!"), Point{}, Point{}
	}
	currentPipe := grid[pos.line][pos.col]

	// Switch-statement in go do not automatically fall through.
	switch currentPipe {
	case '.':
		return errors.New("No valid Pipe"), Point{}, Point{}
	case 'S':
		// Do noting and allow this step
	case '|':
		if direction.line == 0 {
			return errors.New("No pipe access from this side"), Point{}, Point{}
		}
		// Continue in same direction
	case '-':
		if direction.col == 0 {
			return errors.New("No pipe access from this side"), Point{}, Point{}
		}
		// Continue in same direction
	case 'L':
		if direction.line == -1 || direction.col == 1 {
			return errors.New("No pipe access from this side"), Point{}, Point{}
		}
		// line = 1 -> col = 1
		// col = -1 -> line = -1
		direction.col, direction.line = direction.line, direction.col
	case 'J':
		if direction.line == -1 || direction.col == -1 {
			return errors.New("No pipe access from this side"), Point{}, Point{}
		}
		// line = 1 -> col = -1
		// col = 1 -> line = -1
		direction.col, direction.line = direction.line*-1, direction.col*-1
	case '7':
		if direction.line == 1 || direction.col == -1 {
			return errors.New("No pipe access from this side"), Point{}, Point{}
		}
		// line = -1 -> col = -1
		// col = 1 -> line = 1
		direction.col, direction.line = direction.line, direction.col
	case 'F':
		if direction.line == 1 || direction.col == 1 {
			return errors.New("No pipe access from this side"), Point{}, Point{}
		}
		// line = -1 -> col = 1
		// col = -1 -> line = 1
		direction.col, direction.line = direction.line*-1, direction.col*-1
	}

	// Move one step in the direction
	pos.line += direction.line
	pos.col += direction.col

	// Check if the new position is out of bounds again
	if pos.line < 0 || pos.line >= len(grid) || pos.col < 0 || pos.col >= len(grid[0]) {
		return errors.New("Out of Bounds!"), Point{}, Point{}
	}
	return nil, pos, direction
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")

	err, start := findStart(lines)
	handleError(err, "Getting starting point")
	fmt.Printf("Found starting point: %v\n\n", start)

	// Since we do not know what pipe hides behind the starting point, we have to
	// try all directions.
	directions := map[string]Point{
		"North": {line: -1, col: 0},
		"South": {line: 1, col: 0},
		"East":  {line: 0, col: 1},
		"West":  {line: 0, col: -1},
	}

	steps := make(map[string]int)

	for name, direction := range directions {
		fmt.Printf("Running simulation for direction '%v'...\n", name)

		// Struct with only primitives is copied by value, not by reference in go.
		currentPos := start
		// Use a random pipe to start the simulation. We can't use 'S' because
		// that would end the simulation immediately.
		currentPipe := '.'

		for currentPipe != 'S' {
			// Perform one step in the simulation
			err, currentPos, direction = step(currentPos, direction, lines)
			if err != nil {
				// If the simulation fails, remove the direction from the list
				fmt.Printf("Failed! %v\n", err)
				delete(steps, name)
				break
			}
			steps[name]++
			// Get the current pipe to check if the simulation is done
			currentPipe = rune(lines[currentPos.line][currentPos.col])
		}
		fmt.Println()
	}

	// Exactly two directions should form a loop
	result := 0
	fmt.Println("Simulation Done.\nTwo Directions found:")
	for key, value := range steps {
		fmt.Printf("%v (%v steps)\n", key, value)
		result += value
	}

	// Kind of a hacky way to get the result, but it works.
	// The result is in the middle of the loop, so we have to divide by 4 to get
	// the center between the two directions.
	// THIS ONLY WORKS IF THE LOOP IS DIVISIBLE BY 2! (Which it is in this case)
	fmt.Printf("Result: %v\n", result/4)
}
