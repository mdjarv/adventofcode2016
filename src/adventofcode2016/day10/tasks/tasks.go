package tasks

import (
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	id      string
	chips   []int
	outLow  string
	outHigh string
}

func getLow(node Node) int {
	if node.chips[0] < node.chips[1] {
		return node.chips[0]
	} else {
		return node.chips[1]
	}
}

func getHigh(node Node) int {
	if node.chips[0] > node.chips[1] {
		return node.chips[0]
	} else {
		return node.chips[1]
	}
}

var foundAt = ""

func balanceNodes(nodes map[string]Node, output map[string]int, ids []string, chip1 int, chip2 int) (map[string]Node, map[string]int) {
	if len(ids) == 0 {
		return nodes, output
	}

	changedIds := []string{}

	for _, id := range ids {
		startNode := nodes[id]

		if isComplete(startNode.chips, chip1, chip2) {
			foundAt = startNode.id
		}

		if len(startNode.chips) == 2 {
			lowNode := nodes[startNode.outLow]
			highNode := nodes[startNode.outHigh]

			lowNode.chips = append(lowNode.chips, getLow(startNode))
			highNode.chips = append(highNode.chips, getHigh(startNode))

			nodes[lowNode.id] = lowNode
			if isOutput(startNode.outLow) {
				output[startNode.outLow] = getLow(startNode)
			} else {
				changedIds = append(changedIds, lowNode.id)
			}

			nodes[highNode.id] = highNode
			if isOutput(startNode.outHigh) {
				output[startNode.outHigh] = getHigh(startNode)
			} else {
				changedIds = append(changedIds, highNode.id)
			}

			startNode.chips = []int{}
			nodes[startNode.id] = startNode
		}
	}

	return balanceNodes(nodes, output, changedIds, chip1, chip2)
}

func isComplete(chips []int, chip1 int, chip2 int) bool {
	if len(chips) < 2 {
		return false
	}

	return (chips[0] == chip1 && chips[1] == chip2) || (chips[1] == chip1 && chips[0] == chip2)
}

func isOutput(id string) bool {
	s := strings.Split(id, " ")
	return s[0] == "output"
}

func ProcessLines(lines []string, chip1 int, chip2 int) (map[string]Node, map[string]int) {
	nodes := map[string]Node{}
	output := map[string]int{}

	valueRegex := regexp.MustCompile(`value (?P<value>\d+) goes to (?P<bot>bot \d+)`)
	givesRegex := regexp.MustCompile(`(?P<origin>bot \d+) gives low to (?P<low>\w+ \d+) and high to (?P<high>\w+ \d+)`)

	// Setup rules
	for _, line := range lines {
		givesSubstrings := givesRegex.FindStringSubmatch(line)
		if len(givesSubstrings) > 0 {
			id := givesSubstrings[1]
			lowDest := givesSubstrings[2]
			highDest := givesSubstrings[3]

			nodes[id] = Node{id, []int{}, lowDest, highDest}
			if isOutput(lowDest) {
				nodes[lowDest] = Node{lowDest, []int{}, "-", "-"}
			}
			if isOutput(highDest) {
				nodes[highDest] = Node{highDest, []int{}, "-", "-"}
			}
		}
	}
	for foundAt == "" {
		for _, line := range lines {
			valueSubstrings := valueRegex.FindStringSubmatch(line)

			if len(valueSubstrings) > 0 {
				id := valueSubstrings[2]
				value, _ := strconv.Atoi(valueSubstrings[1])

				node := nodes[id]
				node.chips = append(node.chips, value)
				nodes[id] = node
				nodes, output = balanceNodes(nodes, output, []string{id}, chip1, chip2)
			}
		}
	}

	return nodes, output
}

func Day10Part1(lines []string, chip1 int, chip2 int) string {
	foundAt = ""
	ProcessLines(lines, chip1, chip2)
	return foundAt
}

func Day10Part2(lines []string, chip1 int, chip2 int) int {
	foundAt = ""
	_, output := ProcessLines(lines, chip1, chip2)
	return output["output 0"] * output["output 1"] * output["output 2"]
}
