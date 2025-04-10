package main

import (
	"fmt"
	"strings"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	result := solve("../input.txt", 103, 101)
	fmt.Println("result: ", result)
}

func solve(inputfile string, maxRow int, maxCol int) int {
	lines := u.ReadFileLinesMust(inputfile)
	channel := make(chan u.Coordinate, len(lines))
	for _, line := range lines {
		go robo_move(line, channel, maxRow, maxCol)
	}
	quadrant := make(map[rune]int)
	middleRow := (maxRow - 1) / 2
	middleCol := (maxCol - 1) / 2
	for i := 0; i < len(lines); i++ {
		pos := <-channel
		if pos.Row == middleRow || pos.Col == middleCol {
			continue
		}
		if pos.Row < maxRow/2 {
			if pos.Col < maxCol/2 {
				quadrant['0'] += 1
			} else {
				quadrant['1'] += 1
			}

		} else {
			if pos.Col < maxCol/2 {
				quadrant['2'] += 1
			} else {
				quadrant['3'] += 1
			}
		}
	}
	total := quadrant['0'] * quadrant['1'] * quadrant['2'] * quadrant['3']
	return total
}

func robo_move(line string, c chan u.Coordinate, maxRow int, maxCol int) {
	pos, velo := parse_robo(line)
	for secs := 0; secs < 100; secs++ {
		pos.Add(velo)
		if pos.Col < 0 {
			pos.Col += maxCol
		}
		if pos.Row < 0 {
			pos.Row += maxRow
		}
		if pos.Col >= maxCol {
			pos.Col -= maxCol
		}
		if pos.Row >= maxRow {
			pos.Row -= maxRow
		}
	}
	c <- pos
}

func parse_robo(line string) (u.Coordinate, u.Coordinate) {
	pos := u.Coordinate{}
	velo := u.Coordinate{}
	parts := strings.Split(line, " ")

	pPart := strings.TrimPrefix(parts[0], "p=")
	pValues := strings.Split(pPart, ",")

	vPart := strings.TrimPrefix(parts[1], "v=")
	vValues := strings.Split(vPart, ",")

	pos.Col = u.AtoiMust(pValues[0])
	pos.Row = u.AtoiMust(pValues[1])
	velo.Col = u.AtoiMust(vValues[0])
	velo.Row = u.AtoiMust(vValues[1])

	return pos, velo

}
