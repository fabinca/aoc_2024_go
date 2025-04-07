package aoc_utils

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

func replaceChar(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func printgrid(grid []string) {
	for row := range grid {
		println(grid[row])
	}
	println(" ")
}
