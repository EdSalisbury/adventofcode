package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nice := 0

	for scanner.Scan() {
		word := scanner.Text()
		pairIndex := make(map[string]int)
		hasPair := false
		for i := 0; i < len(word)-1; i++ {
			pair := word[i : i+2]
			_, pairExists := pairIndex[pair]
			if !pairExists {
				pairIndex[pair] = i
			} else {
				if i-pairIndex[pair] > 1 {
					hasPair = true
					continue
				}
			}
		}

		hasChars := false
		for i := 2; i < len(word); i++ {
			if word[i] == word[i-2] {
				hasChars = true
				continue
			}
		}

		if hasPair && hasChars {
			nice++
		}

	}
	fmt.Printf("Number of nice strings: %d\n", nice)
}
