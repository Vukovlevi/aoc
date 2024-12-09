package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    data, _ := os.ReadFile("input.txt")
    str := string(data)
    str = str[:len(str) - 1]
    diskmap := strings.Split(str, "")
    disk := make([]string, 0)
    disk2 := make([]string, 0)
    sum := 0
    sum2 := 0

    fileid := 0
    for i, space := range diskmap {
        num, _ := strconv.Atoi(space)
        char := "."
        if i % 2 == 0 {
            char = strconv.Itoa(fileid)
            fileid++
        }

        for j := 0; j < num; j++ {
            disk = append(disk, char)
            disk2 = append(disk2, char)
        }
    }

    iterate := true
    for i := len(disk) - 1; i >= 0 && iterate; i-- {
        index := findFreeSpaceIndex(disk)
        if index >= i {
            iterate = false
            continue
        }

        disk[index] = disk[i]
        disk[i] = "."
    }

    for i := len(disk2) - 1; i >= 0; i-- {
        file, _ := strconv.Atoi(disk2[i])
        if disk2[i] != "." && file < fileid {
            fileid--
            fileSize := 0
            for j := i; j >= 0 && disk2[j] == strconv.Itoa(fileid); j-- {
                fileSize++
                i = j
            }

            index := findFreeSpaceIndexWithSize(disk2, fileSize)
            if index == -1 || index > i {
                continue
            }

            for j := index; j < index + fileSize; j++ {
                disk2[j] = strconv.Itoa(fileid)
            }

            for j := i; j < i + fileSize; j++ {
                disk2[j] = "."
            }
        }
    }

    for i, char := range disk {
        if char == "." {
            break
        }
        num, _ := strconv.Atoi(char)
        sum += i * num
    }

    for i, char := range disk2 {
        if char != "." {
            num, _ := strconv.Atoi(char)
            sum2 += i * num
        }
    }

    fmt.Printf("the checksum: %d\n", sum)
    fmt.Printf("the checksum with whole files moved: %d\n", sum2)
}

func findFreeSpaceIndex(disk []string) int {
    for i, _ := range disk {
        if disk[i] == "." {
            return i
        }
    }

    return -1
}

func findFreeSpaceIndexWithSize(disk []string, size int) int {
    for i := 0; i < len(disk); i++ {
        if disk[i] == "." {
            found := true
            for j := i; j < i + size; j++ {
                if j >= len(disk) || disk[j] != "." {
                    i = j
                    found = false
                    break
                }
            }

            if found {
                return i
            }
        }
    }

    return -1
}
