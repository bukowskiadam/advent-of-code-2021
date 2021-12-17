package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	from Point
	to   Point
}

func parsePoint(str string) Point {
	coordinates := strings.Split(str, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])

	return Point{x: x, y: y}
}

func parseLine(str string) Line {
	points := strings.Split(str, " -> ")

	return Line{from: parsePoint(points[0]), to: parsePoint(points[1])}
}

func parseInput(lines []string) []Line {
	ret := make([]Line, len(lines))

	for i, line := range lines {
		ret[i] = parseLine(line)
	}

	return ret
}

const SIZE = 1000

func direction(from, to int) int {
	if from < to {
		return 1
	}
	if from > to {
		return -1
	}
	return 0
}

func main() {
	// READ THE INPUT
	inputLines, _ := utils.ReadLines("5/input.txt")
	lines := parseInput(inputLines)

	var field [SIZE][SIZE]int

	for _, line := range lines {
		directionX := direction(line.from.x, line.to.x)
		directionY := direction(line.from.y, line.to.y)

		x, y := line.from.x, line.from.y
		for {
			field[y][x] += 1
			if x == line.to.x && y == line.to.y {
				break
			}
			x, y = x+directionX, y+directionY
		}
	}

	crossings := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if field[i][j] > 1 {
				crossings += 1
			}
		}
	}

	fmt.Println(crossings)
}
