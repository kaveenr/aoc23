package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Part1(input string) (result int) {
	sc := parseSchematic(input)
	for rowIdx := 0; rowIdx < sc.Rows(); rowIdx++ {
		for colIdx := 0; colIdx < sc.Cols(); colIdx++ {
			cur := NewCoord(rowIdx, colIdx)
			if unicode.IsNumber(sc.Get(cur)) {
				ref := make([]rune, 0)
				isAdjecent := false
				var number int
				for scanIdx := cur.col; scanIdx < sc.Cols(); scanIdx++ {
					ahead := NewCoord(cur.row, scanIdx)
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
	sc := parseSchematic(input)
	gearMap := make(map[Coord][]int)
	for rowIdx := 0; rowIdx < sc.Rows(); rowIdx++ {
		for colIdx := 0; colIdx < sc.Cols(); colIdx++ {
			cur := NewCoord(rowIdx, colIdx)
			var currentSymbol Coord
			if unicode.IsNumber(sc.Get(cur)) {
				ref := make([]rune, 0)
				isAdjecent := false
				var currentRatio int
				for scanIdx := cur.col; scanIdx < sc.Cols(); scanIdx++ {
					ahead := NewCoord(cur.row, scanIdx)
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

func parseSchematic(input string) (res Schematic) {
	res = make(Schematic, 0)
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

func checkIfAdjecent(sc Schematic, match ComponentMacher, check Coord) (bool, Coord) {
	for checkRowIdx := check.row - 1; checkRowIdx <= check.row+1; checkRowIdx++ {
		for checkColIdx := check.col - 1; checkColIdx <= check.col+1; checkColIdx++ {
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

type Coord struct {
	row, col int
}

func NewCoord(row, col int) Coord {
	return Coord{row: row, col: col}
}

type Schematic [][]rune

func (s Schematic) Get(c Coord) rune {
	return s[c.row][c.col]
}

func (s Schematic) Rows() int {
	return len(s)
}

func (s Schematic) Cols() int {
	if len(s) > 0 {
		return len(s[0])
	}
	return 0
}

func (s Schematic) viz() {
	for rowIdx := range s {
		for colIdx := range s[rowIdx] {
			fmt.Print(string(s[rowIdx][colIdx]))
		}
		fmt.Println()
	}
}
