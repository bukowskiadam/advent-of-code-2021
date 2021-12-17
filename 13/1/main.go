package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const MAX_SIZE = 1400

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())

	dots := [MAX_SIZE][MAX_SIZE]bool{}
	maxX := 0
	maxY := 0
	emptyLineIndex := 0

	for i, l := range lines {
		if len(l) == 0 {
			emptyLineIndex = i
		}
	}

	for _, line := range lines[:emptyLineIndex] {
		pos := strings.Split(line, ",")
		posX, _ := strconv.Atoi(pos[0])
		posY, _ := strconv.Atoi(pos[1])
		dots[posY][posX] = true

		if posX > maxX {
			maxX = posX
		}
		if posY > maxY {
			maxY = posY
		}
	}
	maxX++
	maxY++

	// for y := 0; y < maxY; y++ {
	// 	for x := 0; x < maxX; x++ {
	// 		if dots[y][x] {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	for _, line := range lines[emptyLineIndex+1:] {
		var axis string
		l := strings.Split(line, "=")

		fmt.Sscanf(l[0], "fold along %s", &axis)
		where, _ := strconv.Atoi(l[1])

		fmt.Println(axis, where)

		if axis == "x" {
			for y := 0; y < maxY; y++ {
				for x := 0; x < where; x++ {
					dots[y][x] = dots[y][x] || dots[y][2*where-x]
				}
			}
			maxX = where
		} else {
			for y := 0; y < where; y++ {
				for x := 0; x < maxX; x++ {
					dots[y][x] = dots[y][x] || dots[2*where-y][x]
				}
			}
			maxY = where
		}

		// fmt.Println("---")
		// for y := 0; y < maxY; y++ {
		// 	for x := 0; x < maxX; x++ {
		// 		if dots[y][x] {
		// 			fmt.Print("#")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Print("\n")
		// }
		break
	}

	visible := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if dots[y][x] {
				visible++
			}
		}
	}

	fmt.Println(visible)
}
