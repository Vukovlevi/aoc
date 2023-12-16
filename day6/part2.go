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
	time := 0
	record := 0
	sum := 0

	strTime := strings.Join(strTimes, "")
	num, err := strconv.Atoi(strTime)
	if err != nil {
		fmt.Printf("Error converting time: %s", err.Error())
		return
	}
	time = num

	strRecord := strings.Join(strRecords, "")
	num, err = strconv.Atoi(strRecord)
	if err != nil {
		fmt.Printf("Error converting time: %s", err.Error())
		return
	}
	record = num
	
	for i := 1; i < time; i++ {
		distance := i * (time - i)
		if distance > record {
			sum++
		}
	}

	fmt.Printf("You can beat the records in: %d ways.", sum)
}