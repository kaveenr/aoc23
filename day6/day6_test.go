package main

import (
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day6_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 288, result)
}

func Test_Day6_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 71503, result)
}

func Test_Day6_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 625968, result)
}

func Test_Day6_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 43663323, result)
}

func Benchmark_Day6_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(testInput)
	}
}

func Benchmark_Day6_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInput        = `Time:      7  15   30
	Distance:  9  40  200`
)
