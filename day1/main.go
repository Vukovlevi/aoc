package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)

    lines := strings.Split(str, "\n")
    left := make([]int, len(lines))
    right := make([]int, len(lines))

    for i, id := range lines[:len(lines) - 1] {
        parts := strings.Split(id, "   ")
        num1, _ := strconv.Atoi(parts[0])
        left[i] = num1
        num2, _ := strconv.Atoi(parts[1])
        right[i] = num2
    }

    slices.Sort(left)
    slices.Sort(right)

    difSum := 0
    for i, _ := range left {
        dif := right[i] - left[i]
        if dif < 0 {
            dif *= -1
        }
        difSum += dif
    }

    fmt.Printf("the dif is: %d\n", difSum)

    //part 2
    appears := make(map[int]int)

    for _, num := range right {
        if _, ok := appears[num]; !ok {
            appears[num] = 1
            continue
        }
        appears[num]++
    }

    sum := 0
    for _, num := range left {
        if multiplier, ok := appears[num]; ok {
            sum += (num * multiplier)
        }
    }

    fmt.Printf("the sum is: %d", sum)
}
