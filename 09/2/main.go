package main

import (
	"advent/utils"
	"fmt"
	"sort"
)

func main() {
	lines := utils.Input()

	maxY := len(lines)
	maxX := len(lines[0])

	grid := [][]byte{}

	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	for y, row := range grid {
		for x := range row {
			grid[y][x] -= '0'
		}
	}

	visited := make([][]bool, maxY)
	for i := range visited {
		visited[i] = make([]bool, maxX)
	}

	value := func(x, y int) byte {
		if x < 0 || y < 0 || x >= maxX || y >= maxY {
			return 10
		}
		return grid[y][x]
	}

	wasVisited := func(x, y int) bool {
		if x < 0 || y < 0 || x >= maxX || y >= maxY {
			return true
		}
		return visited[y][x]
	}

	var calculateBasinSize func(x, y int) int

	calculateBasinSize = func(x, y int) int {
		visited[y][x] = true
		size := 1

		step := func(x, y int) {
			if !wasVisited(x, y) && value(x, y) < 9 {
				size += calculateBasinSize(x, y)
			}
		}

		step(x-1, y)
		step(x+1, y)
		step(x, y-1)
		step(x, y+1)

		return size
	}

	basins := []int{}

	for y, row := range grid {
		for x, val := range row {
			if wasVisited(x, y) || val == 9 {
				continue
			}
			basinSize := calculateBasinSize(x, y)
			basins = append(basins, basinSize)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	answer := 1

	for _, v := range basins[:3] {
		answer *= v
	}

	fmt.Println(answer)
}
