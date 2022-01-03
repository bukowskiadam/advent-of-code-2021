package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

type SN struct {
	number int
	left   *SN
	right  *SN
}

func (n *SN) isNumber() bool {
	return n.left == nil && n.right == nil
}

func (n *SN) isNumbers() bool {
	return !n.isNumber() && n.left.isNumber() && n.right.isNumber()
}

func findComma(s string) int {
	depth := 0

	for i, char := range s {
		if char == '[' {
			depth++
			continue
		}
		if char == ']' {
			depth--
		}
		if char == ',' && depth == 0 {
			return i
		}
	}
	panic("no comma")
}

func parseNumber(s string) *SN {
	if len(s) == 1 {
		num, _ := strconv.Atoi(s)
		return &SN{number: num}
	}

	ss := s[1 : len(s)-1]
	commaPos := findComma(ss)
	left := parseNumber(ss[:commaPos])
	right := parseNumber(ss[commaPos+1:])

	return &SN{left: left, right: right}
}

func printNumber(num *SN) {
	if num.left == nil && num.right == nil {
		fmt.Print(num.number)
	} else {
		fmt.Print("[")
		printNumber(num.left)
		fmt.Print(",")
		printNumber(num.right)
		fmt.Print("]")
	}
}

func p(num *SN) {
	printNumber(num)
	fmt.Print("\n")
}

func add(left, right *SN) *SN {
	return &SN{
		left:  left,
		right: right,
	}
}

func reduceSplits(num *SN) bool {
	if num.isNumber() {
		if num.number >= 10 {
			// fmt.Print("split: ")
			// p(num)
			leftNum := num.number / 2
			rightNum := num.number / 2
			if num.number%2 == 1 {
				rightNum += 1
			}
			num.left = &SN{number: leftNum}
			num.right = &SN{number: rightNum}
			num.number = 0
			return true
		}
		return false
	}

	if reduceSplits(num.left) {
		return true
	}

	if reduceSplits(num.right) {
		return true
	}

	return false
}

func reduceExplodes(num *SN, depth int) (bool, int, int) {
	if num.isNumber() {
		return false, -1, -1
	}
	if num.isNumbers() && depth >= 4 {
		// fmt.Print("explode: ")
		// p(num)
		l := num.left.number
		r := num.right.number
		num.number = 0
		num.left = nil
		num.right = nil
		return true, l, r
	}

	if reduced, exLeft, exRight := reduceExplodes(num.left, depth+1); reduced {
		if exRight > 0 {
			if num.right.isNumber() {
				num.right.number += exRight
				exRight = -1
			} else {
				z := num.right
				for !z.isNumber() {
					z = z.left
				}
				z.number += exRight
				exRight = -1
			}
		}
		return true, exLeft, exRight
	}

	if reduced, exLeft, exRight := reduceExplodes(num.right, depth+1); reduced {
		if exLeft > 0 {
			if num.left.isNumber() {
				num.left.number += exLeft
				exLeft = -1
			} else {
				z := num.left
				for !z.isNumber() {
					z = z.right
				}
				z.number += exLeft
				exLeft = -1
			}
		}
		return true, exLeft, exRight
	}

	return false, -1, -1
}

func magnitude(num *SN) int {
	if num.isNumber() {
		return num.number
	}

	return 3*magnitude(num.left) + 2*magnitude(num.right)
}

func main() {
	lines := utils.Input()

	var number *SN

	for _, line := range lines {
		if number == nil {
			number = parseNumber(line)
		} else {
			number = add(number, parseNumber(line))
			// p(number)
			// fmt.Println("reduced:")
			for reduced := true; reduced; {
				reduced, _, _ = reduceExplodes(number, 0)
				if !reduced {
					reduced = reduceSplits(number)
				}
				// p(number)
			}
			// fmt.Print("\n")
		}
	}

	fmt.Println(magnitude(number))
}
