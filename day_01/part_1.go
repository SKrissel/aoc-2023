// This one is pretty simple. Just read the file, split it into lines, and
// iterate over each line. For each line, find all the digits, and add the
// first and last digit together. Add that to the result. Print the result.
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
	return string(dat)
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")

	// Regex to find all digits
	re := regexp.MustCompile(`\d`)
	result := 0
	// Iterate over each line
	for _, line := range lines {
		// Skip empty lines (last line)
		if len(line) == 0 {
			continue
		}
		// Find all digits
		digits := re.FindAllString(line, -1)
		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]
		// Convert to int and add to result
		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		result += num
	}

	fmt.Printf("Result: %d\n", result)
}
