package main

import (
	"fmt"
	"os"
	"strings"
)

func gcd(a, b int) int {
	if a == b {
		return a
	}

	if a < b {
		s := a
		a = b
		b = s
	}

	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int, numbers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

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
	var endsWithA []string
	
	jump := 0
	var jumps []int
	coords := make(map[string][]string)
	
	for i := 2; i < len(lines); i++ {
		splitted := strings.Split(lines[i], " = ")
		coord := strings.ReplaceAll(splitted[1], "(", "")
		coord = strings.ReplaceAll(coord, ")", "")
		coord = strings.ReplaceAll(coord, ",", "")
		
		coords[splitted[0]] = strings.Split(coord, " ")
		
		if string(splitted[0][len(splitted[0]) - 1]) == "A" {
			endsWithA = append(endsWithA, splitted[0])
		}
	}
	
	for i := 0; i < len(endsWithA); i++ {
		key := endsWithA[i]

		for string(key[len(key) - 1]) != "Z" {
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
			jump++

			key = nextKey
		}

		jumps = append(jumps, jump)
		jump = 0
	}

	jump = lcm(jumps[0], jumps[1], jumps[2:]...)

	fmt.Printf("The number of jumps: %v", jump)
}