package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFileFromArgs() string {
	if len(os.Args) < 2 {
		panic("Provide an input file")
	}
	return os.Args[1]
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Input() []string {
	lines, _ := ReadLines(ReadFileFromArgs())

	return lines
}

func MapToNumbers(numbers []string) []int {
	vsm := make([]int, len(numbers))
	for i, v := range numbers {
		vsm[i], _ = strconv.Atoi(v)
	}
	return vsm
}

func ReadCommaSeparatedInts(line string) []int {
	s := strings.Split(line, ",")
	values := make([]int, len(s))
	for i, val := range s {
		values[i], _ = strconv.Atoi(val)
	}
	return values
}
