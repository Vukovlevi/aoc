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
		return
	}

	str := string(data)
	lines := strings.Split(str, "\r\n")

	strTimes := strings.Split(lines[0], " ")[1:]
	strRecords := strings.Split(lines[1], " ")[1:]
	var times []int
	var records []int
	sum := 1

	for _, time := range strTimes {
		if time != "" {
			num, err := strconv.Atoi(time)
			if err != nil {
				fmt.Printf("Error converting time: %s", err.Error())
				return
			}

			times = append(times, num)
		}
	}

	for _, record := range strRecords {
		if record != "" {
			num, err := strconv.Atoi(record)
			if err != nil {
				fmt.Printf("Error converting record: %s", err.Error())
				return
			}

			records = append(records, num)
		}
	}

	for i, time := range times {
		rec := records[i]
		beats := 0
		for j := 1; j < time; j++ {
			distance := j * (time - j)
			if distance > rec {
				beats++
			}
		}
		sum *= beats
	}

	fmt.Printf("You can beat the records in: %d ways.", sum)
}