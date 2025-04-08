package main

import (
	"fmt"
	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
}

func solve(inputfile string) int {
	grid := u.ReadFileLinesMust(inputfile)
	coors := parse_coordinates(grid)
	total := 0
	for key, value := range coors {
		if value == '0' {
			total += find_trail_head_score(coors, key)
		}
	}
	return total
}

func find_trail_head_score(coors map[u.Coordinate]byte, start_coord u.Coordinate) int {
	distinct_trails := get_next_positions(coors, start_coord, byte('0'))
	return distinct_trails
}

func get_next_positions(coors map[u.Coordinate]byte, pos u.Coordinate, height byte) int {
	trail_num := 0
	if coors[pos] != height {
		return 0
	}
	if coors[pos] == '9' {
		return 1
	}
	trail_num += get_next_positions(coors, u.Coordinate{Row: pos.Row - 1, Col: pos.Col}, height+1)
	trail_num += get_next_positions(coors, u.Coordinate{Row: pos.Row + 1, Col: pos.Col}, height+1)
	trail_num += get_next_positions(coors, u.Coordinate{Row: pos.Row, Col: pos.Col + 1}, height+1)
	trail_num += get_next_positions(coors, u.Coordinate{Row: pos.Row, Col: pos.Col - 1}, height+1)
	return trail_num
}

func parse_coordinates(grid []string) map[u.Coordinate]byte {
	coors := make(map[u.Coordinate]byte)
	for row := range grid {
		for col := range grid[row] {
			symbol := grid[row][col]
			coors[u.Coordinate{Row: row, Col: col}] = symbol
		}
	}
	return coors
}
