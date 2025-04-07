package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := day3_1("../input.txt")
	fmt.Println("result: ", result)
}

func day3_1(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	for scanner.Scan() {
		line := scanner.Text()
		muls := re.FindAllString(line, -1)
		for i := 0; i < len(muls); i++ {
			println(muls[i])
			nums := strings.Split(muls[i][4:], ",")
			total += AtoiRemoveNonDigit(nums[0]) * AtoiRemoveNonDigit(nums[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total
}

func AtoiRemoveNonDigit(s string) int {
	re := regexp.MustCompile(`\D`) // Match any non-digit character (i.e., [^0-9])
	new := re.ReplaceAllString(s, "")
	i, err := strconv.Atoi(new)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
