// This one required a bit of critical thinking. At first, I thought I could
// just use the same algorithm as in part 1, which did work just fine for the
// given example. However, when I ran the program with the actual input, I
// quickly realized that this would not work because even after over
// 2.370.000.000 Steps, the program still did not find a solution.
// I then realized that I had to find a way to detect cycles in the Path for
// each Node. I decided to store the steps at which a cycle was found and then
// calculate the LCM of all cycles to get the result.
// After getting the result, it was clear that this optimization was necessary
// because otherwise the program would have taken approximately 25 days to
// complete. With the optimization, it only took 0.3 seconds.
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

// Shamelessly stolen from here: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Shamelessly stolen (and adapted) from here: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type Node struct {
	left  string
	right string
}

// Regex to parse the nodes from the input in the form:
// AAA = (BBB, CCC)
var reNode = regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)

// Parse the nodes from the input into a map of string -> Node
// Also return a list of starting nodes (nodes that end with 'A')
func parseNodes(lines []string) (map[string]Node, []string) {
	result := make(map[string]Node)
	var startingNodes []string

	// Skip the first two lines because they contain the steps and a blank line
	for i := 2; i <= len(lines)-1; i++ {
		line := lines[i]

		// Parse the line with the regex and create a Node
		match := reNode.FindStringSubmatch(line)
		nodeName := match[1]
		node := Node{
			left:  match[2],
			right: match[3],
		}

		// Add the Node to the map
		result[nodeName] = node

		// If the Node Name ends with an 'A', it is a starting Node.
		if nodeName[len(nodeName)-1] == 'A' {
			startingNodes = append(startingNodes, nodeName)
		}
	}

	return result, startingNodes
}

// Check if all cycles have been found. This is the case if all cycles are
// greater than 0.
// If yes, return true, else false
func allCyclesFound(cycles []int) bool {
	for _, cycle := range cycles {
		if cycle == 0 {
			return false
		}
	}
	fmt.Println("All cycles found!")

	return true
}

// Check if a cycle has been found for each Node and if yes, store it in the
// cycles array.
func checkCycles(nodes []string, cycles []int, steps int) []int {
	for i, node := range nodes {
		// If the current node and the starting node end with 'Z', a cycle has
		// been found.
		if nodes[0][len(node)-1] == 'Z' && node[len(node)-1] == 'Z' {
			// If the cycle has not been found yet, store it in the cycles array
			// This is necessary because the cycle is found multiple times
			// because the path is circular.
			// The first time the cycle is found is the shortest path.
			if cycles[i] == 0 {
				cycles[i] = steps
				fmt.Printf("Found cycle %v/%v at Node '%v': %v\n", i+1, len(cycles), node, steps)
			}
		}
	}

	return cycles
}

func main() {
	// Read the input file
	input := getInputText()
	lines := strings.Split(input, "\n")
	steps := 0

	path := lines[0]
	nodes, currentNodes := parseNodes(lines)
	pathCycles := make([]int, len(currentNodes))
	fmt.Printf("Starting with nodes: %v\n", currentNodes)

	// Follow the path until all Nodes end with 'Z'
	for !allCyclesFound(pathCycles) {
		// Get the current step from the path. If we reach the end of the path,
		// start from the beginning again.
		currentStep := path[steps%len(path)]

		// Go through each current node and get the next node
		for i, node := range currentNodes {
			// Get the next node from the current node and the current step
			currentNodePath := nodes[node]
			if currentStep == 'R' {
				currentNodes[i] = currentNodePath.right
			} else {
				currentNodes[i] = currentNodePath.left
			}
		}

		steps++

		checkCycles(currentNodes, pathCycles, steps)
	}

	// Calculate the LCM of all cycles to get the result
	result := LCM(pathCycles...)

	fmt.Printf("Result: %v\n", result)
}
