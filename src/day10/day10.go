package day10

import (
	"adventOfCode"
	"fmt"
)

type Point struct {
	X int
	Y int
}

type point struct {
	X  int
	Y  int
	dX int
	dY int
}

func PartOne() {
	rawFileArr := adventOfCode.ReadInputFile("10", "input.txt")
	points := make(map[Point]bool)
	for _, line := range rawFileArr {
		var p point
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &p.X, &p.Y, &p.dX, &p.dY)
		points[Point{p.X, p.Y}] = true
	}

	fmt.Printf("Answer: %d \n", 42)
}

func PartTwo() {
	fmt.Printf("Answer: %d \n", 42)
}
