package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	s := strings.Split(input, ",")
	counters := make([]int, len(s))
	for i, val := range s {
		counters[i], _ = strconv.Atoi(val)
	}
	return counters
}

const DAYS = 80

func main() {
	// READ THE INPUT
	lines, _ := utils.ReadLines("6/input.txt")
	counters := parseInput(lines[0])
	fmt.Println(counters)

	for d := 0; d < DAYS; d++ {
		for i := range counters {
			counters[i] -= 1
			if counters[i] < 0 {
				counters = append(counters, 8)
				counters[i] = 6
			}
		}
	}

	fmt.Println(len(counters))
}
