package day1

import (
	"testing"
	"ukr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day1_Part1_Example(t *testing.T) {
	result := Part1(`1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`)
	assert.Equal(t, 142, result)
}

func Test_Day1_Part2_Example(t *testing.T) {
	result := Part2(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`)
	assert.Equal(t, 281, result)
}

func Test_Day1_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 55002, result)
}

func Test_Day1_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 55093, result)
}

var myPuzzleInput, _ = commons.LoadFile(`../inputs/day1.txt`)
