package main

import (
	"advent/utils"
	"fmt"
)

func main() {
	lines := utils.Input()

	var pos [2]int
	var points [2]int
	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &pos[0])
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &pos[1])

	fmt.Println(pos)

	pos[0]--
	pos[1]--

	dice := 1
	rolls := 0
	player := 0

	for points[0] < 1000 && points[1] < 1000 {
		sum := 3*dice + 3
		pos[player] += sum
		pos[player] = pos[player] % 10
		points[player] += pos[player] + 1
		dice += 3
		if dice > 100 {
			dice -= 100
		}
		rolls += 3
		player = (player + 1) % 2
		fmt.Println(points)
	}

	loser := 1
	if points[0] < points[1] {
		loser = 0
	}
	fmt.Println(points[loser] * rolls)
}
