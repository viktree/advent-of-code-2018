package day05

import (
	"adventOfCode"
	"fmt"
)

func reactPolymer(polymer string) string {
	var newPolymer []byte
	for i := range polymer {
		molecule := polymer[i]
		if len(newPolymer) == 0 {
			newPolymer = append(newPolymer, molecule)
		} else {
			prevMolecule := newPolymer[len(newPolymer)-1]
			if prevMolecule^molecule != 32 {
				newPolymer = append(newPolymer, molecule)
			} else {
				newPolymer = newPolymer[0 : len(newPolymer)-1]
			}
		}
	}
	return string(newPolymer)
}

func PartOne() {

	polymer := adventOfCode.ReadInputFile("05", "input.txt")[0]
	newPolymer := reactPolymer(polymer)
	fmt.Printf("Answer: %d \n", len(newPolymer))

	return
}

func PartTwo() {

	polymer := adventOfCode.ReadInputFile("05", "input.txt")[0]
	// polymer = "dabAcCaCBAcCcaDA"
	numberOfBonds := make(map[string]int)
	var newPolymer []byte
	for i := range polymer {
		molecule := polymer[i]
		if len(newPolymer) == 0 {
			newPolymer = append(newPolymer, molecule)
		} else {
			prevMolecule := newPolymer[len(newPolymer)-1]
			if prevMolecule^molecule != 32 {
				newPolymer = append(newPolymer, molecule)
			} else {
				newPolymer = newPolymer[0 : len(newPolymer)-1]
				if _, ok := numberOfBonds[string(prevMolecule)]; ok {
					numberOfBonds[string(prevMolecule)] += 2
					numberOfBonds[string(molecule)] += 2

				} else {
					numberOfBonds[string(prevMolecule)] = 2
					numberOfBonds[string(molecule)] = 2
				}
			}
		}
	}
	answer := 0
	for i := range numberOfBonds {
		if numberOfBonds[i] > answer {
			answer = numberOfBonds[i]
		}
	}

	fmt.Printf("Answer: %d \n", answer)

	return
}
