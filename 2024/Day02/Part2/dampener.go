package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// check if a report is safe based on original rules
func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Adjacent levels must differ by 1, 2, or 3
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Levels must consistently increase or decrease
		if (diff > 0) != isIncreasing {
			return false
		}
	}
	return true
}

// Use the Problem Dampener to allow one bad level to be removed
func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	// Check each possible removal
	for i := 0; i < len(report); i++ {
		// Create a new report without the i-th level
		modified := append(report[:i], report[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	// each report
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Convert the line into a slice of integers
		report := make([]int, len(fields))
		for i, field := range fields {
			num, _ := strconv.Atoi(field)
			report[i] = num
		}

		// Check if the report is safe with the Problem Dampener
		if isSafeWithDampener(report) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

/*
// check if a report is safe
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
*/
