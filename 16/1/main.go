package main

import (
	"advent/utils"
	"fmt"
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
	versionsSum := 0

	for i := 0; i < len(binary)-(HEADER_LEN+TYPE_ID_LEN); {
		version := parseBin(binary[i : i+HEADER_LEN])
		typeId := parseBin(binary[i+HEADER_LEN : i+HEADER_LEN+TYPE_ID_LEN])
		i = i + HEADER_LEN + TYPE_ID_LEN

		versionsSum += version
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
			// literalValue := parseBin(literalBinary)
			// fmt.Println("Literal", literalValue)
		} else {
			lengthTypeId := parseBin(binary[i : i+1])
			i += 1
			if lengthTypeId == 0 {
				// lenghtInBits := parseBin(binary[i : i+15])
				i += 15
			} else {
				// numberOfSubpackets := parseBin(binary[i : i+11])
				i += 11
			}
		}
	}

	return versionsSum
}

func main() {
	lines, _ := utils.ReadLines(utils.ReadFileFromArgs())

	for _, line := range lines {
		fmt.Println(solve(line))
	}
}
