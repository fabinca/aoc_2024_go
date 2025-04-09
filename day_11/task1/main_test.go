package main

import "testing"

// func TestDay(t *testing.T) {
// 	input := "125 17"
// 	expected := 22
// 	blink_number := 6
// 	actual := solve(input, blink_number)
// 	if actual != expected {
// 		t.Errorf("Expected %d, got %d", expected, actual)
// 	}
// }

func TestDayZero(t *testing.T) {
	input := "0"
	expected := 22
	blink_number := 75
	actual := solve(input, blink_number)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

// func TestDay_(t *testing.T) {
// 	input := "125 17"
// 	expected := 55312
// 	blink_number := 25
// 	actual := solve(input, blink_number)
// 	if actual != expected {
// 		t.Errorf("Expected %d, got %d", expected, actual)
// 	}
// }

// func TestRemoveZero(t *testing.T) {
// 	actual := remove_leading_zeros("000123")
// 	if actual != "123" {
// 		t.Errorf("Expected %s, got %s", "123", actual)
// 	}

// 	actual = remove_leading_zeros("000")
// 	if actual != "0" {
// 		t.Errorf("Expected %s, got %s", "0", actual)
// 	}
// 	actual = remove_leading_zeros("123")
// 	if actual != "123" {
// 		t.Errorf("Expected %s, got %s", "123", actual)
// 	}

// }
