package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")

	directions := lines[0]
	direction := 0
	key := "AAA"

	jumps := 0
	coords := make(map[string][]string)

	for i := 2; i < len(lines); i++ {
		splitted := strings.Split(lines[i], " = ")
		coord := strings.ReplaceAll(splitted[1], "(", "")
		coord = strings.ReplaceAll(coord, ")", "")
		coord = strings.ReplaceAll(coord, ",", "")

		coords[splitted[0]] = strings.Split(coord, " ")
	}

	for key != "ZZZ" {
		if direction == len(directions) {
			direction = 0
		}

		currentDirection := directions[direction]
		nextKey := ""

		switch string(currentDirection) {
		case "L":
			nextKey = coords[key][0]
			break
		case "R":
			nextKey = coords[key][1]
			break
		}

		direction++
		jumps++

		key = nextKey
	}

	fmt.Printf("The number of jumps: %d", jumps)
}