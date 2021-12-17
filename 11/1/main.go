package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const SIZE = 10

type Grid [SIZE][SIZE]int8

func parseInput(lines []string) *Grid {
	var grid Grid

	for i, line := range lines {
		f := strings.Split(line, "")
		for j, val := range f {
			valInt, _ := strconv.Atoi(val)
			grid[i][j] = int8(valInt)
		}
	}

	return &grid
}

func flash(grid *Grid, i, j int) {
	grid[i][j] = 100
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			i1 := i + x
			j1 := j + y
			if i1 < 0 || i1 >= SIZE || j1 < 0 || j1 >= SIZE {
				continue
			}
			grid[i1][j1] += 1
		}
	}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			i1 := i + x
			j1 := j + y
			if i1 < 0 || i1 >= SIZE || j1 < 0 || j1 >= SIZE {
				continue
			}
			if grid[i1][j1] >= 10 && grid[i1][j1] < 100 {
				flash(grid, i1, j1)
			}
		}
	}
}

func p(g *Grid, step int) {
	fmt.Printf("After step %d:\n", step)
	for _, l := range g {
		for _, x := range l {
			fmt.Print(x)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {
	lines, _ := utils.ReadLines("11/input.txt")

	grid := parseInput(lines)

	flashes := 0

	STEPS := 100

	for step := 0; step < STEPS; step++ {
		// p(grid, step)
		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				grid[i][j] += 1
			}
		}
		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				if grid[i][j] >= 10 && grid[i][j] < 100 {
					flash(grid, i, j)
				}
			}
		}
		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				if grid[i][j] >= 10 {
					grid[i][j] = 0
					flashes += 1
				}
			}
		}
	}
	// p(grid, STEPS)

	fmt.Println(flashes)
}
