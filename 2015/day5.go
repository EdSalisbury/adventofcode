package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	vowels := []string{"a", "e", "i", "o", "u"}
	badStrings := []string{"ab", "cd", "pq", "xy"}
	nice := 0

OUTER:
	for scanner.Scan() {
		word := scanner.Text()
		if FindCount(vowels, word) < 3 {
			continue
		}
		if FindCount(badStrings, word) > 0 {
			continue
		}

		var prevChar rune
		for _, char := range word {
			if char == prevChar {
				nice++
				continue OUTER
			}
			prevChar = char
		}
	}
	fmt.Printf("Number of nice strings: %d\n", nice)
}

func FindCount(items []string, val string) int {
	count := 0
	for _, item := range items {
		count += strings.Count(val, item)
	}
	return count
}
