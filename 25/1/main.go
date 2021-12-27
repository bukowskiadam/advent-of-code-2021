package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const empty = byte('.')
const right = byte('>')
const down = byte('v')

func p(g [][]byte) {
	for i, row := range g {
		fmt.Println(string(row[:len(row)-1]))
		if i == len(g)-2 {
			break
		}
	}
	fmt.Println()
}

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())
	rows := len(lines)
	cols := len(lines[0])

	var grid [][]byte

	for _, line := range lines {
		grid = append(grid, append([]byte(line), byte('X')))
	}
	grid = append(grid, []byte(strings.Repeat("X", cols+1)))

	step := 0

	for {
		// fmt.Printf("After %d steps:\n", step)
		// p(grid)
		moved := false

		for i := 0; i < rows; i++ {
			grid[i][cols] = grid[i][0]
		}

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] != right {
					continue
				}
				next := j + 1

				if grid[i][next] == empty {
					moved = true
					grid[i][next%cols] = right
					grid[i][j] = empty
					j++
				}
			}
		}

		for i := 0; i < cols; i++ {
			grid[rows][i] = grid[0][i]
		}

		for i := 0; i < cols; i++ {
			for j := 0; j < rows; j++ {
				if grid[j][i] != down {
					continue
				}
				next := j + 1

				if grid[next][i] == empty {
					moved = true
					grid[next%rows][i] = down
					grid[j][i] = empty
					j++
				}
			}
		}

		step += 1
		if !moved {
			break
		}
	}

	fmt.Println(step)
}
