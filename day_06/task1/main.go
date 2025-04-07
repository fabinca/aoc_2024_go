package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
}

func solve(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int
	var grid []string
	var guard_location Coordinate
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		for col, char := range line {
			if char == '>' || char == '<' || char == '^' || char == 'v' {
				guard_location = Coordinate{row, col, char}
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		if guard_location.direction == 'O' {
			break
		}
		guard_location = walk_in_grid(grid, guard_location)
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'X' {
				total++
			}
		}
	}
	printgrid(grid)

	return total
}

type Coordinate struct {
	row       int
	col       int
	direction rune
}

func inside_grid(grid []string, loc Coordinate) bool {
	if loc.row < 0 {
		return false
	}
	if loc.row >= len(grid) {
		return false
	}
	if loc.col < 0 {
		return false
	}
	if loc.col >= len(grid[0]) {
		return false
	}
	return true
}

func walk_in_grid(grid []string, before Coordinate) Coordinate {
	var next_field Coordinate
	if before.direction == '^' {
		next_field = Coordinate{before.row - 1, before.col, before.direction}
	} else if before.direction == '>' {
		next_field = Coordinate{before.row, before.col + 1, before.direction}
	} else if before.direction == 'v' {
		next_field = Coordinate{before.row + 1, before.col, before.direction}
	} else if before.direction == '<' {
		next_field = Coordinate{before.row, before.col - 1, before.direction}
	}
	if !inside_grid(grid, next_field) {
		grid[before.row] = replaceChar(grid[before.row], 'X', before.col)
		return Coordinate{before.row, before.col, 'O'}
	}
	if grid[next_field.row][next_field.col] == '#' {
		next_field = Coordinate{before.row, before.col, rotate(before.direction)}
	} else {
		grid[before.row] = replaceChar(grid[before.row], 'X', before.col)
	}
	return next_field
}

func rotate(direction rune) rune {
	if direction == '^' {
		return '>'
	} else if direction == '>' {
		return 'v'
	} else if direction == 'v' {
		return '<'
	} else if direction == '<' {
		return '^'
	}
	log.Fatal("Invalid direction") // should never reach here, but to satisfy the compiler.
	return 'X'                     // should never reach here, but to satisfy the compiler.
}

func replaceChar(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func printgrid(grid []string) {
	for row := range grid {
		println(grid[row])
	}
	println(" ")
}
