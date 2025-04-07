package main

import "testing"

func TestDay3_2(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 48
	actual := day3_2(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
