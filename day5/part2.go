package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func mapValue(collection [][]int, current int) int {
	i := 0
	found := false
	for i < len(collection) && !found {
		currentCollection := collection[i]
		if currentCollection[1] > current {
			i++
			continue
		}

		if currentCollection[1] + currentCollection[2] - 1 < current {
			i++
			continue
		}

		dif := current - currentCollection[1]

		current = currentCollection[0] + dif
		found = true
	}

	return current
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")
	var seeds []int
	var seedToSoil [][]int
	var soilToFert [][]int
	var fertToWater [][]int
	var waterToLight [][]int
	var lightToTemp [][]int
	var tempToHum [][]int
	var humToLoc [][]int
	appendIndexes := []*[][]int{&seedToSoil, &soilToFert, &fertToWater, &waterToLight, &lightToTemp, &tempToHum, &humToLoc}
	currentAppendIndex := 0
	appendTo := appendIndexes[currentAppendIndex]

	strSeeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	for _, seed := range strSeeds {
		intSeed, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Printf("Error converting seed to int: %s", err.Error())
			return
		}

		seeds = append(seeds, intSeed)
	}

	for i := 3; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			currentAppendIndex++
			appendTo = appendIndexes[currentAppendIndex]
			i++

			continue
		}

		j := 0
		strNum := ""
		var numbers []int
		for j < len(line) {
			for j < len(line) && unicode.IsDigit(rune(line[j])) {
				strNum += string(line[j])
				j++
			}

			intNum, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error converting to int: %s", err.Error())
				return
			}

			numbers = append(numbers, intNum)
			strNum = ""
			j++
		}

		*appendTo = append(*appendTo, numbers)
	}

	currentSeed := 0
	currentSoil := 0
	currentFert := 0
	currentWater := 0
	currentLight := 0
	currentTemp := 0
	currentHum := 0
	currentLocation := 0
	var locations []int

	for i := 0; i < len(seeds) - 1; i++ {
		for seed := seeds[i]; seed < seeds[i] + seeds[i + 1]; seed++ {
			currentSeed = seed
			currentSoil = mapValue(seedToSoil, currentSeed)
			currentFert = mapValue(soilToFert, currentSoil)
			currentWater = mapValue(fertToWater, currentFert)
			currentLight = mapValue(waterToLight, currentWater)
			currentTemp = mapValue(lightToTemp, currentLight)
			currentHum = mapValue(tempToHum, currentTemp)
			currentLocation = mapValue(humToLoc, currentHum)

			locations = append(locations, currentLocation)
		}
		if i + 1 != len(seeds) - 1 {
			i++
		}
	}

	min := locations[0]
	for i := 1; i < len(locations); i++ {
		if locations[i] < min {
			min = locations[i]
		}
	}

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("Error creating output file: %s", err.Error())
		return
	}

	_, err = file.WriteString(fmt.Sprintf("The lowest location number: %d", min))
	if err != nil {
		fmt.Printf("Error writing output file: %s", err.Error())
		return
	}

	file.Close()
}