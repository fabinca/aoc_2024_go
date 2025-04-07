package main

import "testing"

func TestDay2_1(t *testing.T) {
	inputfile := "./test_input.txt"
	expected := 6
	actual := solve(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
