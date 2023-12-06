package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	races := parseInput(input)
	result = 1
	for _, race := range races {
		var wins int
		for holdTime := 1; holdTime < race.Time; holdTime++ {
			distTravel := holdTime * (race.Time - holdTime)
			if distTravel > race.Distance {
				wins++
			}
		}
		result *= wins
	}
	return result
}

func Part2(input string) (result int) {
	race := parseInputB(input)
	for holdTime := 1; holdTime < race.Time; holdTime++ {
		distTravel := holdTime * (race.Time - holdTime)
		if distTravel > race.Distance {
			result++
		}
	}
	return result
}

func parseInput(input string) (races Races) {
	parts := strings.Split(input, "\n")
	times, dists := make([]int, 0), make([]int, 0)
	for _, part := range strings.Split(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ") {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			times = append(times, num)
		}
	}
	for _, part := range strings.Split(strings.TrimSpace(strings.Split(parts[1], ":")[1]), " ") {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			dists = append(dists, num)
		}
	}
	for idx, time := range times {
		races = append(races, Race{
			Time:     time,
			Distance: dists[idx],
		})
	}
	return races
}

func parseInputB(input string) (race Race) {
	parts := strings.Split(input, "\n")
	time := strings.ReplaceAll(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ", "")
	dist := strings.ReplaceAll(strings.TrimSpace(strings.Split(parts[1], ":")[1]), " ", "")
	parsedTime, _ := strconv.Atoi(time)
	parsedDist, _ := strconv.Atoi(dist)
	return Race{
		Time:     parsedTime,
		Distance: parsedDist,
	}
}

type Races []Race

type Race struct {
	Time     int
	Distance int
}

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
