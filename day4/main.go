package main

import (
	"fmt"
	"os"
	"strings"
)

var (
    directrions = [][]int{
        {0, 1},
        {1, 1},
        {1, 0},
        {1, -1},
        {0, -1},
        {-1, -1},
        {-1, 0},
        {-1, 1},
    }
    letters = []byte{'X', 'M', 'A', 'S'}
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    splitted := strings.Split(str, "\n")
    characters := make([][]string, len(splitted) - 1)
    for i, row := range splitted {
        if i < len(splitted) - 1 {
            characters[i] = strings.Split(row, "")
        }
    }

    occurs := 0
    occursX := 0

    for i, row := range characters {
        for j, char := range row {
            if char == "X" {
                occurs += checkWord(characters, j, i)
            } else if char == "A" {
                occursX += checkXmas(characters, j, i)
            }
        }
    }

    fmt.Printf("the word occurs: %d times\n", occurs)
    fmt.Printf("the word occurs in x shape: %d times\n", occursX)
}

func checkWord(characters [][]string, x, y int) int {
    occurs := 0
    for _, dir := range directrions {
        possible := true
        for i := 0; i < len(letters) && possible; i++ {
            currX := x + dir[0] * i
            currY := y + dir[1] * i
            if checkInBounds(characters, currX, currY) {
                if byte(characters[currY][currX][0]) != letters[i] {
                    possible = false
                }
            } else {
                possible = false
            }
        }

        if possible {
            occurs++
        }
    }

    return occurs
}

func checkInBounds(characters [][]string, x, y int) bool {
    return x >= 0 && x < len(characters[0]) && y >= 0 && y < len(characters)
}

func checkXmas(characters [][]string, x, y int) int {
    if !checkInBounds(characters, x - 1, y - 1) || !checkInBounds(characters, x + 1, y + 1) {
        return 0
    }

    diagonal := false
    if characters[y - 1][x - 1] == "M" && characters[y + 1][x + 1] == "S" {
        diagonal = true
    } else if characters[y - 1][x - 1] == "S" && characters[y + 1][x + 1] == "M" {
        diagonal = true
    }

    if !diagonal {
        return 0
    }

    if !checkInBounds(characters, x - 1, y + 1) || !checkInBounds(characters, x + 1, y - 1) {
        return 0
    }

    diagonal = false
    if characters[y + 1][x - 1] == "M" && characters[y - 1][x + 1] == "S" {
        diagonal = true
    } else if characters[y + 1][x - 1] == "S" && characters[y - 1][x + 1] == "M" {
        diagonal = true
    }

    if !diagonal {
        return 0
    }

    return 1
}
