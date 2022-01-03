package main

import (
	"advent/utils"
	"fmt"
	"strings"
	"unicode"
)

func isBig(cave string) bool {
	for _, r := range cave {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isSmall(cave string) bool {
	for _, r := range cave {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func canVisit(caves []string, next string) bool {
	m := map[string]int{}
	for _, c := range caves {
		if isSmall(c) && c != "start" {
			m[c] = m[c] + 1
		}
	}
	for _, cc := range m {
		if cc > 1 {
			return m[next] == 0
		}
	}
	return true
}

func main() {
	lines := utils.Input()

	caves := make(map[string][]string)

	for _, line := range lines {
		def := strings.Split(line, "-")
		caves[def[0]] = append(caves[def[0]], def[1])
		caves[def[1]] = append(caves[def[1]], def[0])
	}

	pathsCount := 0
	currentPath := []string{}

	var iteratePaths func(string)

	iteratePaths = func(cave string) {
		if cave == "end" {
			pathsCount++
			// fmt.Println(currentPath)
			return
		}
		currentPath = append(currentPath, cave)
		for _, target := range caves[cave] {
			if target == "start" {
				continue
			}
			if isSmall(target) {
				if canVisit(currentPath, target) {
					iteratePaths(target)
				}
			} else {
				iteratePaths(target)
			}
		}
		currentPath = currentPath[:len(currentPath)-1]
	}
	iteratePaths("start")
	fmt.Println(pathsCount)
}
