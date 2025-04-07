package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
	// obs := get_obstacles("./obs.txt")
	// for i := range obs {
	// 	println(obs[i].col, obs[i].row)
	// }

}

func get_obstacles(obsfile string) []Coordinate {
	file, err := os.Open(obsfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	var obs []Coordinate
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])
		obs = append(obs, Coordinate{row: b, col: a, direction: 'O'})
		row++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for i := range obs {
		println(obs[i].col, obs[i].row)
	}
	return obs

}

func solve(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

	grid_task_1 := append([]string(nil), grid...)
	guard_location_task_1 := Coordinate{guard_location.row, guard_location.col, guard_location.direction}
	for {
		if guard_location_task_1.direction == 'E' {
			break
		}
		guard_location_task_1 = walk_in_grid(grid_task_1, guard_location_task_1)
	}
	printgrid(grid_task_1)

	test_runs := 0
	c := make(chan int)
	c_loc := make(chan Coordinate)
	for i := range grid {
		for j := range len(grid[0]) {
			if grid[i][j] == '.' && grid_task_1[i][j] != '.' {
				test_runs++
				grid_copy := append([]string(nil), grid...)
				grid_copy[i] = replaceChar(grid_copy[i], '#', j)
				go try_walking_with_new_obstacle(grid_copy, Coordinate{guard_location.row, guard_location.col, guard_location.direction}, test_runs, c, c_loc, Coordinate{i, j, 'O'})
			}
		}
	}

	var total int
	for i := 0; i < test_runs; i++ {
		x := <-c
		total += x
	}
	var obs []Coordinate
	obs = get_obstacles("./obs.txt")
	for i := 0; i < total; i++ {
		obstacle := <-c_loc
		obs_correct := false
		for i := range obs {
			if obs[i].col == obstacle.col && obs[i].row == obstacle.row {
				obs_correct = true
				break
			}
		}
		if !obs_correct {
			grid[obstacle.row] = replaceChar(grid[obstacle.row], 'O', obstacle.col)
		}
	}
	printgrid(grid)
	return total
}

func try_walking_with_new_obstacle(grid []string, guard_location Coordinate, idx int, c chan int, c_loc chan Coordinate, obstacle Coordinate) {
	for {
		if guard_location.direction == 'E' {
			c <- 0
			return
		}
		if guard_location.direction == '4' {
			c <- 1
			c_loc <- obstacle
			return
		}
		// printgrid(grid)
		guard_location = walk_in_grid(grid, guard_location)
	}
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
		grid[before.row] = replaceChar(grid[before.row], '1', before.col)
		return Coordinate{before.row, before.col, 'E'}
	}
	if grid[next_field.row][next_field.col] == '#' {
		next_field = Coordinate{before.row, before.col, rotate(before.direction)}
	} else {
		grid[before.row] = replaceChar(grid[before.row], '1', before.col)
		if grid[before.row][before.col] == '4' {
			return Coordinate{before.row, before.col, '4'} // it's a loop!
		}
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

	if str[index] == '1' {
		replacement = '2'
	} else if str[index] == '2' {
		replacement = '3'
	} else if str[index] == '3' {
		replacement = '4'
	}
	new_str := str[:index] + string(replacement) + str[index+1:]
	return new_str
}

func printgrid(grid []string) {
	for row := range grid {
		println(grid[row])
	}
	println(" ")
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(es ...T) Set[T] {
	for _, e := range es {
		s[e] = struct{}{}
	}
	return s
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Remove(e T) Set[T] {
	delete(s, e)
	return s
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	for v := range other {
		s.Add(v)
	}

	return s
}
