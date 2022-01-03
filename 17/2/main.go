package main

import (
	"advent/utils"
	"fmt"
)

func main() {
	lines := utils.Input()

	var x1, x2, y1, y2 int
	fmt.Sscanf(lines[0], "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)

	var isWithinTarget = func(x, y int) bool {
		return x >= x1 && x <= x2 && y >= y1 && y <= y2
	}

	total := 0
	for dx := 1; dx <= x2; dx++ {
		for dy := y1; dy <= -y1; dy++ {
			cdx, cdy := dx, dy
			x, y := 0, 0

			for x <= x2 && y >= y1 {
				x += cdx
				y += cdy
				if cdx > 0 {
					cdx -= 1
				}
				cdy -= 1

				if cdx == 0 && x < x1 {
					break
				}

				if isWithinTarget(x, y) {
					total += 1
					break
				}
			}
		}
	}

	fmt.Println(total)
}
