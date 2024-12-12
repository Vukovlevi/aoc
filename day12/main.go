package main

import (
	"fmt"
	"os"
	"strings"
)

type plot struct {
    pos int
    perimeter int
}

var (
    processedMap = make(map[int]bool)
    directions = [][]int {
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
    gardenMap := make([][]string, len(lines))
    sum := 0

    for i, line := range lines {
        gardenMap[i] = strings.Split(line, "")
    }

    for i, row := range gardenMap {
        for j, _ := range row {
            if _, ok := processedMap[j * 1000 + i]; ok {
                continue
            }

            perimeter := 0
            area := 0
            getArea(j, i, &perimeter, &area, gardenMap)
            sum += perimeter * area
        }
    }

    fmt.Printf("the cost of the fence: %d\n", sum)
}

func getArea(x, y int, perimeter, area *int, gardenMap [][]string) {
    if _, ok := processedMap[x * 1000 + y]; ok {
        return
    }

    processedMap[x * 1000 + y] = true
    *area++
    for _, dir := range directions {
        if isOutOfBounds(x + dir[0], y + dir[1], gardenMap) {
            *perimeter++
            continue
        }

        if gardenMap[y][x] != gardenMap[y + dir[1]][x + dir[0]] {
            *perimeter++
            continue
        }

        getArea(x + dir[0], y + dir[1], perimeter, area, gardenMap)
    }
}

func isOutOfBounds(x, y int, gardenMap [][]string) bool {
    return x < 0 || x >= len(gardenMap[0]) || y < 0 || y >= len(gardenMap)
}
