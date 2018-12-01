package day01

import (
	"adventOfCode"
	"fmt"
	"utils"
)

func PartOne() {
	linesOfFile := adventOfCode.ReadInputFile("01", "input.txt")

	sum := 0

	for _, numberAsString := range linesOfFile {
		number := utils.StrToInt(numberAsString)
		sum += number

	}

	fmt.Printf("Answer: %d \n", sum)
}

func PartTwo() {
	linesOfFile := adventOfCode.ReadInputFile("01", "input.txt")

	var IntSet map[int]bool
	IntSet = make(map[int]bool)

	sum := 0

	for true {
		for _, numberAsString := range linesOfFile {
			number := utils.StrToInt(numberAsString)
			sum += number
			// Check if item is in dict
			if _, ok := IntSet[sum]; ok {
				fmt.Printf("Answer: %d \n", sum)
				return
			}
			IntSet[sum] = true
		}
	}

	fmt.Printf("Answer: %d \n", sum)
}
