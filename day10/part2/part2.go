package main

import (
	"fmt"
	"os"
	"strings"
)

type Tile struct {
	direction string
	comesFrom string
	index int
}

func parseNextTile(t Tile, p map[rune]Tile, lineSize int, s *string) Tile {
	var nextTile Tile
	str := *s

	switch strings.Replace(t.direction, t.comesFrom, "", 1) {
	case "north":
		nextTile = p[rune(str[t.index - lineSize])]
		nextTile.comesFrom = "south"
		nextTile.index = t.index - lineSize
		break
	case "west":
		nextTile = p[rune(str[t.index - 1])]
		nextTile.comesFrom = "east"
		nextTile.index = t.index - 1
		break
	case "south":
		nextTile = p[rune(str[t.index + lineSize])]
		nextTile.comesFrom = "north"
		nextTile.index = t.index + lineSize
		break
	case "east":
		nextTile = p[rune(str[t.index + 1])]
		nextTile.comesFrom = "west"
		nextTile.index = t.index + 1
		break
	}

	return nextTile
}

func includes(array []int, value int) bool {
	found := false
	i := 0
	for !found && i < len(array) {
		if array[i] == value {
			found = true
		}

		i++
	}

	return found
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")
	lineSize := len(lines[0])
	str = strings.ReplaceAll(str, "\r\n", "")

	pipes := map[rune]Tile{
		'|': {direction: "northsouth", comesFrom: "", index: -1},
		'-': {direction: "westeast", comesFrom: "", index: -1},
		'L': {direction: "northeast", comesFrom: "", index: -1},
		'J': {direction: "northwest", comesFrom: "", index: -1},
		'7': {direction: "southwest", comesFrom: "", index: -1},
		'F': {direction: "southeast", comesFrom: "", index: -1},
	}

	startIndex := strings.Index(str, "S")
	waysIndexes := []int{startIndex}
	var ways [][]Tile
	if strings.Index(pipes[rune(str[startIndex - lineSize])].direction, "south") != -1 {
		var way []Tile
		tile := pipes[rune(str[startIndex - lineSize])]
		tile.comesFrom = "south"
		tile.index = startIndex - lineSize
		way = append(way, tile)

		ways = append(ways, way)
	}

	if strings.Index(pipes[rune(str[startIndex - 1])].direction, "east") != -1 {
		var way []Tile
		tile := pipes[rune(str[startIndex - 1])]
		tile.comesFrom = "east"
		tile.index = startIndex - 1
		way = append(way, tile)

		ways = append(ways, way)
	}

	if strings.Index(pipes[rune(str[startIndex + lineSize])].direction, "north") != -1 {
		var way []Tile
		tile := pipes[rune(str[startIndex + lineSize])]
		tile.comesFrom = "north"
		tile.index = startIndex + lineSize
		way = append(way, tile)

		ways = append(ways, way)
	}

	if strings.Index(pipes[rune(str[startIndex + 1])].direction, "west") != -1 {
		var way []Tile
		tile := pipes[rune(str[startIndex + 1])]
		tile.comesFrom = "west"
		tile.index = startIndex + 1
		way = append(way, tile)

		ways = append(ways, way)
	}

	i := 0
	for ways[0][len(ways[0]) - 1].index != ways[1][len(ways[1]) - 1].index {
		tile0 := parseNextTile(ways[0][i], pipes, lineSize, &str)
		tile1 := parseNextTile(ways[1][i], pipes, lineSize, &str)
		ways[0] = append(ways[0], tile0)
		ways[1] = append(ways[1], tile1)

		waysIndexes = append(waysIndexes, tile0.index)
		waysIndexes = append(waysIndexes, tile1.index)

		i++
	}

	enclosed := 0

	for j, line := range lines {
		verticals := 0
		for k, char := range line {
			if !includes(waysIndexes, j * lineSize + k) && string(char) != "S" {
				if verticals % 2 == 1 {
					enclosed++
				}
				continue
			}

			if string(char) == "|" || string(char) == "L" || string(char) == "J" {
				verticals++
			}
		}
	}

	enclosed--
	fmt.Printf("The number of enclosed tiles: %d", enclosed)
}