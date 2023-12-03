package day3

import (
	"fmt"
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day3_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 4361, result)
}

func Test_Day3_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 467835, result)
}

func Test_Day3_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 539433, result)
}

func Test_Day3_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 75847567, result)
}

func Benchmark_Day3_Part1(b *testing.B) {
	for _, v := range []string{testInput, myPuzzleInput} {
		b.Run(fmt.Sprintf("input_size_%d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Part1(v)
			}
		})
	}
}

func Benchmark_Day3_Part2(b *testing.B) {
	for _, v := range []string{testInput, myPuzzleInput} {
		b.Run(fmt.Sprintf("input_size_%d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Part2(v)
			}
		})
	}
}

var myPuzzleInput, _ = commons.LoadFile(`../inputs/day3.txt`)
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
