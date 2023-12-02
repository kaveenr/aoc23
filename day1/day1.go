package day1

import (
	"strconv"
	"strings"
	"unicode"
)

func Part1(input string) (result int) {
	for _, line := range strings.Fields(input) {
		var filteredLine []rune
		for _, c := range line {
			if !unicode.IsLetter(c) {
				filteredLine = append(filteredLine, c)
			}
		}
		num, _ := strconv.Atoi(string(filteredLine[0]) + string(filteredLine[len(filteredLine)-1]))
		result += num
	}
	return result
}

func Part2(input string) (result int) {
	for _, line := range strings.Fields(input) {
		parsed := []string{}
		for idx := 0; idx < len(line); idx++ {
			if !unicode.IsLetter(rune(line[idx])) {
				parsed = append(parsed, string(line[idx]))
			} else {
				for num, text := range numbers {
					if idx+len(text) <= len(line) && line[idx:idx+len(text)] == text {
						parsed = append(parsed, strconv.Itoa(num+1))
					}
				}
			}
		}
		if len(parsed) > 0 {
			num, _ := strconv.Atoi(parsed[0] + parsed[len(parsed)-1])
			result += num
		}
	}
	return result
}

var (
	numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)
