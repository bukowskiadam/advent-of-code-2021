package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const STEPS = 10

func main() {
	lines := utils.Input()

	polymer := lines[0]
	rules := map[string]string{}

	for _, r := range lines[2:] {
		rr := strings.Split(r, " -> ")
		rules[rr[0]] = rr[1]
	}

	for step := 0; step < STEPS; step++ {
		for i := 0; i < len(polymer)-1; i++ {
			pair := polymer[i : i+2]
			if insert, ok := rules[pair]; ok {
				polymer = polymer[:i+1] + insert + polymer[i+1:]
				i++
			}
		}
	}

	counts := map[rune]int{}

	for _, v := range polymer {
		counts[v] += 1
	}

	max := 0
	min := len(polymer)
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	fmt.Println(max - min)
}
