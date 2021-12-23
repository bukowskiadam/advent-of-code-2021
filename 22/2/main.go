package main

import (
	"advent/utils"
	"fmt"
)

type Point struct {
	x, y, z int
}

type Cube struct {
	start, end Point
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (cube *Cube) size() int {
	x := cube.end.x - cube.start.x
	y := cube.end.y - cube.start.y
	z := cube.end.z - cube.start.z

	return x * y * z
}

func (cube *Cube) intersects(other *Cube) bool {
	xInRange := (other.start.x < cube.end.x && other.start.x >= cube.start.x) ||
		(other.end.x < cube.end.x && other.end.x >= cube.start.x)

	yInRange := (other.start.y < cube.end.y && other.start.y >= cube.start.y) ||
		(other.end.y < cube.end.y && other.end.y >= cube.start.y)

	zInRange := (other.start.z < cube.end.z && other.start.z >= cube.start.z) ||
		(other.end.z < cube.end.z && other.end.z >= cube.start.z)

	return xInRange && yInRange && zInRange
}

func (cube *Cube) isValid() bool {
	return cube.start.x < cube.end.x &&
		cube.start.y < cube.end.y &&
		cube.start.z < cube.end.z
}

func (cube *Cube) subtract(other *Cube) []Cube {
	smallCubes := []Cube{
		{ // front
			start: cube.start,
			end:   Point{cube.end.x, cube.end.y, min(cube.end.z, other.start.z)},
		},
		{ // back
			start: Point{cube.start.x, cube.start.y, max(cube.start.z, other.end.z)},
			end:   cube.end,
		},
		{ // left
			start: Point{cube.start.x, cube.start.y, max(cube.start.z, other.start.z)},
			end:   Point{min(cube.end.x, other.start.x), cube.end.y, min(cube.end.z, other.end.z)},
		},
		{ // right
			start: Point{max(cube.start.x, other.end.x), cube.start.y, max(cube.start.z, other.start.z)},
			end:   Point{cube.end.x, cube.end.y, min(cube.end.z, other.end.z)},
		},
		{ // up
			start: Point{other.start.x, other.end.y, other.start.z},
			end:   Point{other.end.x, cube.end.y, other.end.z},
		},
		{ //down
			start: Point{other.start.x, cube.start.y, other.start.z},
			end:   Point{other.end.x, other.start.y, other.end.z},
		},
	}
	fmt.Println(smallCubes)

	validCubes := []Cube{}
	for _, c := range smallCubes {
		if c.isValid() {
			validCubes = append(validCubes, c)
		}
	}

	return validCubes
}

var cubes []Cube

func add(newCube Cube) {
	fmt.Println("current cubes", cubes)
	for _, cube := range cubes {
		if newCube.intersects(&cube) {
			fmt.Println("intersects", newCube, cube)
			smallCubes := newCube.subtract(&cube)
			for _, s := range smallCubes {
				add(s)
			}
			return
		}
	}
	cubes = append(cubes, newCube)
}

func subtract(newCube Cube) {

}

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())

	c := Cube{
		start: Point{1, 1, 1},
		end:   Point{5, 5, 5},
	}
	c2 := Cube{
		start: Point{6, 6, 6},
		end:   Point{7, 7, 7},
	}
	fmt.Println(c.subtract(&c2))
	return
	for _, line := range lines {
		state := true
		coordinates := line[3:]
		if line[1] == 'f' {
			state = false
			coordinates = line[4:]
		}
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(coordinates, "x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)

		p := Cube{
			start: Point{x1, y1, z1},
			end:   Point{x2 + 1, y2 + 1, z2 + 1},
		}

		if state {
			add(p)
		} else {
			subtract(p)
		}
	}

	sum := 0
	for _, p := range cubes {
		sum += p.size()
	}

	fmt.Println(sum)
}
