package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	s := strings.Split(input, ",")
	positions := make([]int, len(s))
	for i, val := range s {
		positions[i], _ = strconv.Atoi(val)
	}
	return positions
}

func findMin(input []int) (int, int) {
	m := 0
	pos := 0
	for i, e := range input {
		if i == 0 || e < m {
			m = e
			pos = i
		}
	}
	return m, pos
}

func findMax(input []int) (int, int) {
	m := 0
	pos := 0
	for i, e := range input {
		if i == 0 || e > m {
			m = e
			pos = i
		}
	}
	return m, pos
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// READ THE INPUT
	lines, _ := utils.ReadLines("7/input.txt")
	positions := parseInput(lines[0])

	min, _ := findMin(positions)
	max, _ := findMax(positions)

	fuelUsed := make([]int, max+1)

	for i := min; i <= max; i++ {
		sum := 0
		for _, val := range positions {
			sum += abs(val - i)
		}
		fuelUsed[i] = sum
	}

	min, pos := findMin(fuelUsed)

	fmt.Println(min, pos)

}
