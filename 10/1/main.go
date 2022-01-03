package main

import (
	"advent/utils"
	"fmt"
)

func isOpening(c rune) bool {
	return c == '[' || c == '(' || c == '{' || c == '<'
}

func isMatching(o rune, c rune) bool {
	if c == ']' {
		return o == '['
	}
	if c == '}' {
		return o == '{'
	}
	if c == '>' {
		return o == '<'
	}
	if c == ')' {
		return o == '('
	}
	return false
}

func main() {
	lines := utils.Input()
	scoresMap := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	score := 0

	for _, line := range lines {
		heap := make([]rune, len(line))
		ptr := 0
		for _, char := range line {
			if isOpening(char) {
				heap[ptr] = char
				ptr++
			} else {
				ptr--
				if !isMatching(heap[ptr], char) {
					score += scoresMap[char]
					break
				}
			}
		}
	}

	fmt.Println(score)
}
