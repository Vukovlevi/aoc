package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("test.txt")
	if err!= nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
    }

	str := string(data)
	str = strings.Trim(str, "\r\n")

	steps := strings.Split(str, ",")

	boxes := make([]map[string]int, 256)
	for i := range boxes {
		boxes[i] = make(map[string]int)
	}
	sum := 0

	for _, step := range steps {
		hashCode := 0
		label := ""
		for i := 0; i < len(step); i++ {
			switch string(step[i]) {
			case "=":
				focal, err := strconv.Atoi(string(step[i + 1]))
				if err!= nil {
                    fmt.Printf("Error parsing focal: %s", err.Error())
                    return
                }
				boxes[hashCode][label] = focal
				i++
				break
			case "-":
				delete(boxes[hashCode], label)
				break
			default:
				for string(step[i]) != "=" && string(step[i]) != "-" {
					label += string(step[i])
					hashCode += int(step[i])
                    hashCode *= 17
                    hashCode %= 256
                    i++
				}
				i--
				break
			}
		}
	}

	for i, box := range boxes {
		slot := 1
		for key := range box {
			sum += (i + 1) * slot * box[key]
			slot++
		}
	}

	fmt.Printf("The sum of the hashes: %d", sum)
}