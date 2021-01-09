package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}

	var floor int = 0

	for _, char := range content {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	fmt.Printf("floor: %d\n", floor)
}
