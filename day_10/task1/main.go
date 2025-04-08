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
	next_positions := make(map[u.Coordinate]bool)
	next_positions[start_coord] = true
	for height := '0'; height < '9'; height++ {
		if len(next_positions) == 0 {
			return 0
		}
		next_positions = get_next_positions(coors, next_positions, byte(height))
	}
	return len(next_positions)
}

func get_next_positions(coors map[u.Coordinate]byte, before_pos map[u.Coordinate]bool, height byte) map[u.Coordinate]bool {
	next_positions := make(map[u.Coordinate]bool)
	for pos := range before_pos {
		if coors[u.Coordinate{Row: pos.Row - 1, Col: pos.Col}] == height+1 {
			next_positions[u.Coordinate{Row: pos.Row - 1, Col: pos.Col}] = true
		}
		if coors[u.Coordinate{Row: pos.Row + 1, Col: pos.Col}] == height+1 {
			next_positions[u.Coordinate{Row: pos.Row + 1, Col: pos.Col}] = true
		}
		if coors[u.Coordinate{Row: pos.Row, Col: pos.Col - 1}] == height+1 {
			next_positions[u.Coordinate{Row: pos.Row, Col: pos.Col - 1}] = true
		}
		if coors[u.Coordinate{Row: pos.Row, Col: pos.Col + 1}] == height+1 {
			next_positions[u.Coordinate{Row: pos.Row, Col: pos.Col + 1}] = true
		}
	}
	return next_positions
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
