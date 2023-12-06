package main

import (
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day5_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 35, result)
}

func Test_Day5_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 46, result)
}

func Test_Day5_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 84470622, result)
}

func Test_Day5_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 26714516, result)
}

func Benchmark_Day5_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(testInput)
	}
}

func Benchmark_Day5_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(testInput)
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInput        = `seeds: 79 14 55 13

	seed-to-soil map:
	50 98 2
	52 50 48
	
	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15
	
	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4
	
	water-to-light map:
	88 18 7
	18 25 70
	
	light-to-temperature map:
	45 77 23
	81 45 19
	68 64 13
	
	temperature-to-humidity map:
	0 69 1
	1 0 69
	
	humidity-to-location map:
	60 56 37
	56 93 4`
)
