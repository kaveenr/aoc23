package main

import (
	"fmt"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	puzz := parseInput(input)
	cur := "AAA"
	for {
		for _, dir := range puzz.Directions {
			if dir == 'L' {
				cur = puzz.Nodes[cur].Left
			} else {
				cur = puzz.Nodes[cur].Right
			}
		}
		result += len(puzz.Directions)
		if cur == "ZZZ" {
			break
		}
	}
	return
}

func Part2(input string) (result int) {
	return result
}

func parseInput(input string) (res Map) {
	lines := strings.Split(input, "\n")
	res.Directions = []rune(strings.TrimSpace(lines[0]))

	res.Nodes = make(map[string]MapNode)
	for _, line := range lines[2:] {
		parts := strings.Split(strings.TrimSpace(line), " = ")
		if len(parts) < 2 {
			break
		}
		nodeName := parts[0]
		res.Nodes[nodeName] = parseMapNode(parts[1])
	}

	return
}

func parseMapNode(nodeStr string) MapNode {
	nodeStr = strings.Trim(nodeStr, "()")
	parts := strings.Split(nodeStr, ", ")
	left := parts[0]
	right := parts[1]
	return MapNode{
		Left:  left,
		Right: right,
	}
}

type Map struct {
	Directions []rune
	Nodes      map[string]MapNode
}

type MapNode struct {
	Left  string
	Right string
}

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
