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

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lineSize := len(strings.Split(str, "\r\n")[0])
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
		ways[0] = append(ways[0], parseNextTile(ways[0][i], pipes, lineSize, &str))
		ways[1] = append(ways[1], parseNextTile(ways[1][i], pipes, lineSize, &str))
		i++
	}

	fmt.Printf("The farthest you can go: %d", len(ways[0]))
}