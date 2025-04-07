package main

import "testing"

func TestDay(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 11387
	actual := solve(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
