package main

import (
	"fmt"
	"os"
	"strings"
)

func copy(arr []byte) []byte {
	newArr := arr
	return newArr
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")
	sum := 0

	var characters [][]byte

	for _, line := range lines {
		var characterLine []byte
		for _, char := range line {
			characterLine = append(characterLine, byte(char))
		}

		characters = append(characters, characterLine)
	}

	for i := 0; i < len(characters[0]); i++ {
		for j := 0; j < len(characters); j++ {
			if j == 0 && string(characters[j][i]) == "O" {
				sum += len(characters)
				continue
			}

			startingJ := j
			for j > 0 && string(characters[j][i]) == "O" && string(characters[j - 1][i]) == "." {
				characters[j][i] = byte('.')
				characters[j - 1][i] = byte('O')
				j--
			}

			if string(characters[j][i]) == "O" {
				sum += len(characters) - j
				j = startingJ
			}
		}
	}

	fmt.Printf("The sum of the loads: %d", sum)
}