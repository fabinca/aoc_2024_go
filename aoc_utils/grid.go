package aoc_utils

func InsideGrid(grid []string, loc Coordinate) bool {
	if loc.Row < 0 {
		return false
	}
	if loc.Row >= len(grid) {
		return false
	}
	if loc.Col < 0 {
		return false
	}
	if loc.Col >= len(grid[0]) {
		return false
	}
	return true
}

func ReplaceChar(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func PrintGrid(grid []string) {
	for row := range grid {
		println(grid[row])
	}
	println(" ")
}
