package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	puz := parseInput(input)
	for _, seq := range puz {
		result += interpolate(seq)
	}
	return result
}
func Part2(input string) (result int) {
	puz := parseInput(input)
	for _, seq := range puz {
		slices.Reverse(seq)
		result += interpolate(seq)
	}
	return result
}

func interpolate(in []int) int {
	diffs := [][]int{in}

	for {
		var cur []int
		for idx := 0; idx < len(diffs[len(diffs)-1])-1; idx++ {
			diff := diffs[len(diffs)-1][idx+1] - diffs[len(diffs)-1][idx]
			cur = append(cur, diff)
		}
		if len(cur) == 0 {
			break
		}
		diffs = append(diffs, cur)
	}

	// Thank you ChatGPT to help with this
	// My initial implmentation had issue with the summing
	for idx := len(diffs) - 2; idx >= 0; idx-- {
		for i := len(diffs[idx]) - 1; i > 0; i-- {
			diffs[idx][i] += diffs[idx+1][i-1]
		}
	}

	return diffs[0][len(diffs[0])-1]
}

func parseInput(input string) (res Puzzle) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var nums []int
		for _, num := range strings.Split(strings.TrimSpace(line), " ") {
			val, err := strconv.Atoi(num)
			if err == nil {
				nums = append(nums, val)
			}
		}
		res = append(res, nums)
	}
	return
}

type Puzzle [][]int

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
