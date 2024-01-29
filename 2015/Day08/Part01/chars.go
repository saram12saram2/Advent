package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var totalCodeChars, totalMemoryChars int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalCodeChars += len(line)

		// Remove surrounding quotes
		line = line[1 : len(line)-1]

		// Process escape sequences
		for i := 0; i < len(line); i++ {
			if line[i] == '\\' {
				if line[i+1] == 'x' {
					i += 3 // Skip hexadecimal escape
				} else {
					i++ // Skip escaped character
				}
			}
			totalMemoryChars++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total Code Characters - Total Memory Characters:", totalCodeChars-totalMemoryChars)
}
