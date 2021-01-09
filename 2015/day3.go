package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}

	type loc struct {
		x, y int
	}

	houses := make(map[loc]int)

	sX := 0
	sY := 0
	rsX := 0
	rsY := 0
	isRoboSanta := false
	houses[loc{0, 0}] += 2

	for _, char := range content {
		x := 0
		y := 0
		switch char {
		case '^':
			y = 1
		case 'v':
			y = -1
		case '>':
			x = 1
		case '<':
			x = -1
		}
		if isRoboSanta {
			rsX += x
			rsY += y
			houses[loc{rsX, rsY}]++
		} else {
			sX += x
			sY += y
			houses[loc{sX, sY}]++
		}
		isRoboSanta = !isRoboSanta
	}
	fmt.Printf("Number of houses that got at least one present: %d\n", len(houses))
}
