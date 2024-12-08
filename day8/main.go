package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
    tower = make(map[int]string)
    antinodes = make(map[int]bool)
    antinodes2 = make(map[int]bool)
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    lines := strings.Split(str, "\n")
    lines = lines[:len(lines) - 1]
    nodemap := make([][]string, len(lines))
    for i, _ := range lines {
        nodemap[i] = strings.Split(lines[i], "")
    }

    for i, row := range nodemap {
        for j, col := range row {
            if col != "." {
                tower[i * 1000 + j] = col
            }
        }
    }

    for k1, v1 := range tower {
        for k2, v2 := range tower {
            if v1 == v2 && k1 != k2 {
                antinodeCount(k1, k2, len(nodemap), len(nodemap[0]))
                antinodeCount2(k1, k2, len(nodemap), len(nodemap[0]))
            }
        }
    }

    fmt.Printf("the count of antinode spots: %d\n", len(antinodes))
    fmt.Printf("the count of antinode spots in version 2: %d\n", len(antinodes2))
}

func antinodeCount(pos1, pos2, boundRow, boundCol int) {
    y1 := pos1 / 1000
    x1 := pos1 % 1000

    y2 := pos2 / 1000
    x2 := pos2 % 1000

    mRow := y2 - y1
    mCol := x2 - x1

    currCol := x1 + mCol
    currRow := y1 + mRow

    found := 0
    for currCol >= 0 && currCol < boundCol && currRow >= 0 && currRow < boundRow && found < 2 {
        dis1 := math.Abs(float64(currCol - x1)) + math.Abs(float64(currRow - y1))
        dis2 := math.Abs(float64(currCol - x2)) + math.Abs(float64(currRow - y2))

        if dis1 == dis2 * 2 || dis2 == dis1 * 2 {
            found++
            antinodes[currRow * 1000 + currCol] = true
        }

        currCol += mCol
        currRow += mRow
    }

    currCol = x1 - mCol
    currRow = y1 - mRow

    for currCol >= 0 && currCol < boundCol && currRow >= 0 && currRow < boundRow && found < 2 {
        dis1 := math.Abs(float64(currCol - x1)) + math.Abs(float64(currRow - y1))
        dis2 := math.Abs(float64(currCol - x2)) + math.Abs(float64(currRow - y2))

        if dis1 == dis2 * 2 || dis2 == dis1 * 2 {
            found++
            antinodes[currRow * 1000 + currCol] = true
        }

        currCol -= mCol
        currRow -= mRow
    }
}

func antinodeCount2(pos1, pos2, boundRow, boundCol int) {
    y1 := pos1 / 1000
    x1 := pos1 % 1000

    y2 := pos2 / 1000
    x2 := pos2 % 1000

    mRow := y2 - y1
    mCol := x2 - x1

    currCol := x1
    currRow := y1

    for currCol >= 0 && currCol < boundCol && currRow >= 0 && currRow < boundRow {
        antinodes2[currRow * 1000 + currCol] = true
        currCol += mCol
        currRow += mRow
    }

    currCol = x1 - mCol
    currRow = y1 - mRow

    for currCol >= 0 && currCol < boundCol && currRow >= 0 && currRow < boundRow {
        antinodes2[currRow * 1000 + currCol] = true
        currCol -= mCol
        currRow -= mRow
    }
}
