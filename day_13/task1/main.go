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
	loc := u.Coordinate{Row: 0, Col: 0}
	button_a := 0
	button_b := 0
	for loc.Row < game.prize.Row && loc.Col < game.prize.Col && button_b < 100 {
		button_b++
		loc.Add(game.button_b)
	}
	for button_a < 100 && loc.Row >= 0 && loc.Col >= 0 {
		if game.prize.Equals(loc) {
			return button_a*3 + button_b
		}
		loc.Substract(game.button_b)
		if button_b >= 0 {
			button_b--
		}
		if game.prize.Equals(loc) {
			return button_a*3 + button_b
		}
		for loc.Row < game.prize.Row && loc.Col < game.prize.Col && button_a < 100 {
			loc.Add(game.button_a)
			button_a++
		}
	}
	return 0
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
	current_game.prize = u.Coordinate{Row: u.AtoiRemoveNonDigit(lines[2][prize_x_idx:prize_y_idx]), Col: u.AtoiRemoveNonDigit(lines[2][prize_y_idx:])}
	return current_game
}
