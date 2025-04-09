package main

import (
	u "github.com/fabinca/aoc_2024_go/aoc_utils"
	"testing"
)

func TestDay(t *testing.T) {
	input := u.ReadFileLinesMust("../test_input.txt")
	expected := 1206
	actual := solve(input)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay_Edge(t *testing.T) {
	grid := []string{"OOOOO", "OXOXO", "OOOOO", "OXOXO", "OOOOO"}
	expected := 436
	actual := solve(grid)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay_E(t *testing.T) {
	grid := []string{"EEEEE", "EXXXX", "EEEEE", "EXXXX", "EEEEE"}
	expected := 236
	actual := solve(grid)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay_Small(t *testing.T) {
	grid := []string{"AAAA", "BBCD", "BBCC", "EEEC"}
	expected := 80
	actual := solve(grid)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
