package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const BOARD_SIZE = 5

type Board struct {
	numbers [BOARD_SIZE][]int
}

func makeBoard(lines []string) *Board {
	b := Board{}

	for i, line := range lines {
		b.numbers[i] = utils.MapToNumbers(strings.Fields(line))
	}

	return &b
}

func isWinningBoard(b *Board) bool {
	for i := 0; i < BOARD_SIZE; i++ {
		sumRow := 0
		sumColumn := 0
		for j := 0; j < BOARD_SIZE; j++ {
			sumRow += b.numbers[i][j]
			sumColumn += b.numbers[j][i]
		}
		if sumRow == -BOARD_SIZE || sumColumn == -BOARD_SIZE {
			return true
		}
	}

	return false
}

func calculateScore(b *Board, lastNumber int) int {
	sum := 0
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if b.numbers[i][j] >= 0 {
				sum += b.numbers[i][j]
			}
		}
	}

	return sum * lastNumber
}

func markNumber(b *Board, number int) {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if b.numbers[i][j] == number {
				b.numbers[i][j] = -1
			}
		}
	}
}

func main() {
	lines := utils.Input()
	numbers := utils.MapToNumbers(strings.Split(lines[0], ","))

	var boards []*Board

	for i := 2; i < len(lines); i += BOARD_SIZE + 1 {
		j := i + BOARD_SIZE
		if j > len(lines) {
			j = len(lines)
		}

		boards = append(boards, makeBoard(lines[i:j]))
	}

	// PLAY THE GAME
	for _, n := range numbers {
		for _, b := range boards {
			markNumber(b, n)
		}

		for _, b := range boards {
			if isWinningBoard(b) {
				fmt.Println("score:", calculateScore(b, n))
				return
			}
		}
	}
}
