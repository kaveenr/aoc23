package main

import (
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day9_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 114, result)
}

func Test_Day9_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 2, result)
}

func Test_Day9_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 2043183816, result)
}

func Test_Day9_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 0, result)
}

func Benchmark_Day9_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(testInput)
	}
}

func Benchmark_Day9_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInput        = `0 3 6 9 12 15
	1 3 6 10 15 21
	10 13 16 21 30 45`
)
