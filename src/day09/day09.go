package day09

import (
	"adventOfCode"
	"fmt"
)

func PartOne() {
	rawFileArr := adventOfCode.ReadInputFile("09", "input.txt")
	var numOfPlayers, maxPoints int
	fmt.Sscanf(rawFileArr[0], "%d players; last marble is worth %d points", &numOfPlayers, &maxPoints)
	for i := 0; i < maxPoints; i++ {
		fmt.Printf("Player [%d] places marble marked with #%d. \n", i%numOfPlayers, i)
	}

	fmt.Printf("Answer: %d \n", 42)
}

func PartTwo() {
	fmt.Printf("Answer: %d \n", 42)
}
