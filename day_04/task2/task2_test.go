package main

import "testing"

func TestDay4_2(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 9
	actual := day4_2(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
