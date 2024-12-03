package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    sum := 0

    index := 0
    for index != -1 && index < len(str) {
        index = strings.Index(str, "mul(")
        if index == -1 {
            continue
        }

        end := index + 4 + 8
        if end > len(str) {
            end = len(str)
        }
        num, i := parseMul(str[index + 4:end])
        if num != -1 {
            sum += num
        }
        str = str[index + 4 + i:]
    }


    fmt.Printf("the sum of the muls: %d\n", sum)
}

func parseMul(data string) (int, int) {
    num1Str := ""
    i := 0
    for i < 3 && unicode.IsDigit(rune(data[i])) {
        num1Str += string(data[i])
        i++
    }

    if data[i] != ',' {
        return -1, i + 1
    }

    i++
    j := i + 3

    num2Str := ""
    for i < j && unicode.IsDigit(rune(data[i])) {
        num2Str += string(data[i])
        i++
    }

    if data[i] != ')' {
        return -1, i + 1
    }

    num1, err := strconv.Atoi(num1Str)
    if err != nil {
        return -1, i + 1
    }
    num2, err := strconv.Atoi(num2Str)
    if err != nil {
        return -1, i + 1
    }

    return num1 * num2, i + 1
}
