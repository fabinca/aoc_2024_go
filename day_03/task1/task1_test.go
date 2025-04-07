package main

import "testing"

func TestDay2_1(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 161
	actual := day3_1(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
