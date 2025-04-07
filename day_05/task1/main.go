package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := solve("../input.txt")
	fmt.Println("result: ", result)
}

func solve(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	first_section := true
	var page_ordering_rules []Pair
	var updates [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			first_section = false
			continue
		}
		if first_section {
			nums := strings.Split(line, "|")
			page_ordering_rules = append(page_ordering_rules, Pair{nums[0], nums[1]})
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	c := make(chan int)
	for i := range updates {
		go processUpdate(updates[i], page_ordering_rules, c)
	}

	var total int
	for i := 0; i < len(updates); i++ {
		x := <-c
		total += x
	}
	return total
}

func processUpdate(update []string, page_ordering_rules []Pair, c chan int) {
	if !is_correct_order(update, page_ordering_rules) {
		c <- 0
		return
	}
	middle_num := update[len(update)/2]
	println("Middle num ", middle_num)
	c <- Atoi(middle_num)
}

func is_correct_order(update []string, page_ordering_rules []Pair) bool {
	all_nums := make(map[string]bool)
	for i := range update {
		all_nums[update[i]] = true
	}
	before_nums := make(map[string]bool)
	for i := range update {
		for j := range page_ordering_rules {
			if update[i] == page_ordering_rules[j].b {
				if all_nums[page_ordering_rules[j].a] && !before_nums[page_ordering_rules[j].a] {
					return false
				}
			}
		}
		before_nums[update[i]] = true
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

type Pair struct {
	a string
	b string
}
