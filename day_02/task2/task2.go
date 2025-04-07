package main

// --- Part Two ---

// The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

// The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

// Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

// More of the above example's reports are now safe:

//     7 6 4 2 1: Safe without removing any level.
//     1 2 7 8 9: Unsafe regardless of which level is removed.
//     9 7 6 2 1: Unsafe regardless of which level is removed.
//     1 3 2 4 5: Safe by removing the second level, 3.
//     8 6 4 4 1: Safe by removing the third level, 4.
//     1 3 6 7 9: Safe without removing any level.

// Thanks to the Problem Dampener, 4 reports are actually safe!

// Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := day2_2("../input.txt")
	fmt.Println("result: ", result)
}

func day2_2(inputfile string) int {
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
		safe := false
		for skip_num := 1; skip_num < len(fields); skip_num++ {
			safe = check_line(fields, 1, Atoi(fields[0]), skip_num)
			if safe {
				break
			}
		}
		if !safe {
			safe = check_line(fields, 2, Atoi(fields[1]), 0)
		}
		println(line, "safe: ", safe)

		if safe {
			total += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func check_line(fields []string, start int, last_num int, skip_level int) bool {
	direction := 0
	for i := start; i < len(fields); i++ {
		if i == skip_level {
			continue
		}
		num := Atoi(fields[i])
		diff := num - last_num

		if direction == 0 {
			if diff > 0 {
				direction = 1 // increasing
			} else {
				direction = -1 // decreasing
			}
		}

		if diff == 0 {
			return false
		}
		if (direction*diff <= 0) || Abs(diff) > 3 {
			return false
		}
		last_num = num
	}
	return true
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
