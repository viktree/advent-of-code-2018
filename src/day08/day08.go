package day08

import (
	"adventOfCode"
	"fmt"
	"strings"
	"utils"
)

func printTree(tree []string, sumOfMetadata int) ([]string, int) {
	numberOfSubtrees := utils.StrToInt(tree[0])
	amountOfMetadata := utils.StrToInt(tree[1])
	tree = tree[2:]
	for numberOfSubtrees > 0 {
		tree, sumOfMetadata = printTree(tree, sumOfMetadata)
		numberOfSubtrees--
	}
	if numberOfSubtrees == 0 {
		for i := 0; i < amountOfMetadata; i++ {
			sumOfMetadata += utils.StrToInt(tree[i])
		}
	}
	return tree[amountOfMetadata:], sumOfMetadata
}

func computeValue(tree []string) (int, []string) {
	var value int
	var stack []int

	numberOfSubtrees := utils.StrToInt(tree[0])
	amountOfMetadata := utils.StrToInt(tree[1])

	if numberOfSubtrees == 0 {
		tree = tree[2:]
		for i := 0; i < amountOfMetadata; i++ {
			value += utils.StrToInt(tree[i])
		}
		fmt.Println("No Children", value)
		return value, tree[amountOfMetadata:]
	}
	tempTree := tree[2 : len(tree)-amountOfMetadata]
	var valueOfSubtree int
	fmt.Printf("Has %d Children: \n", numberOfSubtrees)
	i, origNumOfTrees := 1, numberOfSubtrees
	for numberOfSubtrees > 0 {
		valueOfSubtree, tempTree = computeValue(tempTree)
		stack = append(stack, valueOfSubtree)
		fmt.Printf("Value of child %d/%d : %d on %v\n", i, origNumOfTrees, valueOfSubtree, stack)
		numberOfSubtrees--
		i++
	}
	metadataPortion := tree[len(tree)-amountOfMetadata : len(tree)]
	for i := 0; i < amountOfMetadata; i++ {
		ref := utils.StrToInt(metadataPortion[i])
		if ref > 0 && ref <= len(stack) {
			value += stack[ref-1]
		}
		fmt.Println("ref", ref, "stack", stack, "value", value, ref > 0, ref <= len(stack))
	}
	return value, tree
}

func PartOne() {
	rawInput := adventOfCode.ReadInputFile("08", "input.txt")[0]
	tree := strings.Split(rawInput, " ")
	_, sum := printTree(tree, 0)

	fmt.Printf("Answer: %d \n", sum)
}

func PartTwo() {
	rawInput := adventOfCode.ReadInputFile("08", "input.txt")[0]
	// rawInput = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	tree := strings.Split(rawInput, " ")
	answer, _ := computeValue(tree)
	fmt.Printf("Answer: %d \n", answer)
}
