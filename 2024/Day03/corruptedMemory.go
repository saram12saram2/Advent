package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define a regex pattern to match valid mul(X,Y) instructions
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Find all matches for the pattern in the current line
		matches := pattern.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// Extract X and Y from the match and convert them to integers
			x, err1 := strconv.Atoi(match[1])
			y, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting numbers:", err1, err2)
				return
			}

			// Multiply X and Y and add the result to the total sum
			totalSum += x * y
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the total sum of all valid mul(X,Y) instructions
	fmt.Println("Total sum of all valid multiplications:", totalSum)
}
