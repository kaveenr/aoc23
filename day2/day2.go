package day2

import (
	"slices"
	"strconv"
	"strings"
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
		counts := map[string][]int{
			"red":   make([]int, 0),
			"green": make([]int, 0),
			"blue":  make([]int, 0),
		}
		for _, try := range scene {
			for color, count := range try {
				counts[color] = append(counts[color], count)
			}
		}
		result += slices.Max(counts["red"]) * slices.Max(counts["green"]) * slices.Max(counts["blue"])
	}
	return result
}

func parseGame(input string) games {
	games := make(games)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		sets := strings.Split(parts[1], ";")
		games[id] = make(game, len(sets))
		for try, rolls := range sets {
			rollDetail := make(gameRoll)
			for _, roll := range strings.Split(strings.TrimSpace(rolls), ",") {
				rollParts := strings.Split(strings.TrimSpace(roll), " ")
				rollDetail[rollParts[1]], _ = strconv.Atoi(rollParts[0])
			}
			games[id][try] = rollDetail
		}
	}
	return games
}

type gameRoll = map[string]int
type game = []gameRoll
type games = map[int]game

var (
	cubeCount = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)
