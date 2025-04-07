package aoc_utils

import (
	"log"
	"regexp"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AtoiMust(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func AtoiRemoveNonDigit(s string) int {
	re := regexp.MustCompile(`\D`)
	new := re.ReplaceAllString(s, "")
	i, err := strconv.Atoi(new)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
