package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func filter(list []string, index int, val byte) []string {
	var newList []string

	for _, value := range list {
		if value[index] == val {
			newList = append(newList, value)
		}
	}

	return newList
}

func getFinalValue(list []string, one, zero byte) int64 {
	output := list

	for index := 0; index < len(list[0]); index += 1 {
		var zeros, ones int

		for _, value := range output {
			if value[index] == '0' {
				zeros += 1
			} else {
				ones += 1
			}
		}
		if ones >= zeros {
			output = filter(output, index, one)
		} else {
			output = filter(output, index, zero)
		}
		if len(output) == 1 {
			break
		}
	}
	number, _ := strconv.ParseInt(output[0], 2, 64)
	return number
}

func main() {
	lines, _ := utils.ReadLines("3/input.txt")

	oxygen := getFinalValue(lines, '1', '0')
	co2 := getFinalValue(lines, '0', '1')

	fmt.Println(oxygen * co2)
}
