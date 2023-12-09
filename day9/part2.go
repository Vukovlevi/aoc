package main

import (
	"fmt"
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
	var values [][]int
	sum := 0

	for _, line := range lines {
		var lineNums []int
		strNum := ""
		for i := 0; i < len(line); i++ {
			endOfLine := false
			for !endOfLine && (unicode.IsDigit(rune(line[i])) || string(line[i]) == "-") {
				strNum += string(line[i])
				i++

				if i == len(line) {
					endOfLine = true
				}
			}

			intNum, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error converting number: %s", err.Error())
				return
			}

			lineNums = append(lineNums, intNum)
			strNum = ""
		}

		values = append(values, lineNums)
	}

	for _, value := range values {
		var differences [][]int
		differences = append(differences, value)
		endOfTree := false
		i := 0
		for !endOfTree {
			var difference []int
			nulls := 0
			currentValue := differences[i]
			for j := 0; j < len(currentValue) - 1; j++ {
				dif := currentValue[j + 1] - currentValue[j]
				difference = append(difference, dif)

				if dif == 0 {
					nulls++
				}
			}

			if nulls == len(difference) {
				endOfTree = true
			} else {
				differences = append(differences, difference)
				i++
			}
		}

		startNum := differences[len(differences) - 1][0]
		for i = len(differences) - 2; i >= 0; i-- {
			startNum = differences[i][0] - startNum
		}

		sum += startNum
	}

	fmt.Printf("The sum of the next values: %d", sum)
}