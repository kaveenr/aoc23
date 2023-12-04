package main

import (
	"fmt"
	"testing"

	"github.com/kaveenr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Day4_Part1_Example(t *testing.T) {
	result := Part1(testInput)
	assert.Equal(t, 13, result)
}

func Test_Day4_Part2_Example(t *testing.T) {
	result := Part2(testInput)
	assert.Equal(t, 30, result)
}

func Test_Day4_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 23941, result)
}

func Test_Day4_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 5571760, result)
}
func Benchmark_Day4_Part1(b *testing.B) {
	for _, v := range []string{testInput, myPuzzleInput} {
		b.Run(fmt.Sprintf("input_size_%d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Part1(v)
			}
		})
	}
}

func Benchmark_Day4_Part2(b *testing.B) {
	for _, v := range []string{testInput, myPuzzleInput} {
		b.Run(fmt.Sprintf("input_size_%d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Part2(v)
			}
		})
	}
}

var (
	myPuzzleInput, _ = commons.LoadFile(`input.txt`)
	testInput        = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
    Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
    Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
    Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
    Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
    Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
)
