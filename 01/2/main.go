package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.Input()
	numbers := make([]int, len(lines))

	for index, value := range lines {
		numbers[index], _ = strconv.Atoi(value)
	}

	var increases int
	for i := 3; i < len(numbers); i++ {
		if numbers[i] > numbers[i-3] {
			increases += 1
		}
	}

	fmt.Println(increases)
}
