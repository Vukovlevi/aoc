package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")
	
	var emptyRows []int
	var emptyCols []int
	var galaxies []Point
	sum := 0

	for i, line := range lines {
		galaxyNumber := 0
		for _, char := range line {
			if string(char) == "#" {
				galaxyNumber++
			}
		}

		if galaxyNumber == 0 {
			emptyRows = append(emptyRows, i)
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		galaxyNumber := 0
		for j := 0; j < len(lines); j++ {
			if string(lines[j][i]) == "#" {
				galaxies = append(galaxies, Point{x: i, y: j})
				galaxyNumber++
			}
		}

		if galaxyNumber == 0 {
			emptyCols = append(emptyCols, i)
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			difX := int(math.Abs(float64(galaxies[i].x - galaxies[j].x)))
			difY := int(math.Abs(float64(galaxies[i].y - galaxies[j].y)))

			for _, col := range emptyCols {
				if (col > galaxies[i].x && col < galaxies[j].x) || (col < galaxies[i].x && col > galaxies[j].x) {
					difX += 999999
				}
			}

			for _, row := range emptyRows {
				if (row > galaxies[i].y && row < galaxies[j].y) || (row < galaxies[i].y && row > galaxies[j].y)  {
					difY += 999999
				}
			}

			sum += difX + difY
		}
	}

	fmt.Printf("The sum of the distances between galaxies: %d", sum)
}