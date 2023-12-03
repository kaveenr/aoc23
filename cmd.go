package main

import (
	"fmt"
	"os"
	"strconv"
	"ukr/aoc23/commons"
	"ukr/aoc23/day1"
	"ukr/aoc23/day2"
	"ukr/aoc23/day3"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run . <day_number>")
		return
	}

	dayNumber, err := strconv.Atoi(args[0])
	if err != nil || dayNumber < 1 || dayNumber > 25 {
		fmt.Println("Invalid day number. Please provide a day number between 1 and 25.")
		return
	}

	input, _ := commons.LoadFile(fmt.Sprintf("inputs/day%d.txt", dayNumber))

	result1, result2 := 0, 0

	switch dayNumber {
	case 1:
		result1, result2 = day1.Part1(input), day1.Part2(input)
	case 2:
		result1, result2 = day2.Part1(input), day2.Part2(input)
	case 3:
		result1, result2 = day3.Part1(input), day3.Part2(input)
	default:
		fmt.Printf("Implment Day %d please\n", dayNumber)
		return
	}
	fmt.Println("Solution for day", dayNumber)
	fmt.Println("Part 1 Solution:", result1)
	fmt.Println("Part 2 Solution:", result2)
}
