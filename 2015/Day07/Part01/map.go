package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	instructions := make(map[string]string)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		instructions[parts[1]] = parts[0]
	}
}
