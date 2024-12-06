package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(string(b), -1)

	sum := 0

	for _, match := range matches {
		var num1, num2 int
		fmt.Sscanf(match, "mul(%d,%d)", &num1, &num2)
		sum += num1 * num2
	}

	fmt.Println("Answer to part one:", sum)

	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	doMatchesIndexes := doRe.FindAllStringIndex(string(b), -1)
	dontMatchesIndexes := dontRe.FindAllStringIndex(string(b), -1)

	sum = 0
	enabled := true

	newStr := ""
	doMatchIndex := 0
	dontMatchIndex := 0

	for i := 0; i < len(b); {
		if doMatchIndex < len(doMatchesIndexes) && doMatchesIndexes[doMatchIndex][0] == i {
			enabled = true
			i = doMatchesIndexes[doMatchIndex][1]
			doMatchIndex++
		} else if dontMatchIndex < len(dontMatchesIndexes) && dontMatchesIndexes[dontMatchIndex][0] == i {
			enabled = false
			i = dontMatchesIndexes[dontMatchIndex][1]
			dontMatchIndex++
		} else {
			if enabled {
				newStr += string(b[i])
			}
			i++
		}
	}

	matches = mulRe.FindAllString(newStr, -1)

	sum = 0

	for _, match := range matches {
		var num1, num2 int
		fmt.Sscanf(match, "mul(%d,%d)", &num1, &num2)
		sum += num1 * num2
	}

	fmt.Println("Answer to part two:", sum)
}
