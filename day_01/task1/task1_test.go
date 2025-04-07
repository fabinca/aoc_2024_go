package main

import "testing"

func TestDay1(t *testing.T) {
	inputfile := "../test_input.txt"

	// 1 3   2
	// 2 3      1
	// 3 3
	// 3 4      1
	// 3 5      2
	// 4 9      5
	expected := 11
	actual := day1(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
