package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read the entire input
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	text := string(data)

	// Compile a regex that matches either a valid mul(...) instruction or do() / don't()
	// For mul: mul(<digits>,<digits>)
	// For do: do()
	// For don't: don't()
	pattern := `(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringIndex(text, -1)

	enabled := true
	sum := 0

	for _, match := range matches {
		token := text[match[0]:match[1]]

		// Check which type of token we got
		if strings.HasPrefix(token, "mul(") {
			// Extract the numbers
			// token format: mul(X,Y)
			inside := token[4 : len(token)-1] // get the part inside parentheses
			parts := strings.Split(inside, ",")
			if len(parts) == 2 {
				x, errX := strconv.Atoi(parts[0])
				y, errY := strconv.Atoi(parts[1])
				if errX == nil && errY == nil && enabled {
					sum += x * y
				}
			}
		} else if token == "do()" {
			enabled = true
		} else if token == "don't()" {
			enabled = false
		}
	}

	fmt.Println(sum)
}
