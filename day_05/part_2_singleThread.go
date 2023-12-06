// After completing part 1, part 2 seemed pretty straightforward. I just had
// to add a loop to parse a range of seeds and then find the minimum value
// instead of storing all values.
// This however, was not the case. The program ran for a few minutes and then
// Gave a wrong result. After some debugging, I found out that the code from
// part 1 was not working correctly. Especially the part where I was searching
// for the map containing a seed (findCorrespondingMapValue()). I had to
// change the condition from
// seedMap.sourceStart < value
// to
// seedMap.sourceStart <= value
// to make it work. I don't know why it worked in part 1, but I guess it was
// just luck.
//
// After fixing this, the program ran for a few minutes and gave the correct
// result.
//
// This however, was not the end of the story. I was not satisfied with the
// runtime of the program. I thought that it could be improved by using multiple
// threads especially because I am using go, which has great support for
// concurrency.

package main

import (
	"fmt"
	"math"
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

var isNumerical = regexp.MustCompile(`^[\d ]+$`)

type SeedMap struct {
	destStart   int
	sourceStart int
	mapRange    int
}

type SeedRange struct {
	start  int
	length int
}

// Get the List of seeds from the first line of the input as a list of ints
func getSeeds(line string) []SeedRange {
	line = strings.ReplaceAll(line, "seeds: ", "")
	seedsStr := strings.Split(line, " ")

	var result []SeedRange
	index := 0
	for index < len(seedsStr) {
		start, _ := strconv.Atoi(seedsStr[index])
		length, _ := strconv.Atoi(seedsStr[index+1])
		resultSeed := SeedRange{
			start:  start,
			length: length,
		}

		result = append(result, resultSeed)
		index += 2
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

	currentMin := math.MaxInt
	for step, seedRange := range seeds {
		fmt.Printf("Step %v / %v\n", step+1, len(seeds))

		// Go through each seed in the range
		seedRangeIndex := seedRange.start
		for seedRangeIndex <= seedRange.start+seedRange.length-1 {
			currentValue := seedRangeIndex
			// For each Seed, go through the maps and find the corresponding value
			for _, seedMap := range maps {
				currentValue = findCorrespondingMapValue(seedMap, currentValue)
			}

			// Do not save all results, just the minimum
			if currentValue < currentMin {
				currentMin = currentValue
			}
			seedRangeIndex++
		}
	}

	// Print the minimum result
	fmt.Printf("Result: %v\n", currentMin)
}
