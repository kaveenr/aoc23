package commons

type Coord struct {
	Row, Col int
}

func NewCoord(row, col int) Coord {
	return Coord{Row: row, Col: col}
}
