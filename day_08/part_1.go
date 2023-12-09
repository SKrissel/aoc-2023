// This one was very straightforward. The main 'challenge' was to parse the
// Nodes from the input into a map of string -> Node. I used a regex to parse
// the input. Then we just have to follow the path via the steps in the first
// line of the input until we get from AAA to ZZZ and count the steps needed.
package main

import (
	"fmt"
	"os"
	"regexp"
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

type Node struct {
	left  string
	right string
}

// Regex to parse the nodes from the input in the form:
// AAA = (BBB, CCC)
var reNode = regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)

// Parse the nodes from the input into a map of string -> Node
func parseNodes(lines []string) map[string]Node {
	result := make(map[string]Node)

	// Skip the first two lines because they contain the steps and a blank line
	for i := 2; i <= len(lines)-1; i++ {
		line := lines[i]

		// Parse the line with the regex and create a Node
		match := reNode.FindStringSubmatch(line)
		node := Node{
			left:  match[2],
			right: match[3],
		}

		// Add the Node to the map
		result[match[1]] = node
	}

	return result
}

func main() {
	// Read the input file
	input := getInputText()
	lines := strings.Split(input, "\n")
	steps := 0

	path := lines[0]
	nodes := parseNodes(lines)

	// The starting node is alawys AAA
	currentNode := "AAA"
	// Follow the path until we get to ZZZ
	for currentNode != "ZZZ" {
		// Get the current step from the path. If we reach the end of the path,
		// start from the beginning again.
		currentStep := path[steps%len(path)]

		// Get the next node from the current node and the current step
		currentNodePath := nodes[currentNode]
		if currentStep == 'R' {
			currentNode = currentNodePath.right
		} else {
			currentNode = currentNodePath.left
		}
		steps++
	}

	fmt.Printf("Result: %v\n", steps)
}
