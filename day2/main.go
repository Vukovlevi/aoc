package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)

    lines := strings.Split(str, "\n")
    safeReports := 0

    for _, line := range lines[:len(lines) - 1] {
        report := make([]int, 0)
        parts := strings.Split(line, " ")

        for _, part := range parts {
            num, _ := strconv.Atoi(part)
            report = append(report, num)
        }

        if isSafe(report) {
            safeReports++
            continue
        }

        safe := false

        for i, _ := range report {
            sliced := make([]int, 0)
            sliced = append(sliced, report[:i]...)
            sliced = append(sliced, report[i+1:]...)
            if isSafe(sliced) {
                safe = true
                break
            }
        }

        if safe {
            safeReports++
        }
    }

    fmt.Printf("number of safe reports: %d\n", safeReports)
}

func isSafe(report []int) bool {
    increasing := true

    for i := 0; i < len(report) - 1; i++ {
        dif := report[i] - report[i + 1]

        if dif == 0 || math.Abs(float64(dif)) > 3 {
            return false
        }

        if i == 0 {
            if dif < 0 {
                increasing = false
            }
        } else {
            if increasing && dif < 0 {
                return false
            } else if !increasing && dif > 0 {
                return false
            }
        }
    }

    return true
}
