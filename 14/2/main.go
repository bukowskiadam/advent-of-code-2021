package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strings"
)

const STEPS = 40

func main() {
	lines := utils.Input()

	polymer := lines[0]
	rules := map[string]string{}
	pairs := map[string]uint64{}
	counts := map[byte]uint64{}

	for _, r := range lines[2:] {
		rr := strings.Split(r, " -> ")
		rules[rr[0]] = rr[1]
	}

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pairs[pair] += 1
	}

	for i := range polymer {
		counts[polymer[i]] += 1
	}

	for step := 0; step < STEPS; step++ {
		newPairs := map[string]uint64{}
		for pair, pairCount := range pairs {
			if insert, ok := rules[pair]; ok {
				counts[insert[0]] += pairCount
				np1 := pair[:1] + insert
				np2 := insert + pair[1:]
				newPairs[np1] += pairCount
				newPairs[np2] += pairCount
			} else {
				newPairs[pair] = pairCount
			}
		}
		pairs = newPairs
	}

	var max uint64 = 0
	var min uint64 = math.MaxUint64
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
