package day05

import (
	"adventOfCode"
	"fmt"
	"strings"
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
	var newPolymer string
	minLen := len(polymer)
	for _, letter := range "abcdefghijklmnopqrstuvwxwz" {
		newPolymer = strings.Replace(polymer, string(letter), "", -1)
		newPolymer = strings.Replace(newPolymer, strings.ToUpper(string(letter)), "", -1)
		newPolymer = reactPolymer(newPolymer)
		if len(newPolymer) < minLen {
			minLen = len(newPolymer)
		}
	}

	fmt.Printf("Answer: %d \n", minLen)

	return
}
