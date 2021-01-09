package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalPaper := 0
	totalRibbon := 0

	for scanner.Scan() {
		dims := strings.Split(scanner.Text(), "x")
		length, err := strconv.Atoi(dims[0])
		if err != nil {
			panic(err)
		}
		width, err := strconv.Atoi(dims[1])
		if err != nil {
			panic(err)
		}
		height, err := strconv.Atoi(dims[2])
		if err != nil {
			panic(err)
		}
		extraPaper := 99999
		ribbon := 0
		lwArea := length * width
		if lwArea < extraPaper {
			extraPaper = lwArea
			ribbon = (2 * length) + (2 * width)
		}
		lhArea := length * height
		if lhArea < extraPaper {
			extraPaper = lhArea
			ribbon = (2 * length) + (2 * height)
		}
		whArea := width * height
		if whArea < extraPaper {
			extraPaper = whArea
			ribbon = (2 * width) + (2 * height)
		}
		totalPaper += (2 * lwArea) + (2 * lhArea) + (2 * whArea) + extraPaper
		totalRibbon += ribbon + (length * width * height)
	}
	fmt.Printf("Total amount of wrapping paper: %d sf\n", totalPaper)
	fmt.Printf("Total amount of ribbon: %d f\n", totalRibbon)
}
