package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parseTopLeft(i int, s string, lineSize int) string {
	index := i - lineSize - 1
	if !unicode.IsDigit(rune(s[index])) {
		return ""
	}

	number := ""
	for unicode.IsDigit(rune(s[index])) {
		number = string(s[index]) + number
		index--
	}
	index = i - lineSize
	for unicode.IsDigit(rune(s[index])) {
		number += string(s[index])
		index++
	}

	return number
}

func parseTopRight(i int, s string, lineSize int) string {
	index := i - lineSize + 1
	if !unicode.IsDigit(rune(s[index])) || unicode.IsDigit(rune(s[index - 2])) {
		return ""
	}

	number := ""
	for unicode.IsDigit(rune(s[index])) {
		number += string(s[index])
		index++
	}
	index = i - lineSize
	for unicode.IsDigit(rune(s[index])) {
		number = string(s[index]) + number
		index--
	}

	return number
}

func parseTop(i int, s string, lineSize int) string {
	index := i - lineSize
	number := ""

	if unicode.IsDigit(rune(s[index])) {
		number = string(s[index])
	}

	return number
}

func parseLeft(i int, s string, lineSize int) string {
	index := i - 1
	number := ""
	for unicode.IsDigit(rune(s[index])) {
		number = string(s[index]) + number
		index--
	}

	return number
}

func parseRight(i int, s string, lineSize int) string {
	index := i + 1
	number := ""
	for unicode.IsDigit(rune(s[index])) {
		number += string(s[index])
		index++
	}

	return number
}

func parseBottomLeft(i int, s string, lineSize int) string {
	index := i + lineSize - 1
	if !unicode.IsDigit(rune(s[index])) {
		return ""
	}

	number := ""
	for unicode.IsDigit(rune(s[index])) {
		number = string(s[index]) + number
		index--
	}
	index = i + lineSize
	for unicode.IsDigit(rune(s[index])) {
		number += string(s[index])
		index++
	}

	return number
}

func parseBottomRight(i int, s string, lineSize int, fullSize int) string {
	index := i + lineSize + 1
	if !unicode.IsDigit(rune(s[index])) || unicode.IsDigit(rune(s[index - 2])) {
		return ""
	}

	number := ""
	for index < fullSize && unicode.IsDigit(rune(s[index])) {
		number += string(s[index])
		index++
	}
	index = i + lineSize
	for unicode.IsDigit(rune(s[index])) {
		number = string(s[index]) + number
		index--
	}
	

	return number
}

func parseBottom(i int, s string, lineSize int) string {
	index := i + lineSize
	number := ""

	if unicode.IsDigit(rune(s[index])) {
		number = string(s[index])
	}

	return number
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")
	str = strings.Replace(str, "\r\n", "", -1)
	lineSize := len(lines[0])
	fullSize := len(lines) * lineSize

	sum := 0
	for i, char := range str {
		var numbers []int
		numberStr := ""
		if string(char) == "*" {
			numberStr = parseTopLeft(i, str, lineSize)
			if numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error parsing number: %s", err.Error())
					return
				}
				numbers = append(numbers, number)
				numberStr = ""
			}

			numberStr = parseTopRight(i, str, lineSize)
			if numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error parsing number: %s", err.Error())
					return
				}
				numbers = append(numbers, number)
				numberStr = ""
			}

			if len(numbers) == 0 {
				numberStr = parseTop(i, str, lineSize)
				if numberStr != "" {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						fmt.Printf("Error parsing number: %s", err.Error())
						return
					}
					numbers = append(numbers, number)
					numberStr = ""
				}	
			}

			numberStr = parseLeft(i, str, lineSize)
			if numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error parsing number: %s", err.Error())
					return
				}
				numbers = append(numbers, number)
				numberStr = ""
			}

			numberStr = parseRight(i, str, lineSize)
			if numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error parsing number: %s", err.Error())
					return
				}
				numbers = append(numbers, number)
				numberStr = ""
			}

			numberStr = parseBottomLeft(i, str, lineSize)
			if numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error parsing number: %s", err.Error())
					return
				}
				numbers = append(numbers, number)
				numberStr = ""
			}

			numberStr = parseBottomRight(i, str, lineSize, fullSize)
			if numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error parsing number: %s", err.Error())
					return
				}
				numbers = append(numbers, number)
				numberStr = ""
			}

			if parseBottomLeft(i, str, lineSize)  == "" && parseBottomRight(i, str, lineSize, fullSize) == "" {
				numberStr = parseBottom(i, str, lineSize)
				if numberStr != "" {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						fmt.Printf("Error parsing number: %s", err.Error())
						return
					}
					numbers = append(numbers, number)
					numberStr = ""
				}
			}

			if len(numbers) >= 2 {
				sum += numbers[0] * numbers[1]
			}
		}
	}

	fmt.Printf("The sum of ratios: %d", sum)
}