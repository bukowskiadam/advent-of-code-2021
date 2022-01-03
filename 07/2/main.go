package main

import (
	"advent/utils"
	"fmt"
	"math"
)

func main() {
	// READ THE INPUT
	lines, _ := utils.ReadLines("7/input.txt")
	positions := utils.ReadCommaSeparatedInts(lines[0])

	minFuel := math.MaxInt

	for i, max := utils.FindMinInt(positions), utils.FindMaxInt(positions); i <= max; i++ {
		sum := 0
		for _, val := range positions {
			sum += utils.Factorial(utils.Abs(val - i))
		}
		if sum < minFuel {
			minFuel = sum
		}
	}

	fmt.Println(minFuel)
}
