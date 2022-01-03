package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, _ := utils.ReadLines("2/input.txt")
	var horizontal, depth int

	for _, value := range lines {
		fields := strings.Fields(value)
		number, _ := strconv.Atoi(fields[1])
		if fields[0] == "forward" {
			horizontal += number
		}
		if fields[0] == "down" {
			depth += number
		}
		if fields[0] == "up" {
			depth -= number
		}
	}

	fmt.Println(horizontal * depth)
}
