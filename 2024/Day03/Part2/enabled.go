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

	// Define regex patterns for the instructions
	mulPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)

	scanner := bufio.NewScanner(file)
	totalSum := 0
	mulEnabled := true // Multiplications start enabled

	// Read each line from the input file
	for scanner.Scan() {
		line := scanner.Text()

		// Check for do() and don't() instructions
		if doPattern.MatchString(line) {
			mulEnabled = true
		} else if dontPattern.MatchString(line) {
			mulEnabled = false
		}

		// If multiplications are enabled, find and process mul(X, Y) instructions
		if mulEnabled {
			matches := mulPattern.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				x, err1 := strconv.Atoi(match[1])
				y, err2 := strconv.Atoi(match[2])
				if err1 != nil || err2 != nil {
					fmt.Println("Error converting numbers:", err1, err2)
					return
				}
				totalSum += x * y
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total sum of enabled multiplications:", totalSum)
}
