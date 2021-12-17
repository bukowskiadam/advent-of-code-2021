package main

import (
	"advent/utils"
	"fmt"
	"sort"
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
	lines, _ := utils.ReadLines("10/input.txt")
	scoresMap := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	var allScores []int

	for _, line := range lines {
		heap := make([]rune, len(line))
		ptr := 0
		corrupted := false
		for _, char := range line {
			if isOpening(char) {
				heap[ptr] = char
				ptr++
			} else {
				ptr--
				if !isMatching(heap[ptr], char) {
					corrupted = true
					break
				}
			}
		}
		if corrupted {
			continue
		}
		score := 0
		for ptr > 0 {
			ptr--
			score = 5*score + scoresMap[heap[ptr]]
		}
		allScores = append(allScores, score)
	}
	sort.Ints(allScores)
	middleIndex := int(len(allScores) / 2)

	fmt.Println(allScores[middleIndex])
}
