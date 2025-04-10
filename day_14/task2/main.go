package main

import (
	"strings"
	"time"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {

	solve("../input.txt", 101, 103)
}

type robo struct {
	pos  u.Coordinate
	velo u.Coordinate
}

func solve(inputfile string, maxRow int, maxCol int) {
	lines := u.ReadFileLinesMust(inputfile)
	var robo_positions []robo
	for _, line := range lines {
		robo_positions = append(robo_positions, parse_robo(line))
	}
	robo_map := make(map[u.Coordinate]bool)

	for secs := 0; secs < 101*103*2; secs++ {
		if secs%1000 == 0 {
			println(secs)
		}
		clear(robo_map)
		for i, robo := range robo_positions {
			new_pos := robo_move(robo, maxRow, maxCol)
			robo_positions[i] = new_pos
			robo_map[new_pos.pos] = true
		}
		print_this_if_symmetric(maxRow, maxCol, robo_map, secs)
	}
	// print_this_if_symmetric(maxRow, maxCol, robo_map, 100)

}

func print_this_if_symmetric(maxRow int, maxCol int, robo map[u.Coordinate]bool, secs int) {
	// middle_col := (maxCol - 1) / 2
	// for key := range robo {
	// 	mirror := middle_col
	// 	if key.Col < middle_col {
	// 		mirror = middle_col + (middle_col - key.Col)
	// 	} else if key.Col > middle_col {
	// 		mirror = middle_col - (key.Col - middle_col)
	// 	} else {
	// 		continue
	// 	}
	// 	if !robo[u.Coordinate{Row: key.Row, Col: mirror}] {
	// 		return
	// 	}
	// }
	grid := make([]string, maxRow)
	for i := range grid {
		grid[i] = strings.Repeat(".", maxCol)
	}
	for key := range robo {
		grid[key.Row] = u.ReplaceChar(grid[key.Row], '1', key.Col)
	}
	print_grid := 0
	for i := 0; i < maxRow; i++ {
		if strings.Contains(grid[i], strings.Repeat("1", 3)) {
			print_grid++
		}
	}
	if print_grid <= 5 {
		return
	}
	println(secs, "\n")
	u.PrintGrid(grid)
	time.Sleep(1000 * time.Millisecond)
}

func robo_move(robo robo, maxRow int, maxCol int) robo {
	robo.pos.Add(robo.velo)
	if robo.pos.Col < 0 {
		robo.pos.Col += maxCol
	}
	if robo.pos.Row < 0 {
		robo.pos.Row += maxRow
	}
	if robo.pos.Col >= maxCol {
		robo.pos.Col -= maxCol
	}
	if robo.pos.Row >= maxRow {
		robo.pos.Row -= maxRow
	}
	return robo
}

func parse_robo(line string) robo {
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

	return robo{pos: pos, velo: velo}

}
