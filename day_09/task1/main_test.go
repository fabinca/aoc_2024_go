package main

import "testing"

func TestDay(t *testing.T) {
	input := "2333133121414131402"
	expected := 1928
	actual := solve(input)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
