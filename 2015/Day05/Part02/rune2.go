package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeiou", ch)
}

func isNice(s string) bool {
	vowelCount, doubleLetter, containsBadStr := 0, false, false

	for _, badStr := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, badStr) {
			containsBadStr = true
			break
		}
	}

	for i, ch := range s {
		if isVowel(ch) {
			vowelCount++
		}
		if i > 0 && ch == rune(s[i-1]) {
			doubleLetter = true
		}
	}

	return vowelCount >= 3 && doubleLetter && !containsBadStr
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
