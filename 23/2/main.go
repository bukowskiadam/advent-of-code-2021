package main

import (
	"advent/utils"
	"fmt"
	"math"
)

const wall = byte('#')
const empty = byte('.')
const amphipodsCount = 16

var world [][]byte
var amphipods []Amphipod
var states map[State]int

type Position struct {
	x, y int
}

type Step struct {
	position Position
	cost     int
}

type Amphipod struct {
	position Position
	cost     int16
	face     byte
	home     int8
	moves    int8
}

type AmphipodState struct {
	position Position
	face     byte
}

type State struct {
	amphipods [amphipodsCount]AmphipodState
}

func (p *Position) mark(face byte) {
	world[p.y][p.x] = face
}

func (p *Position) field() byte {
	return world[p.y][p.x]
}

func isSolved() bool {
	homesY := 2
	for _, a := range amphipods {
		targetCol := int(a.home)*2 + 1

		if a.position.y < homesY || a.position.x != targetCol {
			return false
		}
	}
	return true
}

func loadAmphipods() {
	costMap := map[byte]int16{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
	homeMap := map[byte]int8{'A': 1, 'B': 2, 'C': 3, 'D': 4}
	for y, row := range world {
		for x, field := range row {
			switch field {
			case 'A', 'B', 'C', 'D':
				amphipods = append(amphipods, Amphipod{
					position: Position{x, y},
					cost:     costMap[field],
					face:     field,
					home:     homeMap[field],
					moves:    2,
				})
			}
		}
	}
}

func createState() State {
	s := State{}
	i := 0
	for y, row := range world {
		for x, field := range row {
			switch field {
			case 'A', 'B', 'C', 'D':
				s.amphipods[i] = AmphipodState{
					position: Position{x, y},
					face:     field,
				}
				i++
			}
		}
	}
	return s
}

func (a *Amphipod) canStop(pos *Position) bool {
	if pos.y == 1 {
		if a.moves == 0 {
			return false
		}

		switch pos.x {
		case 3, 5, 7, 9:
			return false
		default:
			return true
		}
	}

	if int(a.home)*2+1 != pos.x {
		return false
	}

	down := *pos
	down.y++

	return down.field() == wall || down.field() == a.face
}

func (a *Amphipod) getNextPositions() []Step {
	type QueueElement struct {
		position Position
		distance int
	}

	steps := []Step{}
	visited := map[Position]bool{a.position: true}
	queue := []QueueElement{{a.position, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		distance := current.distance + 1
		cost := distance * int(a.cost)

		var nextPositions []Position

		if current.position.y == 1 {
			nextPositions = []Position{
				{current.position.x + 1, current.position.y},
				{current.position.x - 1, current.position.y},
				{current.position.x, current.position.y + 1},
			}
		} else {
			nextPositions = []Position{
				{current.position.x, current.position.y + 1},
				{current.position.x, current.position.y - 1},
			}
		}

		for _, p := range nextPositions {
			if !visited[p] && p.field() == empty {
				visited[p] = true
				queue = append(queue, QueueElement{p, distance})
				if a.canStop(&p) {
					steps = append(steps, Step{p, cost})
				}
			}
		}
	}

	return steps
}

func printWorld() {
	for _, v := range world {
		fmt.Println(string(v))
	}
	fmt.Println()
}

var tries int

func find(level int) int {
	currentState := createState()

	if val, present := states[currentState]; present {
		return val
	}

	// printWorld()
	tries++
	if isSolved() {
		return 0
	}

	min := math.MaxInt / 2
	for i := range amphipods {
		a := &amphipods[i]
		if a.moves == 0 {
			continue
		}
		originalPosition := a.position
		originalPosition.mark(empty)
		a.moves--

		nextPositions := a.getNextPositions()
		for _, next := range nextPositions {
			a.position = next.position
			next.position.mark(a.face)

			nextCost := next.cost + find(level+1)
			if nextCost < min {
				min = nextCost
			}

			next.position.mark(empty)
		}

		a.moves++
		a.position = originalPosition
		originalPosition.mark(a.face)
	}

	states[currentState] = min

	return min
}

func main() {
	lines := utils.Input()

	for _, line := range lines {
		world = append(world, []byte(line))
	}

	loadAmphipods()

	states = make(map[State]int)
	fmt.Println("Result", find(0))
}
