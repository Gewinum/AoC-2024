package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    floors := make([][]int, 0)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        flz := strings.Split(line, " ")
        floor := make([]int, 0)
        for _, fl := range flz {
            fll, _ := strconv.Atoi(fl)
            floor = append(floor, fll)
        }
        floors = append(floors, floor)
    }

    safe := 0

    for _, floor := range floors {
        // in part one of day 2 canBeSafeByRemovingOne was not called
        if isSafe(floor) || canBeSafeByRemovingOne(floor) {
            safe++
        }
    }

    fmt.Println(safe)
}

func isSafe(floor []int) bool {
    increasing := true
    decreasing := true

    for i := 1; i < len(floor); i++ {
        diff := floor[i] - floor[i-1]
        if diff < 1 || diff > 3 {
            increasing = false
        }
        if diff > -1 || diff < -3 {
            decreasing = false
        }
    }

    return increasing || decreasing
}

func canBeSafeByRemovingOne(floor []int) bool {
    for i := 0; i < len(floor); i++ {
        temp := append([]int{}, floor[:i]...)
        temp = append(temp, floor[i+1:]...)
        if isSafe(temp) {
            return true
        }
    }
    return false
}
