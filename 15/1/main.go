package main

import (
	"advent/utils"
	"fmt"
	"math"
)

type pos struct {
	x, y int
}

type gridPoint struct {
	value  int
	length int
}

func main() {
	lines := utils.Input()

	targetX := len(lines[0]) - 1
	targetY := len(lines) - 1

	grid := [][]gridPoint{}

	for _, line := range lines {
		ints := []gridPoint{}
		for _, c := range line {
			value := int(c) - int('0')
			ints = append(ints, gridPoint{value, math.MaxInt})
		}
		grid = append(grid, ints)
	}

	grid[0][0].length = 0
	queue := []pos{}
	queue = append(queue, pos{0, 0})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				if x != 0 && y != 0 {
					continue
				}
				nx := p.x + x
				ny := p.y + y
				if nx >= 0 && nx <= targetX && ny >= 0 && ny <= targetY {
					newLength := grid[p.y][p.x].length + grid[ny][nx].value
					if newLength < grid[ny][nx].length {
						grid[ny][nx].length = newLength
						queue = append(queue, pos{nx, ny})
					}
				}
			}
		}
	}

	fmt.Println(grid[targetY][targetX].length)
}
