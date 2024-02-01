package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	distances, err := readDistancesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	locations := getLocations(distances)
	shortestDistance := math.MaxInt64

	permute(locations, func(perm []string) {
		distance := calculateDistance(perm, distances)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	})

	fmt.Println("The shortest distance is:", shortestDistance)
}

func readDistancesFromFile(filename string) (map[string]map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	distances := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			continue // skip malformed lines
		}
		locations := strings.Split(parts[0], " to ")
		if len(locations) != 2 {
			continue // skip malformed lines
		}
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			continue // skip lines with invalid distances
		}

		if distances[locations[0]] == nil {
			distances[locations[0]] = make(map[string]int)
		}
		if distances[locations[1]] == nil {
			distances[locations[1]] = make(map[string]int)
		}
		distances[locations[0]][locations[1]] = distance
		distances[locations[1]][locations[0]] = distance
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return distances, nil
}

func getLocations(distances map[string]map[string]int) []string {
	locationSet := make(map[string]bool)
	for start := range distances {
		locationSet[start] = true
		for end := range distances[start] {
			locationSet[end] = true
		}
	}

	var locations []string
	for location := range locationSet {
		locations = append(locations, location)
	}
	return locations
}

func permute(values []string, callback func([]string)) {
	var helper func([]string, int)
	helper = func(arr []string, n int) {
		if n == 1 {
			callback(arr)
			return
		}
		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			} else {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			}
		}
	}
	helper(values, len(values))
}

func calculateDistance(route []string, distances map[string]map[string]int) int {
	sum := 0
	for i := 0; i < len(route)-1; i++ {
		sum += distances[route[i]][route[i+1]]
	}
	return sum
}
