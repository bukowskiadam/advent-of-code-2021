package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
)

func toBinaryString(s string) string {
	binary := ""
	for i := 0; i < len(s); i++ {
		val, err := strconv.ParseInt(s[i:i+1], 16, 8)
		if err != nil {
			panic("Reading failed")
		}
		binary = binary + fmt.Sprintf("%04b", val)
	}
	return binary
}

func parseBin(s string) int {
	v, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic("parseBin error")
	}
	return int(v)
}

const HEADER_LEN = 3
const TYPE_ID_LEN = 3

func solve(s string) int {
	binary := toBinaryString(s)
	val, _ := next(binary)
	return val
}

func next(binary string) (int, int) {
	i := 0
	typeId := parseBin(binary[i+HEADER_LEN : i+HEADER_LEN+TYPE_ID_LEN])
	i = i + HEADER_LEN + TYPE_ID_LEN

	if typeId == 4 {
		literalBinary := ""
		for {
			prefixBit := binary[i]
			literalBinary = literalBinary + binary[i+1:i+5]
			i += 5
			if prefixBit == '0' {
				break
			}
		}
		literalValue := parseBin(literalBinary)
		return literalValue, i
	} else {
		lengthTypeId := parseBin(binary[i : i+1])
		i += 1
		values := []int{}
		if lengthTypeId == 0 {
			subpacketsLen := parseBin(binary[i : i+15])
			i += 15
			j := 0
			for j < subpacketsLen {
				val, processedBytes := next(binary[i+j:])
				values = append(values, val)
				j += processedBytes
			}
			i += j
		} else {
			numberOfSubpackets := parseBin(binary[i : i+11])
			i += 11

			j := 0
			for j < numberOfSubpackets {
				j++
				val, processedBytes := next(binary[i:])
				values = append(values, val)
				i += processedBytes
			}
		}

		switch typeId {
		case 0:
			sum := 0
			for _, x := range values {
				sum += x
			}
			return sum, i
		case 1:
			product := 1
			for _, x := range values {
				product *= x
			}
			return product, i
		case 2:
			min := math.MaxInt
			for _, x := range values {
				if x < min {
					min = x
				}
			}
			return min, i
		case 3:
			max := 0
			for _, x := range values {
				if x > max {
					max = x
				}
			}
			return max, i
		case 5:
			if values[0] > values[1] {
				return 1, i
			}
			return 0, i
		case 6:
			if values[0] < values[1] {
				return 1, i
			}
			return 0, i
		case 7:
			if values[0] == values[1] {
				return 1, i
			}
			return 0, i
		}
	}
	panic("something went wrong")
}

func main() {
	lines := utils.Input()

	for _, line := range lines {
		fmt.Println(solve(line))
	}
}
