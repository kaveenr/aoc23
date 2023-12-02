package main

import (
	"fmt"
	"os"
	"strconv"
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

	input, _ := LoadFile(fmt.Sprintf("inputs/day%d.txt", dayNumber))

	switch dayNumber {
	case 1:
		resultP1 := Day1_Part1(input)
		fmt.Println("Part 1 Solution:", resultP1)

		// Solve Part 2
		resultP2 := Day1_Part2(input)
		fmt.Println("Part 2 Solution:", resultP2)
	default:
		fmt.Printf("Implment Day %d please\n", dayNumber)
	}
}
