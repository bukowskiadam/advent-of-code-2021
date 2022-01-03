package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func main() {
	lines, _ := utils.ReadLines("1/input.txt")
	numbers := make([]int, len(lines))

	for index, value := range lines {
		numbers[index], _ = strconv.Atoi(value)
	}

	var increases int
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			increases += 1
		}
	}

	fmt.Println(increases)
}
