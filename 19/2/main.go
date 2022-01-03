package main

import (
	"advent/utils"
	"fmt"
)

type Pos struct {
	x, y, z int
}

type Diff struct {
	dx, dy, dz int
}

type Scanner struct {
	matched bool
	beacons []Pos
}

func transpose(p Pos, i int) Pos {
	switch i {
	case 0:
		return p
	case 1:
		return Pos{p.x, p.z, -p.y}
	case 2:
		return Pos{p.x, -p.y, -p.z}
	case 3:
		return Pos{p.x, -p.z, p.y}
	case 4, 5, 6, 7:
		return transpose(Pos{-p.x, -p.y, p.z}, i-4)
	case 8, 9, 10, 11:
		return transpose(Pos{p.y, -p.x, p.z}, i-8)
	case 12, 13, 14, 15:
		return transpose(Pos{-p.y, p.x, p.z}, i-12)
	case 16, 17, 18, 19:
		return transpose(Pos{p.z, p.y, -p.x}, i-16)
	case 20, 21, 22, 23:
		return transpose(Pos{-p.z, p.y, p.x}, i-20)
	}
	panic("not supported - max 24")
}

func getDiff(p1, p2 Pos) Diff {
	return Diff{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}
}

func isEqual(p1, p2 Pos) bool {
	return p1.x == p2.x && p1.y == p2.y && p1.z == p2.z
}

func add(p Pos, d Diff) Pos {
	return Pos{p.x + d.dx, p.y + d.dy, p.z + d.dz}
}

func matches(unique, tested []Pos) (bool, []Pos, Pos) {
	for ui, u := range unique {
		if ui > len(unique)-11 {
			break
		}
		for ti, t := range tested {
			if ti > len(tested)-11 {
				break
			}
			diff := getDiff(u, t)

			matchedBeaconsIndex := []int{}

			for uui, uu := range unique {
				if uui > len(unique)-12+len(matchedBeaconsIndex) {
					break
				}
				for ti, tt := range tested {
					newPos := add(tt, diff)

					if isEqual(uu, newPos) {
						matchedBeaconsIndex = append(matchedBeaconsIndex, ti)
					}
				}
			}

			if len(matchedBeaconsIndex) >= 12 {
				missingBeacons := []Pos{}
				for i, beacon := range tested {
					skip := false
					for _, ii := range matchedBeaconsIndex {
						if i == ii {
							skip = true
						}
					}
					if skip {
						continue
					}
					missingBeacons = append(missingBeacons, add(beacon, diff))
				}
				scannerPosition := Pos{diff.dx, diff.dy, diff.dz}

				return true, missingBeacons, scannerPosition
			}
		}
	}
	return false, []Pos{}, Pos{}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func city(p1, p2 Pos) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y) + abs(p2.z-p1.z)
}

func main() {
	lines := utils.Input()

	scanners := []Scanner{}

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		if line[:3] == "---" {
			scanners = append(scanners, Scanner{false, []Pos{}})
		} else if len(line) > 0 {
			x, y, z := 0, 0, 0
			fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			scanners[len(scanners)-1].beacons = append(scanners[len(scanners)-1].beacons, Pos{x, y, z})
		}
	}

	scannerPositions := []Pos{}
	scannerPositions = append(scannerPositions, Pos{0, 0, 0})

	uniqueBeacons := []Pos{}
	uniqueBeacons = append(uniqueBeacons, scanners[0].beacons...)
	scanners[0].matched = true

	for found := true; found; {
		found = false
		for s := 1; s < len(scanners); s++ {
			if scanners[s].matched {
				continue
			}
			for transposeIteration := 0; transposeIteration < 24; transposeIteration++ {
				transposed := []Pos{}
				for _, beacon := range scanners[s].beacons {
					transposed = append(transposed, transpose(beacon, transposeIteration))
				}

				if matching, missingPoints, scannerPosition := matches(uniqueBeacons, transposed); matching {
					uniqueBeacons = append(uniqueBeacons, missingPoints...)
					scannerPositions = append(scannerPositions, scannerPosition)
					scanners[s].matched = true
					found = true
					break
				}
			}
		}
	}

	max := 0
	for i, s1 := range scannerPositions {
		for j, s2 := range scannerPositions {
			if i != j {
				len := city(s1, s2)
				if len > max {
					max = len
				}
			}
		}
	}

	fmt.Println(max)
}
