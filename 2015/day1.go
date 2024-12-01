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
	basementEntered := false

	for index, char := range content {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor == -1 && !basementEntered {
			fmt.Printf("Entered basement at position: %d\n", index+1)
			basementEntered = true
		}
	}
	fmt.Printf("floor: %d\n", floor)
}
