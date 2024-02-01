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

	var totalOriginal, totalEncoded int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		encoded := encodeString(line)
		totalOriginal += len(line)
		totalEncoded += len(encoded)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total difference:", totalEncoded-totalOriginal)
}

// encodeString returns the Go encoded representation of a string literal
func encodeString(s string) string {
	encoded := "\""
	for _, c := range s {
		switch c {
		case '\\', '"':
			encoded += "\\" + string(c)
		default:
			encoded += string(c)
		}
	}
	encoded += "\""
	return encoded
}
