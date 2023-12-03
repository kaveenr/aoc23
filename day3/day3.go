package day3

import (
	"strconv"
	"strings"
	"unicode"

	. "github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	sc := parseGrid(input)
	for rowIdx := 0; rowIdx < sc.Rows(); rowIdx++ {
		for colIdx := 0; colIdx < sc.Cols(); colIdx++ {
			cur := NewCoord(rowIdx, colIdx)
			if unicode.IsNumber(sc.Get(cur)) {
				buffer := make([]rune, 0)
				var isAdjecent bool
				var number int
				for scanIdx := cur.Col; scanIdx < sc.Cols(); scanIdx++ {
					ahead := NewCoord(cur.Row, scanIdx)
					if !unicode.IsNumber(sc.Get(ahead)) {
						colIdx = scanIdx
						break
					}
					curAdjecent, _ := checkIfAdjecent(sc, matchAllComponents, ahead)
					isAdjecent = curAdjecent || isAdjecent
					buffer = append(buffer, sc.Get(ahead))
					number, _ = strconv.Atoi(string(buffer))
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
			if unicode.IsNumber(sc.Get(cur)) {
				buffer := make([]rune, 0)
				var isAdjecent bool
				var currentSymbolLoc Coord
				for scanIdx := cur.Col; scanIdx < sc.Cols(); scanIdx++ {
					ahead := NewCoord(cur.Row, scanIdx)
					if !unicode.IsNumber(sc.Get(ahead)) {
						colIdx = scanIdx
						break
					}
					curAdjecent, curSymbol := checkIfAdjecent(sc, matchGearComponents, ahead)
					isAdjecent = curAdjecent || isAdjecent
					buffer = append(buffer, sc.Get(ahead))
					if curAdjecent {
						currentSymbolLoc = curSymbol
					}
				}
				currentRatio, _ := strconv.Atoi(string(buffer))
				gearMap[currentSymbolLoc] = append(gearMap[currentSymbolLoc], currentRatio)
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
	for _, c := range sc.GetAdjecent(check, 1) {
		if match(sc.Get(c)) {
			return true, c
		}
	}
	return false, Coord{}
}

type ComponentMacher func(cmp rune) bool

func matchAllComponents(cmp rune) bool {
	return !unicode.IsNumber(cmp) && cmp != '.'
}

func matchGearComponents(cmp rune) bool {
	return cmp == '*'
}
