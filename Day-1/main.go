package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	l1 := make([]int, 0)
	l2 := make([]int, 0)

	for {
		var n1, n2 int
		_, err := fmt.Fscanf(f, "%d   %d\n", &n1, &n2)
		if err != nil {
			break
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	distance := 0

	sort.Ints(l1)
	sort.Ints(l2)

	for i := 0; i < len(l1); i++ {
		distance += abs(l1[i] - l2[i])
	}

	fmt.Println("Answer to part one is", distance)

	similarity := 0

	for i := 0; i < len(l1); i++ {
		temp := 0
		for j := 0; j < len(l2); j++ {
			if l1[i] == l2[j] {
				temp++
			}
		}
		similarity += l1[i] * temp
	}

	fmt.Println("Answer to part two is", similarity)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
