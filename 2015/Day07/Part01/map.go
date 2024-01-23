package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    instructions := make(map[string]string)
    signals := make(map[string]uint16)

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

    fmt.Println("Signal provided to wire 'a':", getSignal("a", instructions, signals))
}

func getSignal(wire string, instructions map[string]string, signals map[string]uint16) uint16 {
    if val, ok := signals[wire]; ok {
        return val
    }

    if val, err := strconv.ParseUint(wire, 10, 16); err == nil {
        return uint16(val)
    }

    parts := strings.Fields(instructions[wire])
    var signal uint16

    switch len(parts) {
    case 1:
        signal = getSignal(parts[0], instructions, signals)
    case 2:
        signal = ^getSignal(parts[1], instructions, signals)
    case 3:
        a := getSignal(parts[0], instructions, signals)
        b := getSignal(parts[2], instructions, signals)
        switch parts[1] {
        case "AND":
            signal = a & b
        case "OR":
            signal = a | b
        case "LSHIFT":
            signal = a << b
        case "RSHIFT":
            signal = a >> b
        }
    }

    signals[wire] = signal
    return signal
}
