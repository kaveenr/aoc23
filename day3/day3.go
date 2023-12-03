package day3

import (
	"strconv"
	"strings"
	. "ukr/aoc23/commons"
	"unicode"
)

func Part1(input string) (result int) {
	sc := parseGrid(input)
	for rowIdx := 0; rowIdx < sc.Rows(); rowIdx++ {
		for colIdx := 0; colIdx < sc.Cols(); colIdx++ {
			cur := NewCoord(rowIdx, colIdx)
			if unicode.IsNumber(sc.Get(cur)) {
				ref := make([]rune, 0)
				isAdjecent := false
				var number int
				for scanIdx := cur.Col; scanIdx < sc.Cols(); scanIdx++ {
					ahead := NewCoord(cur.Row, scanIdx)
					if !unicode.IsNumber(sc.Get(ahead)) {
						colIdx = scanIdx
						break
					}
					curAdjecent, _ := checkIfAdjecent(sc, matchAllComponents, ahead)
					isAdjecent = curAdjecent || isAdjecent
					ref = append(ref, sc.Get(ahead))
					number, _ = strconv.Atoi(string(ref))
				}
				if isAdjecent {
					result += number
				}
			}
		}
	}

	return result
}

func Part2(input string) (result int) {
	sc := parseGrid(input)
	gearMap := make(map[Coord][]int)
	for rowIdx := 0; rowIdx < sc.Rows(); rowIdx++ {
		for colIdx := 0; colIdx < sc.Cols(); colIdx++ {
			cur := NewCoord(rowIdx, colIdx)
			var currentSymbol Coord
			if unicode.IsNumber(sc.Get(cur)) {
				ref := make([]rune, 0)
				isAdjecent := false
				var currentRatio int
				for scanIdx := cur.Col; scanIdx < sc.Cols(); scanIdx++ {
					ahead := NewCoord(cur.Row, scanIdx)
					if !unicode.IsNumber(sc.Get(ahead)) {
						colIdx = scanIdx
						break
					}
					curAdjecent, curSymbol := checkIfAdjecent(sc, matchGearComponents, ahead)
					if curAdjecent {
						currentSymbol = curSymbol
					}
					isAdjecent = curAdjecent || isAdjecent
					ref = append(ref, sc.Get(ahead))
					currentRatio, _ = strconv.Atoi(string(ref))
				}
				if isAdjecent {
					gearMap[currentSymbol] = append(gearMap[currentSymbol], currentRatio)
				}
			}
		}
	}
	for _, values := range gearMap {
		if len(values) == 2 {
			result += values[0] * values[1]
		}
	}
	return result
}

func parseGrid(input string) (res Grid) {
	res = make(Grid, 0)
	for _, field := range strings.Split(strings.TrimSpace(input), "\n") {
		field := strings.TrimSpace(field)
		line := make([]rune, len(field))
		for idx, char := range field {
			line[idx] = char
		}
		res = append(res, line)
	}
	return res
}

func checkIfAdjecent(sc Grid, match ComponentMacher, check Coord) (bool, Coord) {
	for checkRowIdx := check.Row - 1; checkRowIdx <= check.Row+1; checkRowIdx++ {
		for checkColIdx := check.Col - 1; checkColIdx <= check.Col+1; checkColIdx++ {
			if (checkRowIdx >= 0 && checkRowIdx < sc.Rows()) && (checkColIdx >= 0 && checkColIdx < sc.Cols()) && match(sc[checkRowIdx][checkColIdx]) {
				return true, NewCoord(checkRowIdx, checkColIdx)
			}
		}
	}
	return false, NewCoord(-1, -1)
}

type ComponentMacher func(cmp rune) bool

func matchAllComponents(cmp rune) bool {
	return !unicode.IsNumber(cmp) && cmp != '.'
}

func matchGearComponents(cmp rune) bool {
	return cmp == '*'
}
