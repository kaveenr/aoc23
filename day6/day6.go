package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	result = 1
	for _, race := range parseInput(input) {
		result *= solveRace(race)
	}
	return
}

func Part2(input string) (result int) {
	race := parseInputB(input)
	result += solveRace(race)
	return
}

func solveRace(race Race) (result int) {
	for holdTime := 1; holdTime < race.Time; holdTime++ {
		distTravel := holdTime * (race.Time - holdTime)
		if distTravel > race.Distance {
			result++
		}
	}
	return
}

func parseInput(input string) (races Races) {
	r, _ := regexp.Compile("([0-9]+)")
	parts := strings.Split(input, "\n")
	times, dists := r.FindAllString(parts[0], -1), r.FindAllString(parts[1], -1)
	for idx, time := range times {
		parsedTime, _ := strconv.Atoi(time)
		parsedDist, _ := strconv.Atoi(dists[idx])
		races = append(races, Race{
			Time:     parsedTime,
			Distance: parsedDist,
		})
	}
	return races
}

func parseInputB(input string) (race Race) {
	r, _ := regexp.Compile("([0-9]+)")
	parts := strings.Split(input, "\n")
	parsedTime, _ := strconv.Atoi(strings.Join(r.FindAllString(parts[0], -1), ""))
	parsedDist, _ := strconv.Atoi(strings.Join(r.FindAllString(parts[1], -1), ""))
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
