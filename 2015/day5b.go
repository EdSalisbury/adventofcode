package main

import (
	"bufio"
	"fmt"
	"os"
)

type indexCount struct {
	index int
	count int
}

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nice := 0

	// OUTER:
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)
		pairs := make(map[string]indexCount)
		hasPair := false
		for i := 0; i < len(word)-1; i++ {
			pair := word[i : i+2]
			_, pairExists := pairs[pair]
			if !pairExists {
				pairs[pair] = indexCount{i, 1}
			} else {
				if i-pairs[pair].index > 1 {
					fmt.Printf("->>>>>> %s         hasPair\n", pair)
					hasPair = true
					continue
				}
			}
		}

		hasChars := false
		chars := make(map[byte]indexCount)
		for i := 0; i < len(word); i++ {
			char := word[i]
			_, charExists := chars[char]
			//fmt.Printf("Index of char %c: %d\n", char, i)
			if !charExists {
				chars[char] = indexCount{i, 1}
			} else {
				//fmt.Printf("Current Index of char %c: %d\n", char, chars[char].index)
				if i-chars[char].index == 2 {
					fmt.Printf("->>>>>> %c         hasChars\n", char)
					hasChars = true
					continue
				} else if i-chars[char].index > 2 {
					chars[char] = indexCount{i, 1}
				}
				//fmt.Printf("New Index of char %c: %d\n", char, chars[char].index)
			}
		}

		if hasPair && hasChars {
			nice++
			fmt.Printf("->>>>>>         has both: %d\n", nice)
		}

	}
	fmt.Printf("Number of nice strings: %d\n", nice)
}
