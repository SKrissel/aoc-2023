// This one was refreshingly easy. Just read the file, split it into lines,
// and iterate over each line splitting it into numbers. Then, for each number,
// find the diffence between itself and the previous number. Do this until the
// difference is 0, and then interpolate the sequence in reverse by adding a 0
// to the last sequence and then adding the last value of the current sequence
// to the last value of the previous sequence. The last value of the first
// sequence is the result. Add all the results together and print the result.
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

// Interpolate the sequences
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

		currentSequenceLast := sequence[len(sequence)-1]
		prevSequenceInterpolated := prevSequence[len(prevSequence)-1]
		// The interpolated value is the last value of the current sequence
		// plus the last value of the previous sequence
		interpolated := currentSequenceLast + prevSequenceInterpolated

		sequence = append(sequence, interpolated)
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

		// Interpolate the sequences and add the result to the
		// total result
		result += interpolate(sequences)
	}

	fmt.Printf("Result: %v\n", result)
}
