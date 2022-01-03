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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
			fmt.Println(currentPath)
			return
		}
		currentPath = append(currentPath, cave)
		for _, target := range caves[cave] {
			if isSmall(target) {
				if !contains(currentPath, target) {
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
