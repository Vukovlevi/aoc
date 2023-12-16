package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(array []int) int {
	minimum := array[0]
	for i := 1; i < len(array); i++ {
		if minimum > array[i] {
			minimum = array[i]
		}
	}

	return minimum
}

func max(array []int) int {
	maximum := array[0]
	for i := 1; i < len(array); i++ {
		if maximum < array[i] {
			maximum = array[i]
		}
	}

	return maximum
}

func getNumberFromString(s string) int {
	switch (s) {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input, %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\n")

	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0
	for _, line := range lines {
		numbers := make(map[int]int)
		var indexes []int
		for _, strDigit := range digits {
			index := strings.Index(line, strDigit)
			if index != -1 {
				numbers[index] = getNumberFromString(strDigit)
				indexes = append(indexes, index)
			}

			index = strings.LastIndex(line, strDigit)
			if index != -1 {
				numbers[index] = getNumberFromString(strDigit)
				indexes = append(indexes, index)
			}
		}
		for i := 0; i < len(line); i++ {
			if digit, err := strconv.Atoi(string(line[i])); err == nil {
				numbers[i] = digit
				indexes = append(indexes, i)
			}
		}

		minimum := min(indexes)
		maximum := max(indexes)

		str := strconv.Itoa(numbers[minimum]) + strconv.Itoa(numbers[maximum])
		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting int to string: %s", err.Error())
			return
		}

		sum += number
	}

	fmt.Printf("The sum is: %d", sum)
}