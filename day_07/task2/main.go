package main

import (
	"fmt"
	u "github.com/fabinca/aoc_2024_go/aoc_utils"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
}

func solve(inputfile string) int {
	lines := u.ReadFileLinesMust(inputfile)
	total := 0
	for i := range lines {
		nums := u.StringToIntSliceMust(lines[i])
		test_value := nums[0]
		if calculate(test_value, nums, 1, 0) {
			println(test_value)
			total += test_value
		}
	}
	return total
}

func calculate(test_value int, nums []int, next_id int, current_value int) bool {
	if current_value == test_value && next_id == len(nums) {
		return true
	}
	if current_value > test_value || next_id == len(nums) {
		return false
	}
	if calculate(test_value, nums, next_id+1, current_value+nums[next_id]) {
		return true
	}
	if calculate(test_value, nums, next_id+1, current_value*nums[next_id]) {
		return true
	}
	return calculate(test_value, nums, next_id+1, concatenation_operator(current_value, nums[next_id]))
}

func concatenation_operator(left int, right int) int {
	result := 0 + right
	for right > 0 {
		left *= 10
		right /= 10
	}
	result += left
	return result
}
