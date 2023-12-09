// This might have been the fastest part 2 ever. I just had to change the
// interpolate() function to work in reverse and interpolate the sequences
// at the start instead of the end.
package main

import (
	"fmt"
	"os"
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

// Parse a line of history into a slice of ints
func parseHistory(line string) []int {
	values := strings.Split(line, " ")
	var result []int
	for _, value := range values {
		parsedValue, _ := strconv.Atoi(value)
		result = append(result, parsedValue)
	}

	return result
}

// Generate the sequences until the last sequence is all 0s
// Returns the sequences as a pyramid
func generateSequence(sequences [][]int) [][]int {
	done := false

	// Run until the last sequence is all 0s
	for !done {
		done = true
		currentSequence := sequences[len(sequences)-1]

		var nextSequence []int
		for i := 1; i <= len(currentSequence)-1; i++ {
			val := currentSequence[i]
			prevVal := currentSequence[i-1]
			// The next value is the difference between the current value and
			// the previous value
			diff := val - prevVal

			nextSequence = append(nextSequence, diff)
			if diff != 0 {
				// We are not done yet
				done = false
			}
		}
		sequences = append(sequences, nextSequence)
	}

	return sequences
}

// Interpolate the sequences in reverse
func interpolate(sequences [][]int) int {
	var result int
	startIndex := len(sequences) - 1

	// Append a 0 to the last sequence (only 0s)
	lastSequence := sequences[startIndex]
	lastSequence = append(lastSequence, 0)
	startIndex--

	// Go over the sequences in reverse (Top of the pyramid to bottom)
	for i := startIndex; i >= 0; i-- {
		sequence := sequences[i]
		prevSequence := sequences[i+1]

		currentSequenceFirst := sequence[0]
		prevSequenceInterpolated := prevSequence[0]
		// The interpolated value is the last value of the current sequence
		// plus the last value of the previous sequence
		interpolated := currentSequenceFirst - prevSequenceInterpolated

		// Prepend the interpolated value to the sequence by using a composite literal
		sequence = append([]int{interpolated}, sequence...)
		sequences[i] = sequence
		result = interpolated
	}

	// We only need the last interpolated value
	return result
}

func main() {
	// Read the input file
	input := getInputText()
	lines := strings.Split(input, "\n")
	result := 0

	// Iterate over each line
	for _, line := range lines {
		// Parse all history values into a slice of ints
		history := parseHistory(line)

		// Create a slice of slices of ints, with the first slice being the
		// history, in order to be able to generate the sequences as a pyramid
		var sequences [][]int
		sequences = append(sequences, history)

		// Generate the sequences until the last sequence is all 0s
		sequences = generateSequence(sequences)

		// Interpolate the sequences in reverse and add the result to the
		// total result
		result += interpolate(sequences)
	}

	fmt.Printf("Result: %v\n", result)
}
