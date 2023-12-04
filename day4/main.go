package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")
	sum := 0
	lineCount := make([]int, len(lines))
	for i := 0; i < len(lineCount); i++ {
		lineCount[i] = 1
	}
	var lineMatches []int
	cards := 0

	for _, line := range lines {
		var winningNumbers []int
		var yourNumbers []int
		isAdd := false
		addTo := &winningNumbers

		for i := 0; i < len(line); i++ {
			char := rune(line[i])
			if string(char) == ":" {
				isAdd = true
			}

			if !isAdd {
				continue
			}

			if string(char) == "|" {
				addTo = &yourNumbers
				continue
			}

			if unicode.IsSpace(char) {
				continue
			}

			number := ""
			for i < len(line) && unicode.IsDigit(rune(line[i])) {
				number += string(line[i])
				i++
			}

			if number != "" {
				intNum, err := strconv.Atoi(number)
				if err != nil {
					fmt.Printf("Error converting number: %s", err.Error())
					return
				}

				*addTo = append(*addTo, intNum)
			}
		}

		matches := 0
		for _, number := range yourNumbers {
			found := false
			i := 0
			for !found && i < len(winningNumbers) {
				if winningNumbers[i] == number {
					matches++
					found = true
				}

				i++
			}
		}

		lineMatches = append(lineMatches, matches)
		sum += int(math.Pow(2, float64(matches) - 1))
	}


	for i, match := range lineMatches {
		for j := 0; j < lineCount[i]; j++ {
			for k := 1; k <= match; k++ {
				if i + k < len(lineCount) {
					lineCount[i + k]++
				}
			}
		}
	}

	for _, count := range lineCount {
		cards += count
	}

	fmt.Printf("The sum of the points: %d\n", sum)
	fmt.Printf("The number of cards: %d", cards)
}