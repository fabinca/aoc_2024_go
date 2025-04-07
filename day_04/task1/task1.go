package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	result := day4_1("../input.txt")
	fmt.Println("result: ", result)
}

func day4_1(inputfile string) int {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	c := make(chan int)
	var wg sync.WaitGroup

	var coor []Coordinate
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'X' {
				coor = append(coor, Coordinate{row, col})
				wg.Add(1)
				go func() {
					defer wg.Done()
					find_xmas(grid, Coordinate{row, col}, c)
				}()

			}
		}
	}
	for i := 0; i < len(coor); i++ {
		x := <-c
		total += x
	}
	return total
}

type Coordinate struct {
	row int
	col int
}

func find_xmas(grid []string, coor Coordinate, c chan int) {
	xmas_sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			xmas := 0
			for k := 1; k <= 3; k++ {
				xmas += check_letter(grid, coor.row+i*k, coor.col+j*k, "XMAS"[k])
			}
			if xmas == 3 {
				xmas_sum += 1

			}
		}
	}
	c <- xmas_sum
}

func check_letter(grid []string, row int, col int, letter byte) int {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) {
		return 0
	}
	if grid[row][col] == letter {
		return 1
	}
	return 0

}
