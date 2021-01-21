package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadWires(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Can't open file %s!", filename))
	}
	defer file.Close()
	wires := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "->")
		key := strings.TrimSpace(fields[1])
		value := strings.TrimSpace(fields[0])
		wires[key] = value
	}
	return wires
}

func isInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func stringToUint16(str string) uint16 {
	value, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(value)
}

func uint16ToString(val uint16) string {
	return strconv.Itoa(int(val))
}

func eval(wires map[string]string, key string, value string) uint16 {
	//fmt.Println(value)
	if isInt(value) {
		return stringToUint16(value)
	}
	fields := strings.Split(value, " ")
	if len(fields) == 3 {
		a := fields[0]
		op := fields[1]
		b := fields[2]
		var intA, intB uint16
		if isInt(a) {
			intA = stringToUint16(a)
		} else {
			intA = eval(wires, key, a)
		}
		if isInt(b) {
			intB = stringToUint16(b)
		} else {
			intB = eval(wires, key, b)
		}
		switch op {
		case "AND":
			result := intA & intB
			wires[key] = uint16ToString(result)
			return result
		case "OR":
			result := intA | intB
			wires[key] = uint16ToString(result)
			return result
		case "LSHIFT":
			result := intA << intB
			wires[key] = uint16ToString(result)
			return result
		case "RSHIFT":
			result := intA >> intB
			wires[key] = uint16ToString(result)
			return result
		}
	} else if len(fields) == 2 {
		a := fields[1]
		var intA uint16
		if isInt(a) {
			intA = stringToUint16(a)
		} else {
			intA = eval(wires, key, a)
		}
		result := ^intA
		wires[key] = uint16ToString(result)
		return result
	}

	return eval(wires, value, wires[value])
}

func main() {
	wires := loadWires("day7.txt")
	partA := uint16ToString(eval(wires, "a", "a"))
	fmt.Printf("Answer for part A: %s\n", partA)
	wires = loadWires("day7.txt")
	wires["b"] = partA
	partB := uint16ToString(eval(wires, "a", "a"))
	fmt.Printf("Answer for part B: %s\n", partB)
}
