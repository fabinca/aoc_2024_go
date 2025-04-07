package aoc_utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func StringToIntSliceMust(s string) []int {
	fields := strings.Fields(s)
	var intslice []int
	for i := range fields {
		intslice = append(intslice, AtoiRemoveNonDigit(fields[i]))
	}
	return intslice
}

func ReadFileLinesMust(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
