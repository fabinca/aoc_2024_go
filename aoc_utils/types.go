package aoc_utils

type Coordinate struct {
	Row    int
	Col    int
	Symbol rune
}

func (a *Coordinate) Add(b Coordinate) {
	a.Row += b.Row
	a.Col += b.Col
}

func (a *Coordinate) Substract(b Coordinate) {
	a.Row -= b.Row
	a.Col -= b.Col
}

func (a Coordinate) Equals(b Coordinate) bool {
	return a.Row == b.Row && a.Col == b.Col
}

type Pair struct {
	A string
	B string
}
