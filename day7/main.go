package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func sortType(array []string) {
	points := map[string]int{"A": 4, "K": 3, "Q": 2, "J": 1, "T": 0}

	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - 1 - i; j++ {
			k := 0
			for k < len(array[j]) && array[j][k] == array[j + 1][k] {
				k++
			}

			if unicode.IsDigit(rune(array[j][k])) && unicode.IsDigit(rune(array[j + 1][k])) {
				if array[j][k] < array[j + 1][k] {
					s := array[j]
					array[j] = array[j + 1]
					array[j + 1] = s
				}
			} else if !unicode.IsDigit(rune(array[j][k])) && !unicode.IsDigit(rune(array[j + 1][k])) {
				if points[string(array[j][k])] < points[string(array[j + 1][k])] {
					s := array[j]
					array[j] = array[j + 1]
					array[j + 1] = s
				}
			} else if unicode.IsDigit(rune(array[j][k])) && !unicode.IsDigit(rune(array[j + 1][k])) {
				s := array[j]
				array[j] = array[j + 1]
				array[j + 1] = s
			}

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
	lines := strings.Split(str, "\r\n")
	handToBid := make(map[string]int)
	sum := 0

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		hand, strBid := splitted[0], splitted[1]

		bid, err := strconv.Atoi(strBid)
		if err != nil {
			fmt.Printf("Error converting bid to int: %s", err.Error())
			return
		}

		handToBid[hand] = bid
	}

	var fives []string
	var fours []string
	var fullHouses []string
	var threes []string
	var twos []string
	var ones []string
	var differents []string

	for key := range handToBid {
		amount := make(map[rune]int)		
		for _, char := range key {
			if amount[char] == 0 {
				amount[char] = 1
			} else {
				amount[char]++
			}
		}

		switch len(amount) {
		case 1:
			fives = append(fives, key)
			break
		case 2:
			for k := range amount {
				if amount[k] == 4 || amount[k] == 1 {
					fours = append(fours, key)
				} else {
					fullHouses = append(fullHouses, key)
				}

				break
			}
			break
		case 3:
			for k := range amount {
				if amount[k] == 1 {
					continue
				}

				if amount[k] == 3 {
					threes = append(threes, key)
				} else {
					twos = append(twos, key)
				}
				
				break
			}
			break
		case 4:
			ones = append(ones, key)
			break
		case 5:
			differents = append(differents, key)
			break
		}
	}

	sortType(fives)
	sortType(fours)
	sortType(fullHouses)
	sortType(threes)
	sortType(twos)
	sortType(ones)
	sortType(differents)

	rank := len(lines)
	for _, hand := range fives {
		sum += handToBid[hand] * rank
		rank--
	}
	for _, hand := range fours {
		sum += handToBid[hand] * rank
		rank--
	}
	for _, hand := range fullHouses {
		sum += handToBid[hand] * rank
		rank--
	}
	for _, hand := range threes {
		sum += handToBid[hand] * rank
		rank--
	}
	for _, hand := range twos {
		sum += handToBid[hand] * rank
		rank--
	}
	for _, hand := range ones {
		sum += handToBid[hand] * rank
		rank--
	}
	for _, hand := range differents {
		sum += handToBid[hand] * rank
		rank--
	}

	fmt.Printf("The sum is: %d", sum)
}