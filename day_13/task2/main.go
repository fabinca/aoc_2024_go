package main

import (
	"fmt"
	"strings"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
}

type Game struct {
	button_a u.Coordinate
	button_b u.Coordinate
	prize    u.Coordinate
}

func solve(inputfile string) int {
	lines := u.ReadFileLinesMust(inputfile)
	total := 0
	for i := 0; i < len(lines); i += 4 {
		current_game := parseGame(lines[i : i+3])
		total += playGame(current_game)

	}
	return total
}

func playGame(game Game) int {
	// prize_x = a * buttona_x + b * buttonb_x
	// prize_y = a * buttona_y + b * buttonb_y
	// a = (prize_x - b * button_b_x) / button_a_x
	// a = (prize_y - b * button_b_y) / button_a_y
	// (prize_x - b * button_b_x) / button_a_x == (prize_y - b * button_b_y) / button_a_y
	// button_a_y * prize_x - b * button_bx * button_a_y == prize_y * button_a_x - b * button_by * button_a_x
	// (button_a_y * prize_x - prize_y * button_a_x) / ((button_bx * button_a_y -button_by * button_a_x)) ==  b
	b := (game.button_a.Col*game.prize.Row - game.prize.Col*game.button_a.Row) / (game.button_b.Row*game.button_a.Col - game.button_b.Col*game.button_a.Row)
	if b < 0 {
		return 0
	}
	a := (game.prize.Col - b*game.button_b.Col) / game.button_a.Col
	if a < 0 {
		return 0
	}
	if a*game.button_a.Col+b*game.button_b.Col != game.prize.Col {
		return 0
	}
	if b*game.button_b.Row+a*game.button_a.Row != game.prize.Row {
		return 0
	}
	return a*3 + b
}

func parseGame(lines []string) Game {
	current_game := Game{}
	button_a_x_idx := strings.Index(lines[0], "X+") + 2
	button_a_y_idx := strings.Index(lines[0], "Y+") + 2
	current_game.button_a = u.Coordinate{Row: u.AtoiRemoveNonDigit(lines[0][button_a_x_idx:button_a_y_idx]), Col: u.AtoiRemoveNonDigit(lines[0][button_a_y_idx:])}
	button_b_x_idx := strings.Index(lines[1], "X+") + 2
	button_b_y_idx := strings.Index(lines[1], "Y+") + 2
	current_game.button_b = u.Coordinate{Row: u.AtoiRemoveNonDigit(lines[1][button_b_x_idx:button_b_y_idx]), Col: u.AtoiRemoveNonDigit(lines[1][button_b_y_idx:])}
	prize_x_idx := strings.Index(lines[2], "X=") + 2
	prize_y_idx := strings.Index(lines[2], "Y=") + 2
	current_game.prize = u.Coordinate{Row: 10000000000000 + u.AtoiRemoveNonDigit(lines[2][prize_x_idx:prize_y_idx]), Col: 10000000000000 + u.AtoiRemoveNonDigit(lines[2][prize_y_idx:])}
	return current_game
}
