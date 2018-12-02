package day02

import (
	"adventOfCode"
	"fmt"
)

func PartOne()  {
	linesOfIDs := adventOfCode.ReadInputFile("02", "input.txt")

	doubleCount := 0
	tripleCount := 0

	for _, ID:= range linesOfIDs {
		LetterCountInID := make(map[int32]int)
		hasTwiceRepeatingLetter := false
		hasThriceRepeatingLetter := false

		for _, letter := range ID{
			if _, ok := LetterCountInID[letter]; ok {
				LetterCountInID[letter] += 1
			} else {
				LetterCountInID[letter] = 1
			}
		}

		for _, frequency := range LetterCountInID {
			if frequency == 2{
				hasTwiceRepeatingLetter = true
			}
			if frequency == 3{
				hasThriceRepeatingLetter = true
			}
		}

		if(hasTwiceRepeatingLetter){
			doubleCount += 1
		}
		if(hasThriceRepeatingLetter){
			tripleCount += 1
		}
	}

	fmt.Printf("Answer: %d \n", doubleCount * tripleCount)

}

func PartTwo()  {
	linesOfIDs := adventOfCode.ReadInputFile("02", "input.txt")
	WordSet := make(map[string]string)

	for _, ID := range(linesOfIDs){
		for i, _ := range ID {
			wordWithoutOneChar := ID[0:i] + ID[(i+1):]
			if _, ok := WordSet[wordWithoutOneChar]; ok {
				if(WordSet[wordWithoutOneChar] != ID){
					fmt.Printf("Answer: %s \n", wordWithoutOneChar)
					return
				}
			}
			WordSet[wordWithoutOneChar] = ID
		}
	}

}