package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func isRegister(s string) bool {
	switch s {
	case "w", "x", "y", "z":
		return true
	}
	return false
}

func toInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(s)
		fmt.Println(err)
		panic("wrong number")
	}
	return value
}

const size = 14

func execute(program []string, input [size]int) bool {
	registers := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	inputIndex := 0

	for _, command := range program {
		chunks := strings.Fields(command)
		operation := chunks[0]
		register := chunks[1]

		if operation == "inp" {
			if inputIndex >= len(input) {
				panic("not enough input numbers")
			}

			registers[register] = input[inputIndex]
			inputIndex++
			continue
		}
		other := chunks[2]
		var otherVal int

		if isRegister(other) {
			otherVal = registers[other]
		} else {
			otherVal = toInt(other)
		}

		switch operation {
		case "add":
			registers[register] += otherVal
		case "mul":
			registers[register] *= otherVal
		case "div":
			registers[register] /= otherVal
		case "mod":
			registers[register] %= otherVal
		case "eql":
			if registers[register] == otherVal {
				registers[register] = 1
			} else {
				registers[register] = 0
			}
		}
	}

	return registers["z"] == 0
}

func verify(program []string, input [size]int) {
	success := execute(program, input)

	for _, d := range input {
		fmt.Print(d)
	}

	fmt.Println()
	fmt.Println(success)
}

func main() {
	program, _ := utils.ReadLines(utils.ReadFileFromArgs())

	// Largest
	largest := [size]int{1, 2, 9, 3, 4, 9, 9, 8, 9, 4, 9, 1, 9, 9}

	// Smallest
	smallest := [size]int{1, 1, 7, 1, 1, 6, 9, 1, 6, 1, 2, 1, 8, 9}

	fmt.Println("Largest:")
	verify(program, largest)

	fmt.Println("\nSmallest:")
	verify(program, smallest)
}
