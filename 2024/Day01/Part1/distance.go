package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split each line into two numbers
		parts := strings.Fields(line)
		if len(parts) == 2 {
			leftNum, _ := strconv.Atoi(parts[0])
			rightNum, _ := strconv.Atoi(parts[1])
			left = append(left, leftNum)
			right = append(right, rightNum)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort both slices
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := abs(left[i] - right[i])
		totalDistance += distance
	}

	fmt.Println("Total Distance:", totalDistance)
	// result: 2166959
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
