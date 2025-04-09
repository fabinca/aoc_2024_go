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
	stones_map := make(map[string]int)
	for _, stone := range strings.Split(stones, " ") {
		stones_map[stone]++
	}
	for i := 0; i < blink_number; i++ {
		stones_map = blink(stones_map)
		// println("Blink: ", i, "\n")
		// println("\n")
	}

	total := 0
	for _, value := range stones_map {
		total += value
	}
	return total
}

func blink(stones map[string]int) map[string]int {
	new_stones := make(map[string]int)
	for key, value := range stones {
		if key == "0" {
			new_stones["1"] += value
		} else if len(key)%2 == 0 {
			new_stones[key[:len(key)/2]] += value
			new_stones[remove_leading_zeros(key[len(key)/2:])] += value
		} else {
			new_stones[strconv.Itoa(2024*u.AtoiMust(key))] += value
		}
	}
	return new_stones
}

func remove_leading_zeros(s string) string {

	if len(s) <= 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			return s[i:]
		}
	}
	return "0"
}
