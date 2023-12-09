package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kaveenr/aoc23/commons"
	"golang.org/x/exp/maps"
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
	puzz := parseInput(input)
	var startNodes, endNodes []string
	for _, node := range maps.Keys(puzz.Nodes) {
		if node[2] == 'A' {
			startNodes = append(startNodes, node)
		} else if node[2] == 'Z' {
			endNodes = append(endNodes, node)
		}
	}
	all := slices.Clone(startNodes)
	for solved := 0; solved != len(endNodes); {
		for _, dir := range puzz.Directions {
			for idx, cur := range all {
				if dir == 'L' {
					all[idx] = puzz.Nodes[cur].Left
				} else {
					all[idx] = puzz.Nodes[cur].Right
				}
			}
		}
		result += len(puzz.Directions)
		solved = 0
		for _, node := range all {
			if node[2] == 'Z' {
				solved++
			}
		}
	}
	return
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
	parts := strings.Split(strings.Trim(nodeStr, "()"), ", ")
	return MapNode{
		Left:  parts[0],
		Right: parts[1],
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
