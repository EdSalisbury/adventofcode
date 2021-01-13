package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Global!
var wires map[string]string

func loadWires(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wires = make(map[string]string)
	pattern := regexp.MustCompile(`^(.*?)\s+->\s+(\w+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		fields := pattern.FindStringSubmatch(line)
		wires[fields[2]] = fields[1]
	}
}

func isValue(expr string) bool {
	pattern := regexp.MustCompile(`^\d+$`)
	return pattern.Match([]byte(expr))
}

func intVal(str string) uint16 {
	val, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(val)
}

func trace(wire string) uint16 {
	fmt.Println(wire)
	if isValue(wire) {
		return intVal(wire)
	}
	fields := strings.Split(wire, " ")
	if len(fields) == 3 {
		a := fields[0]
		op := fields[1]
		b := fields[2]
		switch op {
		case "AND":
			return trace(a) & trace(b)
		case "OR":
			return trace(a) | trace(b)
		case "LSHIFT":
			return trace(a) << trace(b)
		case "RSHIFT":
			return trace(a) >> trace(b)
		}
	} else if len(fields) == 2 {
		return ^trace(fields[1])
	}
	_, found := wires[wire]
	if !found {
		panic("Wire not found!")
	}
	return trace(wires[wire])
}

func main() {
	loadWires("day7.txt")
	fmt.Println(trace("a"))
}
