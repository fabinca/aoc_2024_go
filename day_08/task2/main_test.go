package main

import (
	u "github.com/fabinca/aoc_2024_go/aoc_utils"
	"testing"
)

func TestDay(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 34
	actual := solve(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestAntinodesHori(t *testing.T) {
	actual := find_antinodes(u.Coordinate{Col: 0, Row: 0}, u.Coordinate{Col: 1, Row: 1}, 3, 3)
	if len(actual) != 3 {
		t.Errorf("Missing antinode: only have %+v", actual[0])
	}
	for i := range actual {
		if actual[i].Col == actual[i].Row {
			continue
		}
		t.Errorf("Got unexpected antinode: %+v", actual[i])
	}
}
