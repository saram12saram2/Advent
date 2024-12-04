package main

import (
	"bufio"
	"fmt"
	"os"
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
		// split each line into two numbers
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

	// count occurrences of each number in the right list
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}

	// calculate the similarity score
	similarityScore := 0
	for _, num := range left {
		similarityScore += num * rightCount[num]
	}

	fmt.Println("Similarity Score:", similarityScore)
}
