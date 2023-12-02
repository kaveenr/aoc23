package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Day1_Part1(input string) int {
	words := strings.Fields(input)
	var b [][]rune
	for _, line := range words {
		var filteredLine []rune
		for _, c := range line {
			if !unicode.IsLetter(c) {
				filteredLine = append(filteredLine, c)
			}
		}
		b = append(b, filteredLine)
	}

	var acc []int
	for _, i := range b {
		if len(i) > 0 {
			num, _ := strconv.Atoi(string(i[0]) + string(i[len(i)-1]))
			acc = append(acc, num)
		}
	}
	sumP1 := 0
	for _, val := range acc {
		sumP1 += val
	}
	return sumP1
}

func Day1_Part2(input string) int {
	lookup := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	words := strings.Fields(input)
	acc := 0

	for _, line := range words {
		parsed := []string{}
		for idx := 0; idx < len(line); idx++ {
			if !unicode.IsLetter(rune(line[idx])) {
				parsed = append(parsed, string(line[idx]))
			} else {
				for num, text := range lookup {
					if idx+len(text) <= len(line) && line[idx:idx+len(text)] == text {
						parsed = append(parsed, strconv.Itoa(num+1))
						idx += len(text) - 1
						break
					}
				}
			}
		}
		if len(parsed) > 0 {
			num, _ := strconv.Atoi(parsed[0] + parsed[len(parsed)-1])
			acc += num
		}
	}
	return acc
}
