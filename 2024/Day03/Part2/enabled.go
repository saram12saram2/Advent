package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input file line by line
	input := readFileLineByLine("input.txt")

	var occurrences [][]string
	for _, line := range input {
		occurrences = append(occurrences, findValidMulsDoDonts(line)...)
	}
	fmt.Println("Answer for part 2:", calcSum(occurrences))
}

// readFileLineByLine reads the input file and returns a slice of lines
func readFileLineByLine(filename string) []string {
	// Read the file content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// Split lines and return
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

// finds all occurrences of mul(X,Y), do(), and don't() in the input
func findValidMulsDoDonts(input string) [][]string {
	// Regex to find mul(X,Y), do(), and don't()
	r := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|don't\(\)|do\(\)`)
	return r.FindAllStringSubmatch(input, -1)
}

// extracts numbers from mul(X,Y) and returns their product
func findProduct(input string) int {
	// Extract numbers from mul(X,Y) and return their product
	r := regexp.MustCompile(`\d+`)
	nums := r.FindAllString(input, -1)
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return x * y
}

// calculates the sum of products with do()/don't() conditions
func calcSum(input [][]string) int {
	sum := 0
	active := true // Multiplications are active by default
	for _, match := range input {
		if strings.Contains(match[0], "don't") {
			active = false
		} else if strings.Contains(match[0], "do(") {
			active = true
		} else if active {
			sum += findProduct(match[0])
		}
	}
	return sum
}
