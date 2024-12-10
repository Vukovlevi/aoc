package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
    directions = [][]int{
        {0, 1},
        {1, 0},
        {0, -1},
        {-1, 0},
    }
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    lines := strings.Split(str, "\n")
    lines = lines[:len(lines) - 1]
    heightmap := make([][]int, len(lines))
    sum := 0
    ratingSum := 0

    for i, line := range lines {
        heights := make([]int, 0)
        parts := strings.Split(line, "")
        for _, part := range parts {
            num, _ := strconv.Atoi(part)
            heights = append(heights, num)
        }
        heightmap[i] = heights
    }

    for i, row := range heightmap {
        for j, col := range row {
            if col == 0 {
                result := make(map[int]bool)
                rating := 0
                walk(j, i, 0, result, &rating, heightmap)
                sum += len(result)
                ratingSum += rating
            }
        }
    }

    fmt.Printf("the sum of scores: %d\n", sum)
    fmt.Printf("the sum of ratings: %d\n", ratingSum)
}

func walk(x, y, expectedHeight int, result map[int]bool, rating *int, heightmap [][]int) {
    if x < 0 || x >= len(heightmap[0]) || y < 0 || y >= len(heightmap) {
        return
    }

    if heightmap[y][x] != expectedHeight {
        return
    }

    if expectedHeight == 9 {
        result[x * 1000 + y] = true
        *rating++
        return
    }

    for _, dir := range directions {
        walk(x + dir[0], y + dir[1], expectedHeight + 1, result, rating, heightmap)
    }
}
