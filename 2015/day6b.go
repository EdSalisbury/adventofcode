package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type coord struct {
	x int
	y int
}

func main() {
	file, err := os.Open("day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Looking for one of the following patterns:
	// turn on 792,764 through 872,842
	// turn off 664,203 through 694,754
	// toggle 491,615 through 998,836
	pattern := regexp.MustCompile(`^([\w\s]+?)\s+(\d+),(\d+)\s+through\s+(\d+),(\d+)$`)
	scanner := bufio.NewScanner(file)

	var lights [1000][1000]int

	for scanner.Scan() {
		line := scanner.Text()
		fields := pattern.FindStringSubmatch(line)
		command := fields[1]
		x, _ := strconv.Atoi(fields[2])
		y, _ := strconv.Atoi(fields[3])
		start := coord{x, y}
		x, _ = strconv.Atoi(fields[4])
		y, _ = strconv.Atoi(fields[5])
		end := coord{x, y}

		for y = start.y; y <= end.y; y++ {
			for x = start.x; x <= end.x; x++ {
				if command == "turn on" {
					lights[x][y]++
				} else if command == "turn off" {
					lights[x][y]--
					if lights[x][y] < 0 {
						lights[x][y] = 0
					}
				} else if command == "toggle" {
					lights[x][y] += 2
				}
			}
		}
	}

	brightness := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			brightness += lights[x][y]
		}
	}

	fmt.Printf("The total amount of brightness: %d\n", brightness)
}
