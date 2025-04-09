package main

import (
	"fmt"
	"log"
	"maps"
	// "sort"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	grid := u.ReadFileLinesMust("../input.txt")
	result := solve(grid)
	fmt.Println("result: ", result)
}

func solve(grid []string) int {
	grid_plant_map := make(map[rune]map[u.Coordinate]bool)
	outside_grid := u.Coordinate{Row: len(grid), Col: len(grid[0])}
	for row := range grid {
		for col := range grid[row] {
			char := rune(grid[row][col])
			if len(grid_plant_map[char]) == 0 {
				grid_plant_map[char] = make(map[u.Coordinate]bool)
			}
			grid_plant_map[char][u.Coordinate{Row: row, Col: col}] = true
		}
	}
	var total int
	for char, coors := range grid_plant_map {
		fmt.Printf("\nplant %s \n", string(char))
		total += get_distinct_area_data(coors, outside_grid)
	}
	return total
}

func get_distinct_area_data(coors map[u.Coordinate]bool, outside_grid u.Coordinate) int {
	total := 0
	for len(coors) > 0 {
		current_location := some_location(coors)
		plant_area := get_neighbors(coors, current_location, outside_grid)
		number_of_sides := get_number_of_corners(plant_area)
		area := len(plant_area)
		price := area * number_of_sides
		println(area, "*", number_of_sides, "=", price)
		total += price
	}
	return total
}

func some_location(coors map[u.Coordinate]bool) u.Coordinate {
	for k := range coors {
		return k
	}
	log.Fatal("Should never reach this point")
	return u.Coordinate{}
}

func get_neighbors(coors map[u.Coordinate]bool, current_location u.Coordinate, outside_grid u.Coordinate) map[u.Coordinate]bool {
	neighbors := make(map[u.Coordinate]bool)
	if current_location.Row < 0 || current_location.Row >= outside_grid.Row || current_location.Col < 0 || current_location.Col >= outside_grid.Col {
		return neighbors
	}
	if !coors[current_location] {
		return neighbors
	}
	neighbors[current_location] = true
	// fmt.Printf("Added coordinate: %v\n", current_location)
	delete(coors, current_location)
	maps.Copy(neighbors, get_neighbors(coors, u.Coordinate{Row: current_location.Row + 1, Col: current_location.Col}, outside_grid))
	maps.Copy(neighbors, get_neighbors(coors, u.Coordinate{Row: current_location.Row - 1, Col: current_location.Col}, outside_grid))
	maps.Copy(neighbors, get_neighbors(coors, u.Coordinate{Row: current_location.Row, Col: current_location.Col + 1}, outside_grid))
	maps.Copy(neighbors, get_neighbors(coors, u.Coordinate{Row: current_location.Row, Col: current_location.Col - 1}, outside_grid))
	return neighbors
}

func get_number_of_corners(plant_areas map[u.Coordinate]bool) int {
	corners := 0
	for area := range plant_areas {
		left_area := plant_areas[u.Coordinate{Row: area.Row, Col: area.Col - 1}]
		right_area := plant_areas[u.Coordinate{Row: area.Row, Col: area.Col + 1}]
		up_area := plant_areas[u.Coordinate{Row: area.Row - 1, Col: area.Col}]
		down_area := plant_areas[u.Coordinate{Row: area.Row + 1, Col: area.Col}]
		dia_up_left := plant_areas[u.Coordinate{Row: area.Row - 1, Col: area.Col - 1}]
		dia_up_right := plant_areas[u.Coordinate{Row: area.Row - 1, Col: area.Col + 1}]
		dia_down_left := plant_areas[u.Coordinate{Row: area.Row + 1, Col: area.Col - 1}]
		dia_down_right := plant_areas[u.Coordinate{Row: area.Row + 1, Col: area.Col + 1}]
		if !left_area && !up_area {
			corners++
		}
		if !right_area && !up_area {
			corners++
		}
		if !left_area && !down_area {
			corners++
		}
		if !right_area && !down_area {
			corners++
		}
		if !dia_up_left && up_area && left_area {
			corners++
		}
		if !dia_up_right && up_area && right_area {
			corners++
		}
		if !dia_down_left && down_area && left_area {
			corners++
		}
		if !dia_down_right && down_area && right_area {
			corners++
		}

	}
	return corners
}

// type Side struct {
// 	horizontal bool
// 	side  int
// }

// func get_number_of_sides(plant_areas map[u.Coordinate]bool) int {
// 	sides := make(map[Side][]int)
// 	for k := range plant_areas {
// 		if !plant_areas[u.Coordinate{Row: k.Row + 1, Col: k.Col}] {
// 			lower_side := Side{horizontal: true, side: k.Row}
// 			sides[lower_side] = append(sides[lower_side], k.Col)
// 		}
// 		if !plant_areas[u.Coordinate{Row: k.Row - 1, Col: k.Col}] {
// 			upper_side := Side{horizontal: true, side: k.Row - 1}
// 			sides[upper_side] = append(sides[upper_side], k.Col)
// 		}
// 		if !plant_areas[u.Coordinate{Row: k.Row, Col: k.Col + 1}] {
// 			right_side := Side{horizontal: false, side: k.Col}
// 			sides[right_side] = append(sides[right_side], k.Row)
// 		}
// 		if !plant_areas[u.Coordinate{Row: k.Row, Col: k.Col - 1}] {
// 			left_side := Side{horizontal: false, side: k.Col - 1}
// 			sides[left_side] = append(sides[left_side], k.Row)
// 		}
// 	}
// 	number_of_sides := 0
// 	for _, positions := range sides {
// 		sort.Ints(positions)
// 		number_of_sides++
// 		last_pos := positions[0]
// 		for i := 1; i < len(positions); i++ {
// 			if positions[i]-last_pos > 1 {
// 				number_of_sides++
// 			}
// 			last_pos = positions[i]
// 		}
// 	}
// 	return number_of_sides
// }
