package main

import "testing"

func TestDay2_1(t *testing.T) {
	inputfile := "../test_input.txt"
	expected := 4
	actual := day2_2(inputfile)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
