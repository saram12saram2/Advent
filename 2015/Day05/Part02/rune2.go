package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Check if a string contains a pair of any two letters that appears at least twice without overlapping
func hasPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

// Check if a string contains at least one letter which repeats with exactly one letter between them
func hasRepeatWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}



func isNice(s string) bool {
	return hasPair(s) && hasRepeatWithOneBetween(s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	niceCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if isNice(line) {
			niceCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} else {
		fmt.Printf("Number of nice strings: %d\n", niceCount)
	}
}



// compile -> cat input.txt | go run your_program.go
