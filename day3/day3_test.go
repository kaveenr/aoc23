package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day3_Part1(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 4361, result)
}

func Test_Day3_Part2(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 467835, result)
}

var testInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
