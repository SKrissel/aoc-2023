// This is where it gets complicated. Converting the words to numbers is
// relatively simple. Just replace each word with the corresponding number.
// This, however does not work, because some words have common letters.
// For example, looking at the input "twone", we can see that the result
// should be "21", but if we just replace each word with the corresponding
// number, we get "2ne", because "two" and "one" both share the same "e".
// To fix this, we need to create a copy of each line containing only the
// numbers and not replace the words in the original line. Then, we can
// continue like in part 1
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Map of number-words to numbers
var strToNumMap = map[string]string{
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",

	"one":   "1",
	"six":   "6",
	"two":   "2",
	"four":  "4",
	"nine":  "9",
	"five":  "5",
	"seven": "7",
	"three": "3",
	"eight": "8",
}

// Get the length of the longest key in the map
func getMaxKeyLen[T any](input map[string]T) int {
	max := -1 // -1 is guaranteed to be less than length of string
	for key, _ := range input {
		if len(key) < max {
			// Skip shorter string
			continue
		}
		if len(key) > max {
			// Found longer string. Update max and reset result.
			max = len(key)
		}
	}
	return max
}

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

// Convert each word in the line to the corresponding number
// Using a sliding window of size maxSearchLen.
func convertWordsToNumbers(line string, maxSearchLen int) string {
	startIndex := 0
	result := ""
	// Iterate over each character in the line
	for startIndex < len(line) {
		// Copy the map to allow for changes
		mapCopy := make(map[string]string)
		for k, v := range strToNumMap {
			mapCopy[k] = v
		}

		// Iterate over each character in the line starting at startIndex
		for i := 0; i <= maxSearchLen; i++ {
			// If we are at the end of the line, break
			if startIndex+i > len(line)-1 {
				break
			}
			// Iterate over each key (number-word) in the mapCopy
			for k, v := range mapCopy {
				// If our current index is greater than the length of the key,
				// continue to the next key
				if i > len(k)-1 {
					continue
				}
				// If the current character in the line does not match the
				// current character in the key, delete the key from the mapCopy
				if k[i] != line[i+startIndex] {
					delete(mapCopy, k)
					continue
				}
				// If we are at the end of the key, add the value to the result
				// and break
				if i == len(k)-1 {
					result += v
					break
				}
			}
		}

		// Continue to the next character in the line
		startIndex++
	}

	return result
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d`)
	result := 0
	maxKeyLen := getMaxKeyLen(strToNumMap)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		fmt.Printf("%s -> ", line)
		line = convertWordsToNumbers(line, maxKeyLen)
		fmt.Println(line)
		digits := re.FindAllString(line, -1)
		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]
		num, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		fmt.Println(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		result += num
	}

	fmt.Printf("Result: %d\n", result)
}
