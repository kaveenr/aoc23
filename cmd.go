package main

import (
	"fmt"
	"os"
	"strconv"
	"ukr/aoc23/day1"
	"ukr/aoc23/day2"
)

func LoadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

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

	input, _ := LoadFile(fmt.Sprintf("inputs/day%d.txt", dayNumber))

	result1, result2 := 0, 0

	switch dayNumber {
	case 1:
		result1 = day1.Part1(input)
		result2 = day1.Part2(input)
	case 2:
		result1 = day2.Part1(input)
		result2 = day2.Part2(input)
	default:
		fmt.Printf("Implment Day %d please\n", dayNumber)
		return
	}
	fmt.Println("Part 1 Solution:", result1)
	fmt.Println("Part 2 Solution:", result2)
}
