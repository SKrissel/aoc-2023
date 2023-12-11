// After returning to this part 2 after doing day 11, I was pleasantly surprised
// of how easy this part was. Just change the worth of the Joker-card to -1 and
// Add 6 Lines of code to the getPlayType-function.
// Easy as that.
package main

import (
	"fmt"
	"os"
	"sort"
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

// This is one play of the game
type Play struct {
	hand     string
	bid      int
	playType PlayType
}

type PlayType int

const (
	FIVE_OF_KIND  PlayType = 6
	FOUR_OF_KIND  PlayType = 5
	FULL_HOUSE    PlayType = 4
	THREE_OF_KIND PlayType = 3
	TWO_PAIR      PlayType = 2
	ONE_PAIR      PlayType = 1
	HIGH_CARD     PlayType = 0
)

// Jokers are now the lowest card
var cards = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': -1,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

// Get the play type of a hand
func getPlayType(hand string) PlayType {
	cardMap := map[rune]int{}

	// Count the amount of cards of the same type
	for _, card := range hand {
		cardMap[card] += 1
	}

	// Jokers are wild
	// Remove them from the map
	js := cardMap['J']
	delete(cardMap, 'J')

	// If there are no cards left, it's a five of a kind with jokers
	if len(cardMap) == 0 {
		return FIVE_OF_KIND
	}

	// Find the max amount of cards of the same type
	var maxAmount int
	for _, amount := range cardMap {
		if amount > maxAmount {
			maxAmount = amount
		}
	}
	// The Jokers are wild, so add them to the max amount
	maxAmount += js

	switch maxAmount {
	case 5:
		return FIVE_OF_KIND
	case 4:
		return FOUR_OF_KIND
	case 3:
		// If there are max 3 cards of the same type
		// AND there are 2 different types of cards it is a full house
		if len(cardMap) == 2 {
			return FULL_HOUSE
		}
		return THREE_OF_KIND
	case 2:
		// If there are max 2 cards of the same type
		// AND there are 3 different types of cards it is a two pair
		if len(cardMap) == 3 {
			return TWO_PAIR
		}
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

// Get all plays from the input
func getPlays(lines []string) []Play {
	var result []Play
	for _, line := range lines {
		splitStr := strings.Split(line, " ")
		bid, _ := strconv.Atoi(splitStr[1])
		hand := splitStr[0]
		play := Play{
			bid:      bid,
			hand:     hand,
			playType: getPlayType(hand),
		}
		result = append(result, play)
	}
	return result
}

// Sort plays by rank
type ByRank []Play

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	// Return true if a[i] < a[j]

	// Sort by play type first
	if a[i].playType < a[j].playType {
		return true
	} else if a[i].playType > a[j].playType {
		return false
	}

	// If the play types are equal, sort by hand
	for index := range a[i].hand {
		cardI := []rune(a[i].hand)[index]
		cardJ := []rune(a[j].hand)[index]
		if cards[cardI] < cards[cardJ] {
			return true
		} else if cards[cardI] > cards[cardJ] {
			return false
		}
	}
	return false
}

func main() {
	input := getInputText()
	lines := strings.Split(input, "\n")

	plays := getPlays(lines)
	sort.Sort(ByRank(plays))

	winnings := 0
	for i, play := range plays {
		fmt.Printf("%v: %v\n", i+1, play)
		winnings += play.bid * (i + 1)
	}

	fmt.Printf("Result: %v\n", winnings)
}
