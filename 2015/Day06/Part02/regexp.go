package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	const size = 1000
	var grid [size][size]int // Change to integer for brightness levels

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch {
				case line[:6] == "toggle":
					grid[x][y] += 2
				case line[:7] == "turn on":
					grid[x][y] += 1
				case line[:8] == "turn off":
					if grid[x][y] > 0 { // Ensure brightness does not go below zero
						grid[x][y] -= 1
					}
				}
			}
		}
	}

	totalBrightness := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			totalBrightness += grid[x][y]
		}
	}

	fmt.Println("Total brightness of all lights:", totalBrightness)
}
