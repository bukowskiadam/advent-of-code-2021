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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// READ THE INPUT
	inputLines, _ := utils.ReadLines("5/input.txt")
	lines := parseInput(inputLines)

	var field [SIZE][SIZE]int

	for _, line := range lines {
		if line.from.x != line.to.x && line.from.y != line.to.y {
			continue
		}
		startX := min(line.from.x, line.to.x)
		endX := max(line.from.x, line.to.x)
		startY := min(line.from.y, line.to.y)
		endY := max(line.from.y, line.to.y)

		for i := startX; i <= endX; i++ {
			for j := startY; j <= endY; j++ {
				field[j][i] += 1
			}
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
