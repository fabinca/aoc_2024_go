package main

// The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9

// This example data contains six reports each containing five levels.

// The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

//     The levels are either all increasing or all decreasing.
//     Any two adjacent levels differ by at least one and at most three.

// In the example above, the reports can be found safe or unsafe by checking those rules:

//     7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
//     1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
//     9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
//     1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
//     8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
//     1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

// So, in this example, 2 reports are safe.

// Analyze the unusual data from the engineers. How many reports are safe?

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := day2_1("../input.txt")
	fmt.Println("result: ", result)
}

func day2_1(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		last_num := Atoi(fields[0])
		direction := 0 // 0 = undetermined, 1 = increasing, -1 = decreasing
		safe := true

		for i := 1; i < len(fields); i++ {
			num := Atoi(fields[i])
			diff := num - last_num

			// Check if no change (must be increasing or decreasing)
			if diff == 0 {
				safe = false
				break
			}

			// Set direction on first comparison
			if direction == 0 {
				if diff > 0 {
					direction = 1 // increasing
				} else {
					direction = -1 // decreasing
				}
			}

			// Check if direction changed or difference is too large/small
			if (direction*diff <= 0) || Abs(diff) > 3 {
				safe = false
				break
			}

			last_num = num
		}

		if safe {
			total += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
