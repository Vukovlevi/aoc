package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")

	sum := 0
	minSum := 0

	for _, line := range lines {
		gameData := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Split(gameData[0], " ")[1])
		if err != nil {
			fmt.Printf("Error converting id: %s", err.Error())
			return
		}

		reveals := strings.Split(strings.TrimPrefix(gameData[1], " "), ";")
		skip := false
		maxGreen := 0
		maxRed := 0
		maxBlue := 0
		for _, reveal := range reveals {
			oneBalls := strings.Split(reveal, ",")
			for _, oneBall := range oneBalls {
				stat := strings.Split(strings.TrimPrefix(oneBall, " "), " ")
				number, err := strconv.Atoi(stat[0])
				if err != nil {
					fmt.Printf("Error converting ball number: %s", err.Error())
					return
				}
				color := stat[1]

				switch (color) {
					case "red":
						if number > maxRed {
							maxRed = number
						}
						if number > 12 {
							skip = true
						}
						break
					case "green": 
						if number > maxGreen {
							maxGreen = number
						}
						if number > 13 {
							skip = true
						}
						break
					case "blue":
						if number > maxBlue {
							maxBlue = number
						}
						if number > 14 {
							skip = true
						}
						break
				}
			}
		}

		if !skip {
			sum += id
		}

		minSum += maxRed * maxGreen * maxBlue
	}

	fmt.Printf("The sum of the ids of possible games: %d\n", sum)
	fmt.Printf("The sum of the minimum balls required to play all rounds: %d", minSum)
}