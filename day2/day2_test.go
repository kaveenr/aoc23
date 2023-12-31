package main

import (
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day2_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 8, result)
}

func Test_Day2_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 2286, result)
}

func Test_Day2_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 2207, result)
}

func Test_Day2_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 62241, result)
}

func Benchmark_Day2_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(testInput)
	}
}

func Benchmark_Day2_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInput        = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	`
)
