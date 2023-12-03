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
	str = strings.Replace(str, "\r\n", "", -1)
	lineSize := len(lines[0])
	fullSize := lineSize * len(lines)
	sum := 0

	
	for lineNumber, line := range lines {
		number := ""
		isPart := false
		lineLastNumber := false

		for i := 0; i < lineSize; i++ {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				number += string(char)
				if i == lineSize - 1 {
					lineLastNumber = true
				}
			}
			if !unicode.IsDigit(char) || lineLastNumber {
				for j := 0; j < len(number); j++ {
					index := lineNumber * lineSize + i - (len(number) - j)
					if lineLastNumber {
						index++
					}
					if j == 0 {
						if ((index - lineSize - 1 > 0 && index % lineSize != 0) && (!unicode.IsDigit(rune(str[index - lineSize - 1])) && string(str[index - lineSize - 1]) != ".")) ||
						   (index - lineSize > 0 && (!unicode.IsDigit(rune(str[index - lineSize])) && string(str[index - lineSize]) != ".")) ||
						   ((index - 1 > 0 && index % lineSize != 0) && (!unicode.IsDigit(rune(str[index - 1])) && string(str[index - 1]) != ".")) ||
						   ((index + lineSize - 1 < fullSize && index % lineSize != 0) && (!unicode.IsDigit(rune(str[index + lineSize - 1])) && string(str[index + lineSize - 1]) != ".")) ||
						   (index + lineSize < fullSize && (!unicode.IsDigit(rune(str[index + lineSize])) && string(str[index + lineSize]) != ".")) {
							isPart = true
						}
					}
					if j == len(number) - 1 {
						if ((index - lineSize + 1 > 0 && index % lineSize != lineSize - 1) && (!unicode.IsDigit(rune(str[index - lineSize + 1])) && string(str[index - lineSize + 1]) != ".")) ||
						   (index - lineSize > 0 && (!unicode.IsDigit(rune(str[index - lineSize])) && string(str[index - lineSize]) != ".")) ||
						   ((index + 1 < fullSize && index % lineSize != lineSize - 1) && (!unicode.IsDigit(rune(str[index + 1])) && string(str[index + 1]) != ".")) ||
						   ((index + lineSize + 1 < fullSize && index % lineSize != lineSize - 1) && (!unicode.IsDigit(rune(str[index + lineSize + 1])) && string(str[index + lineSize + 1]) != ".")) ||
						   (index + lineSize < fullSize && (!unicode.IsDigit(rune(str[index + lineSize])) && string(str[index + lineSize]) != ".")) {
							isPart = true
						}
					} else if j != 0 {
						if (index - lineSize > 0 && (!unicode.IsDigit(rune(str[index - lineSize])) && string(str[index - lineSize]) != ".")) ||
						   (index + lineSize < fullSize && (!unicode.IsDigit(rune(str[index + lineSize])) && string(str[index + lineSize]) != ".")) {
							isPart = true
						}
					}
	
					if isPart {
						intNum, err := strconv.Atoi(number)
						if err != nil {
							fmt.Printf("Error converting string to int: %s", err.Error())
							return
						}
	
						sum += intNum
						number = ""
						isPart = false
						break
					}
	
				}
				number = ""
			}
		}
	}

	fmt.Printf("The sum of the numbers: %d", sum)
}