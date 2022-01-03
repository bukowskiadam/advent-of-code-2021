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

func matches(unique, tested []Pos) (bool, []Pos) {
	for _, u := range unique {
		for _, t := range tested {
			diff := getDiff(u, t)

			matchedBeaconsIndex := []int{}

			for _, uu := range unique {
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

				return true, missingBeacons
			}
		}
	}
	return false, []Pos{}
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

				if matching, missingPoints := matches(uniqueBeacons, transposed); matching {
					uniqueBeacons = append(uniqueBeacons, missingPoints...)
					scanners[s].matched = true
					found = true
					break
				}
			}
		}
	}

	fmt.Println(len(uniqueBeacons))
}
