package main

import "testing"

func TestDay(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 12
	actual := solve(inputfile, 7, 11)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
