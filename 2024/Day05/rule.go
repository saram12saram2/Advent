package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Run:  go run rule.go input.txt
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ruleRegex := regexp.MustCompile(`^(\d+)\|(\d+)$`)

	var rules [][2]int
	var updates [][]int

	// Read until we no longer get a line matching X|Y to collect rules
	readingRules := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if readingRules {
			// Check if line matches the rule pattern
			m := ruleRegex.FindStringSubmatch(line)
			if m == nil {
				// No longer a rule line, start updates
				readingRules = false
				// Process this line as an update
				update := parseUpdate(line)
				if len(update) > 0 {
					updates = append(updates, update)
				}
			} else {
				// It's a rule
				X, _ := strconv.Atoi(m[1])
				Y, _ := strconv.Atoi(m[2])
				rules = append(rules, [2]int{X, Y})
			}
		} else {
			// We're reading updates
			update := parseUpdate(line)
			if len(update) > 0 {
				updates = append(updates, update)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Now we have all rules and all updates
	// For each update, check if it's correctly ordered
	sum := 0
	for _, update := range updates {
		if isCorrectlyOrdered(update, rules) {
			// Find middle page
			middleIndex := len(update) / 2
			sum += update[middleIndex]
		}
	}

	fmt.Println(sum)
}

// parseUpdate parses a line of comma-separated page numbers into a slice of ints.
func parseUpdate(line string) []int {
	parts := strings.Split(line, ",")
	var update []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		num, err := strconv.Atoi(p)
		if err != nil {
			// If there's any non-integer, ignore (though puzzle implies always integers)
			continue
		}
		update = append(update, num)
	}
	return update
}

// isCorrectlyOrdered checks the given update against the rules.
// We only consider rules where both pages are in the update.
func isCorrectlyOrdered(update []int, rules [][2]int) bool {
	// Create a map for quick index lookup
	indexMap := make(map[int]int, len(update))
	for i, page := range update {
		indexMap[page] = i
	}

	for _, r := range rules {
		X, Y := r[0], r[1]
		ix, xFound := indexMap[X]
		iy, yFound := indexMap[Y]
		if xFound && yFound {
			if ix >= iy {
				// This violates the order rule
				return false
			}
		}
	}

	return true
}
