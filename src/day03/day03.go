package day03

import (
	"adventOfCode"
	"utils"
	"strings"
	"fmt"
)

type Claim struct {
	id 		string
	xOffset	int
	yOffset int
	xAmount	int
	yAmount int
}

type Point struct {
    x  int
    y  int
}

func isRawClaimDelimiter(r rune) bool {
	return r == ',' || r == 'x' || r == ':' || r == '#' || r == '@'
}

func parseClaim(rawClaim string) Claim {
	claimParts := strings.FieldsFunc(rawClaim, isRawClaimDelimiter)
	return Claim{
		claimParts[0],
		utils.StrToInt(claimParts[1][1:]),
		utils.StrToInt(claimParts[2]),
		utils.StrToInt(claimParts[3][1:]),
		utils.StrToInt(claimParts[4]),
	}
}

func PartOne()  {
	linesOfClaims := adventOfCode.ReadInputFile("03", "input.txt")

	SquaresClaims := make(map[Point]int)
	squaresClaimedTwiceCount := 0

	for _, rawClaim := range linesOfClaims {

		claim := parseClaim(rawClaim)

		// Tally up the number of claims for each square
		for y := claim.yOffset + 1; y <= claim.yOffset + claim.yAmount; y++ {
			for x := claim.xOffset + 1; x <= claim.xOffset + claim.xAmount; x++  {
				if _, ok := SquaresClaims[Point{x, y}]; ok {
					SquaresClaims[Point{x, y}] += 1
				} else {
					SquaresClaims[Point{x, y}] = 1
				}
			}
		}
	}

	// Look for two or more claims
	for _, numberOfClaims := range SquaresClaims {
		if(numberOfClaims >= 2){
			squaresClaimedTwiceCount += 1
		}

	}

	fmt.Printf("Answer: %d \n", squaresClaimedTwiceCount)

}

func PartTwo()  {
	linesOfClaims := adventOfCode.ReadInputFile("03", "input.txt")

	IDsOfSingleClaims := make([]string, len(linesOfClaims))
	SquaresClaims := make(map[Point]int)

	for claimNumber, rawClaim := range linesOfClaims {

		claim := parseClaim(rawClaim)

		// Store the id till it's claimed again
		IDsOfSingleClaims[claimNumber] = claim.id

		// Tally up the number of claims for each square
		for y := claim.yOffset + 1; y <= claim.yOffset + claim.yAmount; y++ {
			for x := claim.xOffset + 1; x <= claim.xOffset + claim.xAmount; x++  {
				if _, ok := SquaresClaims[Point{x, y}]; ok {
					// Square was previously claimed, remove stored Ids
					prevClaim := SquaresClaims[Point{x, y}]
					IDsOfSingleClaims[prevClaim] = ""
					IDsOfSingleClaims[claimNumber] = ""

				} else {
					SquaresClaims[Point{x, y}] = claimNumber
				}
			}
		}
	}

	for _, claimID := range IDsOfSingleClaims {
		if(claimID != ""){
			fmt.Printf("Answer: %s \n", claimID)
			return
		}
	}
}