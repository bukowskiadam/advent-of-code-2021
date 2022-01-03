package main

import (
	"advent/utils"
	"fmt"
)

const END_SCORE = 21

type Roll struct {
	sum   int
	times int
}

var possibleRolls []Roll

func play(scores, pos [2]int, player int) (int, int) {
	if scores[0] >= END_SCORE {
		return 1, 0
	}
	if scores[1] >= END_SCORE {
		return 0, 1
	}

	otherPlayer := (player + 1) % 2
	sum1, sum2 := 0, 0
	for _, roll := range possibleRolls {
		var newPos [2]int
		var newScores [2]int
		newPos[player] = (pos[player] + roll.sum) % 10
		newPos[otherPlayer] = pos[otherPlayer]
		newScores[player] = scores[player] + newPos[player] + 1
		newScores[otherPlayer] = scores[otherPlayer]

		win1, win2 := play(newScores, newPos, otherPlayer)
		sum1 += win1 * roll.times
		sum2 += win2 * roll.times
	}
	return sum1, sum2
}

func main() {
	lines := utils.Input()

	possibleRolls = []Roll{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}

	var pos [2]int
	var scores [2]int
	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &pos[0])
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &pos[1])

	// make it 0-9, instead of 1-10
	pos[0]--
	pos[1]--

	player := 0

	w1, w2 := play(scores, pos, player)

	if w1 > w2 {
		fmt.Println(w1)
	} else {
		fmt.Println(w2)
	}
}
