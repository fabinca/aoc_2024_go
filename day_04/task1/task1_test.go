package main

import "testing"

func TestDay4_1(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 18
	actual := day4_1(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
