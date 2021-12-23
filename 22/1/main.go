package main

import (
	"advent/utils"
	"fmt"
)

const size = 101

var cube [size][size][size]bool

func b(i int) int {
	if i < -50 {
		return -50
	}
	if i > 50 {
		return 50
	}
	return i
}

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())

	for _, line := range lines {
		state := true
		coordinates := line[3:]
		if line[1] == 'f' {
			state = false
			coordinates = line[4:]
		}
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(coordinates, "x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)

		if x1 > 50 || x2 < -50 || y1 > 50 || y2 < -50 || z1 > 50 || z2 < -50 {
			continue
		}

		for x := b(x1); x <= b(x2); x++ {
			for y := b(y1); y <= b(y2); y++ {
				for z := b(z1); z <= b(z2); z++ {
					cube[x+50][y+50][z+50] = state
				}
			}
		}
	}
	ans := 0

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				if cube[x][y][z] {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
