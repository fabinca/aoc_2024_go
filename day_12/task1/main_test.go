package main

import (
	"testing"

	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func TestDay(t *testing.T) {
	input := u.ReadFileLinesMust("../test_input.txt")
	expected := 1930
	actual := solve(input)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay_(t *testing.T) {
	grid := []string{"OOOOO", "OXOXO", "OOOOO", "OXOXO", "OOOOO"}
	expected := 772
	actual := solve(grid)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
