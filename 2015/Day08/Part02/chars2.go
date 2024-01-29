package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var totalOriginalChars, totalEncodedChars int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalOriginalChars += len(line)

		// Encode the line
		encodedLine := "\"" + strings.ReplaceAll(line, "\\", "\\\\")
		encodedLine = strings.ReplaceAll(encodedLine, "\"", "\\\"") + "\""

		totalEncodedChars += len(encodedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total Encoded Characters - Total Original Characters:", totalEncodedChars-totalOriginalChars)
}
