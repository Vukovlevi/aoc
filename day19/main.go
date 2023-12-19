package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseWork(work string, works map[string]string, rating map[string]int, sum *int) {
	switch work {
	case "A":
		*sum += rating["x"] + rating["m"] + rating["a"] + rating["s"]
		return
	case "R":
		return
	}

	current := works[work]
	eol := false

	for !eol {
		if len(current) <= 5 {
			parseWork(current[:len(current) - 1], works, rating, sum)
			return
		}

		lookUpValue := rating[string(current[0])]
		statement := current[1]
		toMatchValue := ""
		i := 2
		for string(current[i]) != ":" && string(current[i]) != "}" {
			toMatchValue += string(current[i])
    	    i++
		}
		
		if string(current[i]) == "}" {
			parseWork(current[:len(current) - 1], works, rating, sum)
			return
		}

		intNum, err := strconv.Atoi(toMatchValue)
		if err != nil {
    	    fmt.Printf("Error parsing toMatchValue: %s", err.Error())
			return
    	}

		i++
		todoWork := ""
		for string(current[i]) != "," && string(current[i]) != "}" {
			todoWork += string(current[i])
            i++
		}

		if string(current[i]) == "}" {
			eol = true
		} else {
			switch string(statement) {
			case "<":
				if lookUpValue < intNum {
					parseWork(todoWork, works, rating, sum)
					return
                }
				break
			case ">":
				if lookUpValue > intNum {
                    parseWork(todoWork, works, rating, sum)
					return
                }
                break
			}

			current = current[i + 1:]
		}
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
	}

	str := string(data)
	lines := strings.Split(str, "\n")

	workflows := make(map[string]string)
	var ratings []map[string]int

	isWorkflow := true
	sum := 0

	for _, line := range lines {
		if line == "" {
			isWorkflow = false
            continue
		}

		if isWorkflow {
			splitted := strings.Split(line, "{")
			workflows[splitted[0]] = splitted[1]
		} else {
			rating := make(map[string]int)
			line = strings.Replace(line, "{", "", 1)
			line = strings.Replace(line, "}", "", 1)
			splitted := strings.Split(line, ",")
			for _, value := range splitted {
				keyValue := strings.Split(value, "=")

				intNum, err := strconv.Atoi(keyValue[1])
				if err!= nil {
                    fmt.Printf("Error converting value to int: %s", err.Error())
                    return
                }
                rating[keyValue[0]] = intNum
			}
			ratings = append(ratings, rating)
		}
	}

	for _, rating := range ratings {
		parseWork("in", workflows, rating, &sum)
	}

	fmt.Printf("The sum of the accepted ratings: %d", sum)
}