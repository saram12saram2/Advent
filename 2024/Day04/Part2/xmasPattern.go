package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to check if a position forms an X-MAS pattern
func isXMAS(grid [][]rune, row, col int) bool {
	// Dimensions of the grid
	rows := len(grid)
	cols := len(grid[0])

	// Define the positions for the "X" shape
	deltas := [][2]int{
		{-1, -1}, // Top-left M
		{1, -1},  // Bottom-left M
		{-1, 1},  // Top-right M
		{1, 1},   // Bottom-right M
		{0, 0},   // Center A
	}

	// Check "M.S MAS" orientation
	matches := [5]rune{'M', 'M', 'M', 'M', 'A'}
	for i, delta := range deltas {
		r := row + delta[0]
		c := col + delta[1]
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != matches[i] {
			break
		}
		if i == len(matches)-1 {
			return true
		}
	}

	// Check "S.M SAM" reversed orientation
	matches = [5]rune{'S', 'S', 'S', 'S', 'A'}
	for i, delta := range deltas {
		r := row + delta[0]
		c := col + delta[1]
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != matches[i] {
			break
		}
		if i == len(matches)-1 {
			return true
		}
	}

	return false
}

// Function to count all X-MAS patterns in the grid
func countXMAS(grid [][]rune) int {
	count := 0
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {
			if isXMAS(grid, row, col) {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the grid
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Count all X-MAS patterns
	totalXMAS := countXMAS(grid)

	fmt.Println("Total X-MAS patterns:", totalXMAS)
}
