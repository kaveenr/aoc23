package main

import (
	"fmt"
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day1_Part1_Example(t *testing.T) {
	result := Part1(testInputPart1)
	assert.Equal(t, 142, result)
}

func Test_Day1_Part2_Example(t *testing.T) {
	result := Part2(testInputPart2)
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

func Benchmark_Day1_Part1(b *testing.B) {
	for _, v := range []string{testInputPart1, myPuzzleInput} {
		b.Run(fmt.Sprintf("input_size_%d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Part1(v)
			}
		})
	}
}

func Benchmark_Day1_Part2(b *testing.B) {
	for _, v := range []string{testInputPart2, myPuzzleInput} {
		b.Run(fmt.Sprintf("input_size_%d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Part2(v)
			}
		})
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInputPart1   = `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	testInputPart2 = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`
)
