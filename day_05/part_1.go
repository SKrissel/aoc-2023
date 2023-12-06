// This day was a tough one. Even the description was hard to understand.
// After reading it a few times and still not understanding it, I decided
// to just start coding and see what happens. I started with the input
// parsing, which was pretty straightforward. I then started to implement
// the algorithm described in the description. By carefully reading the
// description and implementing it step by step, I was able to get the
// correct result.

package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
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

var isNumerical = regexp.MustCompile(`^[\d ]+$`)

type SeedMap struct {
	destStart   int
	sourceStart int
	mapRange    int
}

// Get the List of seeds from the first line of the input as a list of ints
func getSeeds(line string) []int {
	line = strings.ReplaceAll(line, "seeds: ", "")
	seedsStr := strings.Split(line, " ")

	var result []int
	for _, seed := range seedsStr {
		seedInt, err := strconv.Atoi(seed)
		handleError(err, "Atoi")
		result = append(result, seedInt)
	}
	return result
}

// Parse a single map from a string
func parseMap(line string) SeedMap {
	mapValues := strings.Split(line, " ")
	destStart, _ := strconv.Atoi(mapValues[0])
	sourceStart, _ := strconv.Atoi(mapValues[1])
	mapRange, _ := strconv.Atoi(mapValues[2])

	return SeedMap{
		destStart:   destStart,
		sourceStart: sourceStart,
		mapRange:    mapRange,
	}
}

// Get the maps from the input as a list of lists of SeedMaps
// where each entry in the outer list is a list of SeedMaps for a single map
func getMaps(lines []string) [][]SeedMap {
	// Skip first line because it contains the seeds
	index := 1
	var result [][]SeedMap
	var currentMap []SeedMap
	currentMap = nil

	for index < len(lines) {
		line := lines[index]
		// Skip empty lines
		if line == "\n" {
			index++
			continue
		}

		// If the line is not numerical, it's a title
		isTitle := !isNumerical.MatchString(line)
		if isTitle {
			// Finish the current map if there is one
			if currentMap != nil {
				result = append(result, currentMap)
			}
			currentMap = nil
			index++
			continue
		}

		// If the line is numerical, it's a map
		// Parse the map
		currentMap = append(currentMap, parseMap(line))
		index++

	}

	result = append(result, currentMap)
	return result
}

// Find the map containing a seed in a list of maps
func findCorrespondingMapValue(seedMaps []SeedMap, value int) int {
	// Unmapped seeds are their own value
	result := value

	var containingMap SeedMap
	for _, seedMap := range seedMaps {
		// If the seed is in the range of the map, it's mapped
		if seedMap.sourceStart <= value &&
			seedMap.sourceStart+seedMap.mapRange > value {
			containingMap = seedMap
			break
		}
	}

	// Range is never 0, so if the range is actually 0, no containing map
	// was found.
	if containingMap.mapRange != 0 {
		result = containingMap.destStart - containingMap.sourceStart + value
	}

	return result

}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")

	seeds := getSeeds(lines[0])
	maps := getMaps(lines)

	var results []int

	for _, seed := range seeds {
		currentValue := seed
		// For each Seed, go through the maps and find the corresponding value
		for _, seedMap := range maps {
			currentValue = findCorrespondingMapValue(seedMap, currentValue)
		}
		results = append(results, currentValue)
	}

	// Print the minimum result
	fmt.Printf("Result: %v\n", slices.Min(results))
}
