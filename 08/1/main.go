package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.Input()
	answer := 0

	for _, line := range lines {
		x := strings.Split(line, "|")
		digits := strings.Fields(x[1])
		for _, digit := range digits {
			switch len(digit) {
			case 2:
				answer++
			case 3:
				answer++
			case 4:
				answer++
			case 7:
				answer++
			}
		}
	}

	fmt.Println(answer)
}
