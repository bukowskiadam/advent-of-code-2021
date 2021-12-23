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

func (cube *Cube) contains(other *Cube) bool {
	xInRange := (other.start.x < cube.end.x && other.start.x >= cube.start.x) ||
		(other.end.x <= cube.end.x && other.end.x > cube.start.x)

	yInRange := (other.start.y < cube.end.y && other.start.y >= cube.start.y) ||
		(other.end.y <= cube.end.y && other.end.y > cube.start.y)

	zInRange := (other.start.z < cube.end.z && other.start.z >= cube.start.z) ||
		(other.end.z <= cube.end.z && other.end.z > cube.start.z)

	return xInRange && yInRange && zInRange
}

func (cube *Cube) intersects(other *Cube) bool {
	return cube.contains(other) || other.contains(cube)
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
			start: Point{max(cube.start.x, other.start.x), max(cube.start.y, other.end.y), max(cube.start.z, other.start.z)},
			end:   Point{min(cube.end.x, other.end.x), cube.end.y, min(cube.end.z, other.end.z)},
		},
		{ //down
			start: Point{max(cube.start.x, other.start.x), cube.start.y, max(cube.start.z, other.start.z)},
			end:   Point{min(cube.end.x, other.end.x), min(cube.end.y, other.start.y), min(cube.end.z, other.end.z)},
		},
	}

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
	for _, cube := range cubes {
		if cube.intersects(&newCube) {
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
	var newCubes []Cube

	for i := 0; i < len(cubes); {
		if cubes[i].intersects(&newCube) {
			smallCubes := cubes[i].subtract(&newCube)
			newCubes = append(newCubes, smallCubes...)
			cubes = append(cubes[:i], cubes[i+1:]...)
		} else {
			i++
		}
	}

	cubes = append(cubes, newCubes...)
}

func sum() int {
	sum := 0
	for _, p := range cubes {
		sum += p.size()
	}
	return sum
}

// part 1 - naive
const size = 101

var cube [size][size][size]bool

func b(i int) int {
	if i < -50 {
		return -50
	}
	if i > 50 {
		return 50
	}
	return i
}

func sumPt1() int {
	ans := 0

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				if cube[x][y][z] {
					ans++
				}
			}
		}
	}

	return ans
}

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())

	for _, line := range lines {
		state := true
		coordinates := line[3:]
		if line[1] == 'f' {
			state = false
			coordinates = line[4:]
		}
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(coordinates, "x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)

		if !(x1 > 50 || x2 < -50 || y1 > 50 || y2 < -50 || z1 > 50 || z2 < -50) {
			for x := b(x1); x <= b(x2); x++ {
				for y := b(y1); y <= b(y2); y++ {
					for z := b(z1); z <= b(z2); z++ {
						cube[x+50][y+50][z+50] = state
					}
				}
			}
		}

		p := Cube{
			start: Point{x1, y1, z1},
			end:   Point{x2 + 1, y2 + 1, z2 + 1},
		}

		fmt.Printf("\n%s\n", line)
		fmt.Printf("cube %t size: %d\n", state, p.size())
		if !p.isValid() {
			panic("invalid cube")
		}

		if state {
			add(p)
		} else {
			subtract(p)
		}

		// verify cubes after the step
		for i, c := range cubes {
			for i2, c2 := range cubes {
				if i == i2 {
					continue
				}
				if c.intersects(&c2) {
					fmt.Println(i, c)
					fmt.Println(i2, c2)
					panic("error")
				}
			}
		}

		// result
		fmt.Println("pt 1:", sumPt1())
		fmt.Println("pt 2:", sum())
	}

	// fmt.Println(sum())
}
