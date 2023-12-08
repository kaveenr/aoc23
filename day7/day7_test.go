package main

import (
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day7_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 6440, result)
}

func Test_Day7_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 5905, result)
}

func Test_Day7_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 248569531, result)
}

func Test_Day7_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 250382098, result)
}

func Benchmark_Day7_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(testInput)
	}
}

func Benchmark_Day7_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInput        = `32T3K 765
	T55J5 684
	KK677 28
	KTJJT 220
	QQQJA 483`
)
