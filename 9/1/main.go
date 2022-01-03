package main

import (
	"advent/utils"
	"fmt"
)

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())

	riskLevel := 0

	grid := [][]byte{}

	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	var max byte = '9' + 1

	value := func(x, y int) byte {
		if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
			return max
		}
		return grid[y][x]
	}

	for y, row := range grid {
		for x, val := range row {
			if value(x-1, y) > val &&
				value(x+1, y) > val &&
				value(x, y-1) > val &&
				value(x, y+1) > val {
				riskLevel += int(val-'0') + 1
			}
		}
	}

	fmt.Println(riskLevel)
}
