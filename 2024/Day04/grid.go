package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the map type
type Grid map[[2]int]rune

// Helper to convert boolean to integer
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Part 1: Check if the "XMAS" pattern exists in all 8 directions
func xmas(grid Grid, r, c int) int {
	return b2i(xmasUp(grid, r, c)) +
		b2i(xmasUpRight(grid, r, c)) +
		b2i(xmasRight(grid, r, c)) +
		b2i(xmasDownRight(grid, r, c)) +
		b2i(xmasDown(grid, r, c)) +
		b2i(xmasDownLeft(grid, r, c)) +
		b2i(xmasLeft(grid, r, c)) +
		b2i(xmasUpLeft(grid, r, c))
}

func xmasUp(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r - 1, c}] == 'M' &&
		grid[[2]int{r - 2, c}] == 'A' && grid[[2]int{r - 3, c}] == 'S'
}

func xmasUpRight(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r - 1, c + 1}] == 'M' &&
		grid[[2]int{r - 2, c + 2}] == 'A' && grid[[2]int{r - 3, c + 3}] == 'S'
}

func xmasRight(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r, c + 1}] == 'M' &&
		grid[[2]int{r, c + 2}] == 'A' && grid[[2]int{r, c + 3}] == 'S'
}

func xmasDownRight(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r + 1, c + 1}] == 'M' &&
		grid[[2]int{r + 2, c + 2}] == 'A' && grid[[2]int{r + 3, c + 3}] == 'S'
}

func xmasDown(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r + 1, c}] == 'M' &&
		grid[[2]int{r + 2, c}] == 'A' && grid[[2]int{r + 3, c}] == 'S'
}

func xmasDownLeft(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r + 1, c - 1}] == 'M' &&
		grid[[2]int{r + 2, c - 2}] == 'A' && grid[[2]int{r + 3, c - 3}] == 'S'
}

func xmasLeft(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r, c - 1}] == 'M' &&
		grid[[2]int{r, c - 2}] == 'A' && grid[[2]int{r, c - 3}] == 'S'
}

func xmasUpLeft(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'X' && grid[[2]int{r - 1, c - 1}] == 'M' &&
		grid[[2]int{r - 2, c - 2}] == 'A' && grid[[2]int{r - 3, c - 3}] == 'S'
}

// Part 2: Check the X-MAS shape
func xmasPart2(grid Grid, r, c int) bool {
	return xmasDiagonal1(grid, r, c) && xmasDiagonal2(grid, r, c)
}

func xmasDiagonal1(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'A' &&
		((grid[[2]int{r + 1, c - 1}] == 'M' && grid[[2]int{r - 1, c + 1}] == 'S') ||
			(grid[[2]int{r + 1, c - 1}] == 'S' && grid[[2]int{r - 1, c + 1}] == 'M'))
}

func xmasDiagonal2(grid Grid, r, c int) bool {
	return grid[[2]int{r, c}] == 'A' &&
		((grid[[2]int{r + 1, c + 1}] == 'M' && grid[[2]int{r - 1, c - 1}] == 'S') ||
			(grid[[2]int{r + 1, c + 1}] == 'S' && grid[[2]int{r - 1, c - 1}] == 'M'))
}

func main() {
	// Read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	grid := make(Grid)
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for col, char := range line {
			grid[[2]int{row, col}] = char
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Part 1
	part1 := 0
	for pos := range grid {
		r, c := pos[0], pos[1]
		part1 += xmas(grid, r, c)
	}
	fmt.Println("Part 1:", part1)

	// Part 2
	part2 := 0
	for pos := range grid {
		r, c := pos[0], pos[1]
		part2 += b2i(xmasPart2(grid, r, c))
	}
	fmt.Println("Part 2:", part2)
}
