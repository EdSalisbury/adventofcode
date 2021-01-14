package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/golang-collections/collections/stack"
)

// Global!
var wires map[string]string
var Key *stack.Stack

func loadWires(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wires = make(map[string]string)
	//pattern := regexp.MustCompile(`^(.*?)\s+->\s+(\w+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		//fields := pattern.FindStringSubmatch(line)
		fields := strings.Split(line, "->")
		wires[strings.TrimSpace(fields[1])] = strings.TrimSpace(fields[0])
	}
}

func trace(wire string) uint16 {
	fmt.Println(wire)
	val, err := strconv.ParseUint(wire, 10, 16)
	if err == nil {
		//wires[key] = wire
		//fmt.Printf("Assigning %s = %s\n", key, wire)
		// fmt.Printf("Assigning %s = %s\n", Key.Peek(), string(val))
		// wires[Key.Pop().(string)] = string(val)
		fmt.Printf("Popping %s off the stack\n", Key.Pop().(string))
		fmt.Printf("Returning %s\n", wire)
		return uint16(val)
	}

	fields := strings.Split(wire, " ")
	if len(fields) == 3 {
		var intA uint16
		var intB uint16
		a := fields[0]
		op := fields[1]
		b := fields[2]
		valA, err := strconv.ParseUint(a, 10, 16)
		if err == nil {
			intA = uint16(valA)
		} else {
			intA = trace(a)
		}

		valB, err := strconv.ParseUint(b, 10, 16)
		if err == nil {
			intB = uint16(valB)
		} else {
			intB = trace(b)
		}

		switch op {
		case "AND":
			traceA := intA
			fmt.Printf("traceA = %d\n", traceA)
			traceB := intB
			fmt.Printf("traceB = %d\n", traceB)
			result := traceA & traceB
			fmt.Printf("result = %d\n", result)
			fmt.Printf("Assigning %s = %d\n", Key.Peek(), result)
			wires[Key.Pop().(string)] = strconv.Itoa(int(result))
			return result
		case "OR":
			traceA := intA
			fmt.Printf("traceA = %d\n", traceA)
			traceB := intB
			fmt.Printf("traceB = %d\n", traceB)
			result := traceA | traceB
			fmt.Printf("result = %d\n", result)
			fmt.Printf("Assigning %s = %d\n", Key.Peek(), result)
			wires[Key.Pop().(string)] = strconv.Itoa(int(result))
			return result
		case "LSHIFT":
			traceA := intA
			fmt.Printf("traceA = %d\n", traceA)
			//traceB := trace(b)
			traceB := intB
			fmt.Printf("traceB = %d\n", traceB)
			result := traceA << traceB
			fmt.Printf("result = %d\n", result)
			fmt.Printf("Assigning %s = %d\n", Key.Peek(), result)
			wires[Key.Pop().(string)] = strconv.Itoa(int(result))
			return result
		case "RSHIFT":
			traceA := trace(a)
			fmt.Printf("traceA = %d\n", traceA)
			traceB := intB
			fmt.Printf("traceB = %d\n", traceB)
			result := traceA >> traceB
			fmt.Printf("result = %d\n", result)
			fmt.Printf("Assigning %s = %d\n", Key.Peek(), result)
			wires[Key.Pop().(string)] = strconv.Itoa(int(result))
			return result
		}
	} else if len(fields) == 2 {
		var intB uint16
		valB, err := strconv.ParseUint(fields[1], 10, 16)
		if err == nil {
			intB = uint16(valB)
		} else {
			intB = trace(fields[1])
		}
		traceA := intB
		fmt.Printf("traceA = %d\n", traceA)
		result := ^traceA
		fmt.Printf("result = %d\n", result)
		fmt.Printf("Assigning %s = %d\n", Key.Peek(), result)
		wires[Key.Pop().(string)] = strconv.Itoa(int(result))
		return result
	}
	_, found := wires[wire]
	if !found {
		panic(fmt.Sprintf("Wire %s not found!", wire))
	}
	fmt.Printf("Pushing wire %s to stack\n", wire)
	Key.Push(wire)
	return trace(wires[wire])
}

func main() {
	loadWires("day7_ray.txt")
	Key = stack.New()
	fmt.Printf("%s = %d\n", os.Args[1], trace(os.Args[1]))
	fmt.Printf("%s\n", wires["b"])
	//fmt.Println(len(wires))
}
