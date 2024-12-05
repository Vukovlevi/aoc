package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
    ordering = make(map[int][]int)
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    parts := strings.Split(str, "\n\n")
    rules := strings.Split(parts[0], "\n")
    pageUpdates := strings.Split(parts[1], "\n")
    pageUpdates = pageUpdates[:len(pageUpdates) - 1]
    sum := 0
    notCorrectSum := 0

    for _, rule := range rules {
        nums := strings.Split(rule, "|")
        num1, _ := strconv.Atoi(nums[0])
        num2, _ := strconv.Atoi(nums[1])

        if _, ok := ordering[num2]; ok {
            ordering[num2] = append(ordering[num2], num1)
        } else {
            ordering[num2] = []int{num1}
        }
    }

    for _, update := range pageUpdates {
        pages := strings.Split(update, ",")
        printedPages := make(map[int]bool)

        correct := true
        outer:
        for _, page := range pages {
            pageNum, _ := strconv.Atoi(page)
            if _, ok := ordering[pageNum]; ok {
                for _, before := range ordering[pageNum] {
                    if !slices.Contains(pages, strconv.Itoa(before)) {
                        continue
                    }

                    if _, ok := printedPages[before]; !ok {
                        correct = false
                        break outer
                    }
                }
            }

            printedPages[pageNum] = true
        }

        if !correct {
            notCorrectSum += sortPages(pages)
            continue
        }

        middleIndex := math.Floor(float64(len(pages) / 2))
        middle, _ := strconv.Atoi(pages[int(middleIndex)])
        sum += middle
    }

    fmt.Printf("the sum of the middle pages: %d\n", sum)
    fmt.Printf("the sum of the middle pages of the not correct ones: %d\n", notCorrectSum)
}

func sortPages(pages []string) int {
    pageNums := make([]int, len(pages))
    for i, _ := range pages {
        num, _ := strconv.Atoi(pages[i])
        pageNums[i] = num
    }

    for i := 0; i < len(pageNums); i++ {
        for j := i; j < len(pageNums); j++ {
            if slices.Contains(ordering[pageNums[i]], pageNums[j]) {
                temp := pageNums[j]
                pageNums[j] = pageNums[i]
                pageNums[i] = temp
            }
        }
    }

    middleIndex := math.Floor(float64(len(pageNums) / 2))
    return pageNums[int(middleIndex)]
}
