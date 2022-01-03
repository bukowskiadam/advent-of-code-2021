package main

import (
	"advent/utils"
	"fmt"
	"math"
)

func main() {
	lines := utils.Input()
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
