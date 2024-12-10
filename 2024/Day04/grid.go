package main

import (
	"bufio"
	"fmt"
	"os"
)

// Check if "XMAS" starts at a given position in the grid
func checkWord(grid [][]rune, word string, row, col, rowDir, colDir int) bool {
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		r := row + i*rowDir
		c := col + i*colDir

		// Check bounds
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return false
		}
		// Check character match
		if grid[r][c] != rune(word[i]) {
			return false
		}
	}
	return true
}

// Function to count all occurrences of "XMAS" in the grid
func countOccurrences(grid [][]rune, word string) int {
	count := 0
	directions := [][2]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Diagonal Down-Right
		{-1, -1}, // Diagonal Up-Left
		{1, -1},  // Diagonal Down-Left
		{-1, 1},  // Diagonal Up-Right
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			for _, dir := range directions {
				if checkWord(grid, word, row, col, dir[0], dir[1]) {
					count++
				}
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

	word := "XMAS"

	// Count occurrences
	totalOccurrences := countOccurrences(grid, word)

	fmt.Println("Total occurrences of 'XMAS':", totalOccurrences)
}
