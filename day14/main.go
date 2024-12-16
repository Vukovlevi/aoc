package main

import (
	"fmt"
	"os"
    "math"
	"strconv"
	"strings"
)

const (
    width = 101
    heigth = 103
)

type Robot struct {
    x int
    y int
    xVel int
    yVel int
}

func (r *Robot) Simulate(seconds int) {
    for i := 0; i < seconds; i++ {
        r.x += r.xVel
        r.y += r.yVel

        if r.x < 0 {
            r.x = int(width - math.Abs(float64(r.x)))
        } else if r.x >= width {
            r.x = r.x % width
        }

        if r.y < 0 {
            r.y = int(heigth - math.Abs(float64(r.y)))
        } else if r.y >= heigth {
            r.y = r.y % heigth
        }
    }
}

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    lines := strings.Split(str, "\n")
    lines = lines[:len(lines) - 1]
    robots := make([]Robot, len(lines))
    quadrants := make([]int, 4)
    sum := 1

    for i, line := range lines {
        firstCommaIndex := strings.Index(line, ",")
        firstSpaceIndex := strings.Index(line, " ")
        lastCommaIndex := strings.LastIndex(line, ",")
        x, _ := strconv.Atoi(line[2:firstCommaIndex])
        y, _ := strconv.Atoi(line[firstCommaIndex + 1:firstSpaceIndex])
        xVel, _ := strconv.Atoi(line[firstSpaceIndex + 3: lastCommaIndex])
        yVel, _ := strconv.Atoi(line[lastCommaIndex + 1:])
        robots[i] = Robot{
            x: x,
            y: y,
            xVel: xVel,
            yVel: yVel,
        }
        robots[i].Simulate(100)
    }

    middleRow := heigth / 2
    middleCol := width / 2

    for _, robot := range robots {
        if robot.x < middleCol {
            if robot.y < middleRow {
                quadrants[0] += 1
            } else if robot.y > middleRow {
                quadrants[1] += 1
            }
        } else if robot.x > middleCol {
            if robot.y < middleRow {
                quadrants[2] += 1
            } else if robot.y > middleRow {
                quadrants[3] += 1
            }
        }
    }

    for _, quadrant := range quadrants {
        sum *= quadrant
    }

    fmt.Printf("the safety factor: %d\n", sum)
}
