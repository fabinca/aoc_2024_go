package main

import (
	"fmt"
	"log"
	"maps"

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
		perimiter := get_perimeter(plant_area)
		area := len(plant_area)
		price := area * perimiter
		// println(area, "*", perimiter, "=", price)
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

func get_perimeter(plant_areas map[u.Coordinate]bool) int {
	perimeter := 0
	for k := range plant_areas {
		if !plant_areas[u.Coordinate{Row: k.Row + 1, Col: k.Col}] {
			perimeter++
		}
		if !plant_areas[u.Coordinate{Row: k.Row - 1, Col: k.Col}] {
			perimeter++
		}
		if !plant_areas[u.Coordinate{Row: k.Row, Col: k.Col + 1}] {
			perimeter++
		}
		if !plant_areas[u.Coordinate{Row: k.Row, Col: k.Col - 1}] {
			perimeter++
		}
	}
	return perimeter
}
