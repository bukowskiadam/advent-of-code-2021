package main

import (
	"advent/utils"
	"fmt"
)

func main() {
	lines := utils.Input()
	var gamma, ypsilon int

	for index := 0; index < len(lines[0]); index += 1 {
		var zeros, ones int

		for _, value := range lines {
			if value[index] == '0' {
				zeros += 1
			} else {
				ones += 1
			}
		}
		gamma = gamma << 1
		ypsilon = ypsilon << 1
		if ones > zeros {
			gamma = gamma | 1
		} else {
			ypsilon = ypsilon | 1
		}
	}

	fmt.Println(gamma * ypsilon)
}
