package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	for _, line := range strings.Fields(input) {
		var filtered []string
		for _, c := range line {
			if !unicode.IsLetter(c) {
				filtered = append(filtered, string(c))
			}
		}
		num, _ := strconv.Atoi(filtered[0] + filtered[len(filtered)-1])
		result += num
	}
	return result
}

func Part2(input string) (result int) {
	for _, line := range strings.Fields(input) {
		var filtered []string
		for idx := 0; idx < len(line); idx++ {
			if !unicode.IsLetter(rune(line[idx])) {
				filtered = append(filtered, string(line[idx]))
			} else {
				for num, text := range numbers {
					if idx+len(text) <= len(line) && line[idx:idx+len(text)] == text {
						filtered = append(filtered, strconv.Itoa(num+1))
					}
				}
			}
		}
		num, _ := strconv.Atoi(filtered[0] + filtered[len(filtered)-1])
		result += num
	}
	return result
}

var (
	numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
