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
				galaxyNumber++
			}
		}

		if galaxyNumber == 0 {
			emptyCols = append(emptyCols, i)
		}
	}

	for i := 0; i < len(emptyRows); i++ {
		after := lines[emptyRows[i] + i:]
		lines = lines[:emptyRows[i] + i]

		row := ""
		for j := 0; j < len(lines[0]); j++ {
			row += "."
		}

		lines = append(lines, row)
		lines = append(lines, after...)
	}

	for i := 0; i < len(emptyCols); i++ {
		for j := 0; j < len(lines); j++ {
			lines[j] = lines[j][:emptyCols[i] + i] + "." + lines[j][emptyCols[i] + i:]
		}
	}

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "#" {
				galaxies = append(galaxies, Point{x: j, y: i})
			}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += int(math.Abs(float64(galaxies[i].x - galaxies[j].x)) + math.Abs(float64(galaxies[i].y - galaxies[j].y)))
		}
	}

	fmt.Printf("The sum of the distances between galaxies: %d", sum)
}