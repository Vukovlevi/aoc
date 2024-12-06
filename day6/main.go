package main

import (
	"fmt"
	"os"
)

var (
    visited = make(map[int]bool)
    guardMap = make([][]byte, 0)
    dir = '^'

    startX = 0
    startY = 0
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)

    i := 0
    for i < len(str) {
        row := make([]byte, 0)
        for str[i] != '\n' {
            row = append(row, str[i])

            if str[i] == '^' {
                actualI := i - len(guardMap)
                startY = actualI / len(guardMap[0])
                startX = actualI % len(guardMap[0])
            }

            i++
        }
        guardMap = append(guardMap, row)
        i++
    }

    going := true
    for going {
        visited[startX * 1000 + startY] = true
        switch dir {
        case '^':
            if IsOff(startY - 1, startX) {
                going = false
                break
            }

            if guardMap[startY - 1][startX] == '#' {
                dir = '>'
                break
            }

            startY--
        case '>':
            if IsOff(startY, startX + 1) {
                going = false
                break
            }

            if guardMap[startY][startX + 1] == '#' {
                dir = 'v'
                break
            }

            startX++
        case '<':
            if IsOff(startY, startX - 1) {
                going = false
                break
            }

            if guardMap[startY][startX - 1] == '#' {
                dir = '^'
                break
            }

            startX--
        case 'v':
            if IsOff(startY + 1, startX) {
                going = false
                break
            }

            if guardMap[startY + 1][startX] == '#' {
                dir = '<'
                break
            }

            startY++
        }
    }

    fmt.Printf("visited spots: %d\n", len(visited))
}

func IsOff(x, y int) bool {
    return x < 0 || x >= len(guardMap[0]) || y < 0 || y >= len(guardMap)
}
