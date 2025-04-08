package main

import (
	u "github.com/fabinca/aoc_2024_go/aoc_utils"
	"testing"
)

func TestDay(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 14
	actual := solve(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestAntinodes(t *testing.T) {
	actual := find_antinodes(u.Coordinate{Col: 4, Row: 3}, u.Coordinate{Col: 5, Row: 5})
	if len(actual) != 2 {
		t.Errorf("Missing antinode: only have %+v", actual[0])
	}
	for i := range actual {
		if actual[i].Col == 6 && actual[i].Row == 7 {
			continue
		}
		if actual[i].Col == 3 && actual[i].Row == 1 {
			continue
		}
		t.Errorf("Got unexpected antinode: %+v", actual[i])
	}
}
