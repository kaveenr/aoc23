package dayn

import (
	"testing"
	"ukr/aoc23/commons"

	"github.com/stretchr/testify/assert"
)

func Test_Dayn_Part1_Example(t *testing.T) {
	result := Part1(``)
	assert.Equal(t, 0, result)
}

func Test_Dayn_Part2_Example(t *testing.T) {
	result := Part2(``)
	assert.Equal(t, 0, result)
}

func Test_Dayn_Part1(t *testing.T) {
	result := Part1(myPuzzleInput)
	assert.Equal(t, 0, result)
}

func Test_Dayn_Part2(t *testing.T) {
	result := Part2(myPuzzleInput)
	assert.Equal(t, 0, result)
}

var myPuzzleInput, _ = commons.LoadFile(`../inputs/dayn.txt`)
