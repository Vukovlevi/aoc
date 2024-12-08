package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
    operators = []string{"+", "*"}
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    lines := strings.Split(str, "\n")
    lines = lines[:len(lines) - 1]
    sum := 0

    for _, line := range lines {
        parts := strings.Split(line, ": ")
        numpart := strings.Split(parts[1], " ")
        eq, _ := strconv.Atoi(parts[0])
        nums := make([]int, len(numpart))

        for i := range len(numpart) {
            nums[i], _ = strconv.Atoi(numpart[i])
        }

        if tryOperators(eq, nums[0], nums[1:]) {
            sum += eq
        }
    }

    fmt.Printf("the sum of the possible equations: %d\n", sum)
}

func tryOperators(maxvalue, value int, remainingNums []int) bool {
    if value > maxvalue {
        return false
    }

    for _, op := range operators {
        val := value
        switch op {
        case "+":
            val += remainingNums[0]

            if len(remainingNums) == 1 && val == maxvalue {
                return true
            }

            if len(remainingNums) > 1 && tryOperators(maxvalue, val, remainingNums[1:]) {
                return true
            }
        case "*":
            val *= remainingNums[0]

            if len(remainingNums) == 1 && val == maxvalue {
                return true
            }

            if len(remainingNums) > 1 && tryOperators(maxvalue, val, remainingNums[1:]) {
                return true
            }
        }
    }

    return false
}
