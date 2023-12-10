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

var cards = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
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

func getPlayType(hand string) PlayType {
	cardMap := map[rune]int{}

	for _, card := range hand {
		cardMap[card] += 1
	}

	var maxAmount int
	for _, amount := range cardMap {
		if amount > maxAmount {
			maxAmount = amount
		}
	}

	switch maxAmount {
	case 5:
		return FIVE_OF_KIND
	case 4:
		return FOUR_OF_KIND
	case 3:
		if len(cardMap) == 2 {
			return FULL_HOUSE
		}
		return THREE_OF_KIND
	case 2:
		if len(cardMap) == 3 {
			return TWO_PAIR
		}
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

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

type ByRank []Play

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	if a[i].playType < a[j].playType {
		return true
	} else if a[i].playType > a[j].playType {
		return false
	}
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
