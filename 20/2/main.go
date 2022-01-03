package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

const ITERATIONS = 50

func p(s []string) {
	for _, i := range s {
		fmt.Println(i)
	}
	fmt.Print("\n")
}

func extendEmpty(image []string) []string {
	for lineIndex, line := range image {
		image[lineIndex] = "." + line + "."
	}
	length := len(image[0])
	image = append([]string{strings.Repeat(".", length)}, image...)
	image = append(image, strings.Repeat(".", length))

	return image
}

func extend(image []string) []string {
	char := image[0][:1]
	for lineIndex, line := range image {
		image[lineIndex] = char + line + char
	}
	length := len(image[0])
	image = append([]string{strings.Repeat(char, length)}, image...)
	image = append(image, strings.Repeat(char, length))

	return image
}

func shrink(image []string, frame int) []string {
	n := make([]string, len(image)-2*frame)

	for i := frame; i < len(image)-frame; i++ {
		n[i-frame] = image[i][frame : len(image[i])-frame]
	}

	return n
}

func makeEmpty(lines, cols int) []string {
	empty := make([]string, lines)
	for j := 0; j < lines; j++ {
		empty[j] = strings.Repeat(".", cols)
	}
	return empty
}

func transform(image []string, dict string) []string {
	lines := len(image)
	cols := len(image[0])
	newImage := makeEmpty(lines, cols)

	for i := 0; i < lines-2; i++ {
		line := "."
		for j := 0; j < cols-2; j++ {
			n := 0
			for z := 0; z < 9; z++ {
				val := 0
				if image[i+z/3][j+z%3] == '#' {
					val = 1
				}
				n = 2*n + val
			}
			line += string(dict[n])
		}
		line += "."
		newImage[i+1] = line
	}
	return newImage
}

func main() {
	lines := utils.Input()

	dict := lines[0]
	image := lines[2:]

	fmt.Println(len(dict), len(image))

	image = extendEmpty(image)
	image = extendEmpty(image)
	for i := 0; i < ITERATIONS; i++ {
		image = extend(image)
		// p(image)
		image = transform(image, dict)
		image = shrink(image, 1)
		image = extend(image)
		// p(image)
	}

	// p(image)
	c := 0
	for _, line := range image {
		for _, char := range line {
			if char == '#' {
				c += 1
			}
		}
	}
	fmt.Println(c)
}
