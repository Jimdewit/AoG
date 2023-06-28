package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func castToInt(castTarget string) int {
	convertedInt, err := strconv.Atoi(castTarget)
	if err != nil {
		panic(err)
	}
	return convertedInt
}

type coords struct {
	x int
	y int
}

type claim struct {
	elfNumber          int
	leftUpperBoundary  coords
	leftLowerBoundary  coords
	rightUpperBoundary coords
	rightLowerBoundary coords
}

func fillMap(singleClaim claim, mappedClaims map[coords][]int) map[coords][]int {
	for y := singleClaim.leftUpperBoundary.y; y < singleClaim.leftLowerBoundary.y; y++ {
		for x := singleClaim.leftUpperBoundary.x; x < singleClaim.rightUpperBoundary.x; x++ {
			claimCoordinates := coords{x, y}
			mappedClaims[claimCoordinates] = append(mappedClaims[claimCoordinates], singleClaim.elfNumber)
		}
	}
	return mappedClaims
}

func parseClaim(claimContents []string) claim {
	splitClaim := strings.Split(claimContents[0], " ")
	elfNumber := castToInt(splitClaim[0][1:])
	claimWidth := castToInt(strings.Split(splitClaim[len(splitClaim)-1], "x")[0])
	claimHeight := castToInt(strings.Split(splitClaim[len(splitClaim)-1], "x")[1])
	leftUpperX := castToInt(strings.Split(splitClaim[2], ",")[0])
	leftUpperY := castToInt(strings.Split(splitClaim[2], ",")[1][:len(strings.Split(splitClaim[2], ",")[1])-1])

	parsedClaim := claim{elfNumber,
		coords{leftUpperX, leftUpperY},
		coords{leftUpperX, leftUpperY + claimHeight},
		coords{leftUpperX + claimWidth, leftUpperY},
		coords{leftUpperX + claimWidth, leftUpperY + claimHeight}}

	return parsedClaim
}

func processClaimList(claimList [][]string) map[coords][]int {
	mappedClaims := make(map[coords][]int)
	for e := 0; e < len(claimList); e++ {
		claimContents := claimList[e]
		parsedClaim := parseClaim(claimContents)
		mappedClaims = fillMap(parsedClaim, mappedClaims)
	}

	return mappedClaims
}

func countDoubleClaims(mappedClaims map[coords][]int) int {
	counter := 0
	for k := range mappedClaims {
		if len(mappedClaims[k]) > 1 {
			counter++
		}
	}
	return counter
}

func checkForOverlaps(claimList [][]string, mappedClaims map[coords][]int, out chan<- int) {
	for e := 0; e < len(claimList); e++ {
		overlaps := 0
		claimContents := claimList[e]
		parsedClaim := parseClaim(claimContents)
		for _, v := range mappedClaims {
			for _, claimNumber := range v {
				if parsedClaim.elfNumber == claimNumber && len(v) > 1 {
					overlaps++
				} else {
					continue
				}
			}
		}
		if overlaps == 0 {
			out <- parsedClaim.elfNumber
		}
	}
}

func main() {
	inputFilename := "input.txt"
	claimList := getInputFromInputFile(inputFilename)
	mappedClaims := processClaimList(claimList)
	doubleClaims := countDoubleClaims(mappedClaims)
	fmt.Printf("Double claims %d\n", doubleClaims)

	start := time.Now().UnixMilli()
	chanRes := make(chan int, 1)
	go checkForOverlaps(claimList, mappedClaims, chanRes)
	uniqueClaim := <-chanRes
	fmt.Printf("Unique claim %d\n", uniqueClaim)
	finish := time.Now().UnixMilli()
	timeTaken := finish - start
	fmt.Printf("Time taken %d milliseconds", timeTaken)
}
