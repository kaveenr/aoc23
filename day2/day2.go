package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	for id, scene := range parseGame(input) {
		possible := true
		for _, try := range scene {
			for color, count := range try {
				if count > cubeCount[color] {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			result += id
		}
	}
	return result
}

func Part2(input string) (result int) {
	for _, scene := range parseGame(input) {
		counts := make(map[string][]int)
		for _, try := range scene {
			for color, count := range try {
				counts[color] = append(counts[color], count)
			}
		}
		result += slices.Max(counts["red"]) * slices.Max(counts["green"]) * slices.Max(counts["blue"])
	}
	return result
}

func parseGame(input string) PuzzleInput {
	PuzzleInput := make(PuzzleInput)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			break
		}
		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		sets := strings.Split(parts[1], ";")
		PuzzleInput[id] = make(Game, len(sets))
		for try, rolls := range sets {
			rollDetail := make(Roll)
			for _, roll := range strings.Split(strings.TrimSpace(rolls), ",") {
				rollParts := strings.Split(strings.TrimSpace(roll), " ")
				rollDetail[rollParts[1]], _ = strconv.Atoi(rollParts[0])
			}
			PuzzleInput[id][try] = rollDetail
		}
	}
	return PuzzleInput
}

type Roll = map[string]int
type Game = []Roll
type PuzzleInput = map[int]Game

var (
	cubeCount = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
