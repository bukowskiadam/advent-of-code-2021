package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	s := strings.Split(input, ",")
	fishes := make([]int, 9)
	for _, val := range s {
		life, _ := strconv.Atoi(val)
		fishes[life] += 1
	}
	return fishes
}

const DAYS = 256

func main() {
	lines := utils.Input()
	fishes := parseInput(lines[0])
	fmt.Println(fishes)

	for d := 0; d < DAYS; d++ {
		newFish := fishes[0]
		fishes = append(fishes[1:], 0)
		fishes[6] += newFish
		fishes[8] += newFish
	}

	var sum int64
	for _, v := range fishes {
		sum += int64(v)
	}

	fmt.Println(sum)
}
