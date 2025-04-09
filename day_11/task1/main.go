package main

import (
	"fmt"
	"strconv"
	"strings"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	result := solve("872027 227 18 9760 0 4 67716 9245696", 75)
	fmt.Println("result: ", result)
}

func solve(stones string, blink_number int) int {
	println(stones)
	for i := 0; i < blink_number; i++ {
		stones = blink(stones)
		println("Blink: ", i, "\n", len(stones))
		if len(stones) < 100 {
			println(stones)
		}
		println("\n")
	}

	total := len(strings.Split(stones, " "))
	return total
}

func blink(stones string) string {
	single_stones := strings.Split(stones, " ")
	var new_stones string
	for i := range single_stones {
		stone := strings.TrimSpace(single_stones[i])
		if stone == " " {
			continue
		}
		if stone == "0" {
			new_stones += "1"
		} else if len(stone)%2 == 0 {
			new_stones += stone[:len(stone)/2]
			new_stones += " "
			new_stones += remove_leading_zeros(stone[len(stone)/2:])
		} else {
			new_stones += strconv.Itoa(2024 * u.AtoiMust(stone))
		}
		new_stones += " "
	}
	return new_stones[:len(new_stones)-1]
}

func remove_leading_zeros(s string) string {

	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			return s[i:]
		}
	}
	return "0"
}
