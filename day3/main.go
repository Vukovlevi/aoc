package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    sum := 0
    dos := strings.Split(str, "do()")
    for i, do := range dos {
        dos[i] = strings.Split(do, "don't()")[0]
    }

    for _, exec := range dos {
        sum += parseMul(exec)
    }

    fmt.Printf("the sum of the muls: %d\n", sum)
}

func parseMul(data string) int {
    re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
    matches := re.FindAllStringSubmatch(data, -1)
    sum := 0
    for _, match := range matches {
        num1, _ := strconv.Atoi(match[1])
        num2, _ := strconv.Atoi(match[2])
        sum += num1 * num2
    }
    return sum
}
