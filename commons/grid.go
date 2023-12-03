package commons

import (
	"fmt"
	"strings"
)

type Grid [][]rune

func (s Grid) Get(c Coord) rune {
	return s[c.Row][c.Col]
}

func (s Grid) Rows() int {
	return len(s)
}

func (s Grid) Cols() int {
	if len(s) > 0 {
		return len(s[0])
	}
	return 0
}

func (s Grid) GetAdjecent(c Coord, size int) (res []Coord) {
	res = make([]Coord, 0)
	for checkRowIdx := c.Row - size; checkRowIdx <= c.Row+size; checkRowIdx++ {
		for checkColIdx := c.Col - size; checkColIdx <= c.Col+size; checkColIdx++ {
			if (checkRowIdx >= 0 && checkRowIdx < s.Rows()) && (checkColIdx >= 0 && checkColIdx < s.Cols()) {
				res = append(res, NewCoord(checkRowIdx, checkColIdx))
			}
		}
	}
	return res
}

func (s Grid) viz() {
	for rowIdx := range s {
		for colIdx := range s[rowIdx] {
			fmt.Print(string(s[rowIdx][colIdx]))
		}
		fmt.Println()
	}
}

func NewGrid(input string) (res Grid) {
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
