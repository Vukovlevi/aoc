package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Printf("Error reading input file: %s", err.Error())
		return
    }

	str := string(data)
	str = strings.Trim(str, "\r\n") + ","

	sum := 0
	hashCode := 0

	for _, char := range str {
		if string(char) == "," {
			sum += hashCode
            hashCode = 0
			continue
		}

		hashCode += int(char)
		hashCode *= 17
		hashCode %= 256
	}

	fmt.Printf("The sum of the hashes: %d", sum)
}