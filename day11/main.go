package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
    blinkCount = 75
    multiplier = 2024
)

var (
    cache = make(map[string]int)
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    stones := strings.Split(strings.TrimSpace(str), " ")
    sum := 0

    for _, stone := range stones {
        sum += blink(stone, blinkCount)
    }


    fmt.Printf("the number of stones: %d\n", sum)
}

func blink(stone string, iteration int) int {
    if iteration == 0 {
        return 1
    }

    if v, ok := cache[fmt.Sprintf("%d-%s", iteration, stone)]; ok {
        return v
    }

    if stone == "0" {
        res := blink("1", iteration - 1)
        cache[fmt.Sprintf("%d-%s", iteration, stone)] = res
        return res
    }

    if len(stone) % 2 == 0 {
        halfIndex := len(stone) / 2
        num, _ := strconv.Atoi(stone[halfIndex:])
        res := blink(stone[:halfIndex], iteration - 1) + blink(strconv.Itoa(num), iteration - 1)
        cache[fmt.Sprintf("%d-%s", iteration, stone)] = res
        return res
    }

    num, _ := strconv.Atoi(stone)
    num *= multiplier
    res := blink(strconv.Itoa(num), iteration - 1)
    cache[fmt.Sprintf("%d-%s", iteration, stone)] = res
    return res
}
