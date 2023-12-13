package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkLine(line *string, lineCheck []int) bool {
	currentLine := *line

	var currentCheck []int

	for i := 0; i < len(currentLine); i++ {
		num := 0
		for i < len(currentLine) && string(currentLine[i]) == "#" {
			num++
			i++
		}

		if num != 0 {
			currentCheck = append(currentCheck, num)
		}
	}

	if len(lineCheck) != len(currentCheck) {
		return false
	}

	for i := range lineCheck {
		if lineCheck[i] != currentCheck[i] {
			return false
		}
	}

	return true
}

func tryPossibilities(possibilities *int, line *string, unknownIndexes []int, index int, lineCheck []int) {
	if index == len(unknownIndexes) {
		if checkLine(line, lineCheck) {
			*possibilities++
		}

		return
	}

	states := []string{".", "#"}

	for _, state := range states {
		currentLine := *line
		currentIndex := unknownIndexes[index]
		if currentIndex == len(currentLine) - 1 {
			currentLine = currentLine[:currentIndex] + state
		} else {
			currentLine = currentLine[:currentIndex] + state + currentLine[currentIndex + 1:]
		}

		tryPossibilities(possibilities, &currentLine, unknownIndexes, index + 1, lineCheck)
	}
}

func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)

	lines := strings.Split(str, "\r\n")
	var unknowns [][]int
	var lineChecks [][]int
	possibilities := 0

	for lineIndex, line := range lines {
		var unknown []int
		var checks []int
		isCheck := false
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "?" {
				unknown = append(unknown, i) 
			}

			if string(line[i]) == " " {
				isCheck = true
				continue
			}
			
			if !isCheck {
				continue
			}

			strNum := ""
			for i < len(line) && unicode.IsDigit(rune(line[i])) {
				strNum += string(line[i])
				i++
			}

			intNum, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error converting check number: %s", err.Error())
				return
			}

			strNum = ""
			checks = append(checks, intNum)
		}

		currentLine := strings.Split(line, " ")[0]
		lines[lineIndex] = currentLine
		currentCheck := checks
		currentUnknown := unknown
		for i := 0; i < 4; i++ {
			lines[lineIndex] += "?" + currentLine
			checks = append(checks, currentCheck...)
			for j := 0; j < len(currentUnknown); j++ {
				unknown = append(unknown, unknown[len(unknown) - len(currentUnknown)] + len(currentLine) + 1)
			}
		}
		lineChecks = append(lineChecks, checks)
		unknowns = append(unknowns, unknown)
	}

	fmt.Println(lines, lineChecks, unknowns)

	for i, line := range lines {
		tryPossibilities(&possibilities, &line, unknowns[i], 0, lineChecks[i])
	}

	fmt.Printf("The number of possibilities: %d", possibilities)
}