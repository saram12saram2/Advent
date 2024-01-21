package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
    const size = 1000
    var grid [size][size]bool

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
    for scanner.Scan() {
        line := scanner.Text()
        matches := re.FindStringSubmatch(line)

        x1, _ := strconv.Atoi(matches[1])
        y1, _ := strconv.Atoi(matches[2])
        x2, _ := strconv.Atoi(matches[3])
        y2, _ := strconv.Atoi(matches[4])

        for x := x1; x <= x2; x++ {
            for y := y1; y <= y2; y++ {
                switch {
                case line[:6] == "toggle":
                    grid[x][y] = !grid[x][y]
                case line[:7] == "turn on":
                    grid[x][y] = true
                case line[:8] == "turn off":
                    grid[x][y] = false
                }
            }
        }
    }

    count := 0
    for x := 0; x < size; x++ {
        for y := 0; y < size; y++ {
            if grid[x][y] {
                count++
            }
        }
    }

    fmt.Println("Number of lights that are lit:", count)
}
