package main

import (
	"advent/utils"
	"fmt"
	"sort"
	"strings"
)

func sortDigit(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func includes(s1, s2 string) bool {
	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}
	return true
}

func main() {
	defer utils.MeasureExecutionTime("main")()
	lines := utils.Input()
	sum := 0

	for _, line := range lines {
		left, right := func(s string) ([]string, []string) {
			x := strings.Split(line, "|")
			l := strings.Fields(x[0])
			r := strings.Fields(x[1])

			for i, v := range l {
				l[i] = sortDigit(v)
			}
			for i, v := range r {
				r[i] = sortDigit(v)
			}

			return l, r
		}(line)

		dict := make(map[int]string, 10)

		for _, digit := range left {
			switch len(digit) {
			case 2:
				dict[1] = digit
			case 3:
				dict[7] = digit
			case 4:
				dict[4] = digit
			case 7:
				dict[8] = digit
			}
		}

		for _, digit := range left {
			if len(digit) == 6 {
				if includes(digit, dict[4]) {
					dict[9] = digit
				}
			}
		}

		for _, digit := range left {
			if len(digit) == 5 {
				if includes(digit, dict[1]) {
					dict[3] = digit
				} else if includes(dict[9], digit) {
					dict[5] = digit
				} else {
					dict[2] = digit
				}
			}
		}

		for _, digit := range left {
			if len(digit) == 6 {
				if !includes(digit, dict[4]) {
					if includes(digit, dict[5]) {
						dict[6] = digit
					} else {
						dict[0] = digit
					}
				}
			}
		}

		fmt.Println(dict, right)

		invertedDict := make(map[string]int, 10)
		for num, digit := range dict {
			invertedDict[digit] = num
		}

		val := 0
		for _, d := range right {
			val = val*10 + invertedDict[d]
		}

		fmt.Println(val)

		sum += val
	}

	fmt.Println(sum)
}
