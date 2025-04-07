package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
	// obs := get_obstacles("./obs.txt")
	// for i := range obs {
	// 	println(obs[i].Col, obs[i].Row)
	// }

}

func solve(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid []string
	var guard_location u.Coordinate
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		for col, char := range line {
			if char == '>' || char == '<' || char == '^' || char == 'v' {
				guard_location = u.Coordinate{Row: row, Col: col, Symbol: char}
			}
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid_task_1 := append([]string(nil), grid...)
	guard_location_task_1 := u.Coordinate{Row: guard_location.Row, Col: guard_location.Col, Symbol: guard_location.Symbol}
	for {
		if guard_location_task_1.Symbol == 'E' {
			break
		}
		guard_location_task_1 = walk_in_grid(grid_task_1, guard_location_task_1)
	}
	u.PrintGrid(grid_task_1)

	test_runs := 0
	c := make(chan int)
	c_loc := make(chan u.Coordinate)
	for i := range grid {
		for j := range len(grid[0]) {
			if grid[i][j] == '.' && grid_task_1[i][j] != '.' {
				test_runs++
				grid_copy := append([]string(nil), grid...)
				grid_copy[i] = u.ReplaceChar(grid_copy[i], '#', j)
				go try_walking_with_new_obstacle(grid_copy, u.Coordinate{Row: guard_location.Row, Col: guard_location.Col, Symbol: guard_location.Symbol}, test_runs, c, c_loc, u.Coordinate{Row: i, Col: j, Symbol: 'O'})
			}
		}
	}

	var total int
	for i := 0; i < test_runs; i++ {
		x := <-c
		total += x
		println(total)
	}
	for i := 0; i < total; i++ {
		obstacle := <-c_loc
		grid[obstacle.Row] = u.ReplaceChar(grid[obstacle.Row], 'O', obstacle.Col)
	}
	u.PrintGrid(grid)
	return total
}

func try_walking_with_new_obstacle(grid []string, guard_location u.Coordinate, idx int, c chan int, c_loc chan u.Coordinate, obstacle u.Coordinate) {
	for {
		if guard_location.Symbol == 'E' {
			c <- 0
			return
		}
		if guard_location.Symbol == '4' {
			c <- 1
			c_loc <- obstacle
			return
		}
		guard_location = walk_in_grid(grid, guard_location)
	}
}

func walk_in_grid(grid []string, before u.Coordinate) u.Coordinate {
	var next_field u.Coordinate
	if before.Symbol == '^' {
		next_field = u.Coordinate{Row: before.Row - 1, Col: before.Col, Symbol: before.Symbol}
	} else if before.Symbol == '>' {
		next_field = u.Coordinate{Row: before.Row, Col: before.Col + 1, Symbol: before.Symbol}
	} else if before.Symbol == 'v' {
		next_field = u.Coordinate{Row: before.Row + 1, Col: before.Col, Symbol: before.Symbol}
	} else if before.Symbol == '<' {
		next_field = u.Coordinate{Row: before.Row, Col: before.Col - 1, Symbol: before.Symbol}
	}
	if !u.InsideGrid(grid, next_field) {
		grid[before.Row] = u.ReplaceChar(grid[before.Row], '1', before.Col)
		return u.Coordinate{Row: before.Row, Col: before.Col, Symbol: 'E'}
	}
	if grid[next_field.Row][next_field.Col] == '#' {
		next_field = u.Coordinate{Row: before.Row, Col: before.Col, Symbol: rotate(before.Symbol)}
	} else {
		increase_guard_step(grid, before)
		if grid[before.Row][before.Col] == '4' {
			return u.Coordinate{Row: before.Row, Col: before.Col, Symbol: '4'} // it's a loop!
		}
	}
	return next_field
}

func increase_guard_step(grid []string, before u.Coordinate) {
	var count rune
	before_count := grid[before.Row][before.Col]
	switch before_count {
	case '1':
		count = '2'
	case '2':
		count = '3'
	case '3':
		count = '4'
	case '.':
		count = '1'
	default:
		println("Guard start direction", string(before_count))
		count = '1'
	}
	grid[before.Row] = u.ReplaceChar(grid[before.Row], count, before.Col)
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
