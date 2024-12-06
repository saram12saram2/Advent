package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to check if a report is safe
func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check if the difference is within the valid range
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Check if the sequence maintains consistent direction
		if (diff > 0) != isIncreasing {
			return false
		}
	}

	return true
}

// Function to check if a report can be made safe by removing one level
func canBeMadeSafe(report []int) bool {
	for i := 0; i < len(report); i++ {
		// Create a new slice excluding the current level
		modified := append(report[:i], report[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}
	return false
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	// Process each line in the input file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Convert the line into a slice of integers
		report := make([]int, len(fields))
		for i, field := range fields {
			report[i], _ = strconv.Atoi(field)
		}

		// Check if the report is safe or can be made safe
		if isSafe(report) || canBeMadeSafe(report) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the number of safe reports
	fmt.Println("Number of safe reports:", safeCount)
}
