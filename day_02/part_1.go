// This one is simple as well. Just read the file, split it into lines, and
// iterate over each line. For each line, find the Game number and the plays.
// Then, construct a Game struct checking if the game is possible. If it is,
// add the game number to the result. Print the result.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Game struct to hold the game number and the plays
type Game struct {
	number int
	plays  []Play
}

// Play struct to hold the number of dice of each color
type Play struct {
	red   int
	green int
	blue  int
}

// Number of available dice of each color
var availableDice = Play{
	red:   12,
	green: 13,
	blue:  14,
}

// Regexes to find the game number and the number of dice of each color
var reGame = regexp.MustCompile(`Game (\d+): `)
var reRed = regexp.MustCompile(`(\d+) red`)
var reGreen = regexp.MustCompile(`(\d+) green`)
var reBlue = regexp.MustCompile(`(\d+) blue`)

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

// Get the number of dice of a color from a play by matching a regex
func getColorDice(play string, re *regexp.Regexp) int {
	result := 0
	match := re.FindStringSubmatch(play)
	// If the regex matches, the first element of the result is the whole
	// string, and the second element is the number of dice
	// If the regex doesn't match, the result is nil and no dice of that color
	// are used
	if match != nil {
		result, _ = strconv.Atoi(match[1])
	}
	return result
}

// Parse a play from a string
func parsePlay(play string) Play {
	return Play{
		red:   getColorDice(play, reRed),
		green: getColorDice(play, reGreen),
		blue:  getColorDice(play, reBlue),
	}
}

// Parse the game number and the plays from a line of input
func parseGame(line string) Game {
	// Find the game number
	gameInfo := reGame.FindStringSubmatch(line)
	gameNr, _ := strconv.Atoi(gameInfo[1])

	// Remove the game number from the line
	line = strings.Replace(line, gameInfo[0], "", 1)

	playsStr := strings.Split(line, ";")

	// Parse each play
	var plays []Play
	for _, playStr := range playsStr {
		plays = append(plays, parsePlay(playStr))
	}

	result := Game{number: gameNr, plays: plays}
	return result
}

// Check if a game is possible by checking each play against the available
// dice of each color
func gamePossible(game Game) bool {
	for _, play := range game.plays {
		if play.red > availableDice.red ||
			play.green > availableDice.green ||
			play.blue > availableDice.blue {
			return false
		}
	}
	return true
}

func main() {
	// Read the input file
	input := getInputText()
	lines := strings.Split(input, "\n")
	result := 0

	// Iterate over each line
	for _, line := range lines {
		game := parseGame(line)
		// If the game is possible, add the game number to the result
		if gamePossible(game) {
			result += game.number
		}
	}

	fmt.Printf("Result: %v\n", result)
}
