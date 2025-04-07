package main

import (
	"fmt"
	u "github.com/fabinca/aoc_2024_go/aoc_utils"
	"maps"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
}

func solve(inputfile string) int {
	grid := u.ReadFileLinesMust(inputfile)
	nodes := getNodes(grid)
	// printNodes(nodes)
	antinodes := make(map[u.Coordinate]bool)
	for key, value := range nodes {
		new := get_antinodes(value, grid)
		fmt.Printf("Antinodes for %c:\n %+v\n\n", key, new)
		maps.Copy(antinodes, new)
	}
	for key, _ := range antinodes {
		grid[key.Row] = u.ReplaceChar(grid[key.Row], '#', key.Col)
	}
	u.PrintGrid(grid)
	total := len(antinodes)
	return total
}

func getNodes(grid []string) map[rune][]u.Coordinate {
	nodes := make(map[rune][]u.Coordinate)
	for row := range grid {
		for col := range grid[row] {
			symbol := rune(grid[row][col])
			if symbol != '.' {
				node := u.Coordinate{Col: col, Row: row, Symbol: symbol}
				nodes[symbol] = append(nodes[symbol], node)
			}
		}
	}
	return nodes
}

func get_antinodes(node_coors []u.Coordinate, grid []string) map[u.Coordinate]bool {
	antinodes := make(map[u.Coordinate]bool)
	for i := 0; i < len(node_coors)-1; i++ {
		for j := i + 1; j < len(node_coors); j++ {
			new_antinodes := find_antinodes(node_coors[i], node_coors[j])
			for idx := range new_antinodes {
				if u.InsideGrid(grid, new_antinodes[idx]) {
					antinodes[new_antinodes[idx]] = true
				}
			}
		}
	}
	return antinodes
}

func find_antinodes(a u.Coordinate, b u.Coordinate) []u.Coordinate {
	var two_antinodes []u.Coordinate
	dx := b.Col - a.Col
	dy := b.Row - a.Row
	two_antinodes = append(two_antinodes, u.Coordinate{Col: b.Col + dx, Row: b.Row + dy})
	two_antinodes = append(two_antinodes, u.Coordinate{Col: a.Col - dx, Row: a.Row - dy})
	return two_antinodes
}

func printNodes(nodes map[rune][]u.Coordinate) {
	for symbol, coords := range nodes {
		fmt.Printf("%c: %+v\n", symbol, coords)
	}
}
