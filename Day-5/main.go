package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	rules := make(map[int][]Rule)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		flz := strings.Split(line, "|")
		before, _ := strconv.Atoi(flz[0])
		after, _ := strconv.Atoi(flz[1])
		if _, ok := rules[before]; !ok {
			rules[before] = make([]Rule, 0)
		}
		rule := Rule{Before: before, After: after}
		rules[before] = append(rules[before], rule)
	}

	orders := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fl := strings.Split(line, ",")
		order := make([]int, 0)
		for _, fa := range fl {
			skibi, _ := strconv.Atoi(fa)
			order = append(order, skibi)
		}
		orders = append(orders, order)
	}

	sum := 0

	incorrectOrders := make([]int, 0)

	for aa, order := range orders {
		isCorrect := true
		for i := 0; i < len(order); i++ {
			if _, ok := rules[order[i]]; ok {
				for _, rule := range rules[order[i]] {
					for j := 0; j < i; j++ {
						if rule.After == order[j] {
							isCorrect = false
							break
						}
					}
					if !isCorrect {
						break
					}
				}
				if !isCorrect {
					break
				}
			}
		}

		if isCorrect {
			sum += order[(len(order)-1)/2]
		} else {
			incorrectOrders = append(incorrectOrders, aa)
		}
	}

	fmt.Println("Answer to part one is", sum)

	for _, incorrectOrderInd := range incorrectOrders {
		incorrectOrder := orders[incorrectOrderInd]
		for i := 0; i < len(incorrectOrder); i++ {
			for j := 0; j < len(incorrectOrder)-i-1; j++ {
				for _, rule := range rules[incorrectOrder[j]] {
					if rule.After == incorrectOrder[j+1] {
						incorrectOrder[j], incorrectOrder[j+1] = incorrectOrder[j+1], incorrectOrder[j]
					}
				}
			}
		}
		orders[incorrectOrderInd] = incorrectOrder
	}

	sum = 0

	for _, order := range incorrectOrders {
		sum += orders[order][(len(orders[order])-1)/2]
	}

	fmt.Println("Answer to part two is", sum)
}
