package commons

import "fmt"

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

func (s Grid) viz() {
	for rowIdx := range s {
		for colIdx := range s[rowIdx] {
			fmt.Print(string(s[rowIdx][colIdx]))
		}
		fmt.Println()
	}
}
