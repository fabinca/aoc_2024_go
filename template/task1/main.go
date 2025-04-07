package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total
}
