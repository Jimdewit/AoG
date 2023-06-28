package main

import (
	"fmt"
	"os"
	"strings"
)

func countOccurrences(wordToCheck string, desiredOccurrences int) int {

	for positionToCheck := 0; positionToCheck < len(wordToCheck); positionToCheck++ {
		letter := wordToCheck[positionToCheck : positionToCheck+1]

		if strings.Count(wordToCheck, letter) == desiredOccurrences {
			return 1
		}
	}
	return 0
}

func compareWordToOthers(wordToCheck string, allWords []string) string {
	for wordPosition := 0; wordPosition < len(allWords); wordPosition++ {
		differences := 0
		letters := ""
		comparisonWord := allWords[wordPosition]

		for positionToCheck := 0; positionToCheck < len(wordToCheck); positionToCheck++ {
			letter := wordToCheck[positionToCheck : positionToCheck+1]

			if comparisonWord[positionToCheck:positionToCheck+1] != letter {
				differences += 1
			} else {
				letters += letter
			}
		}
		if differences == 1 {
			return letters
		}
	}
	return ""
}

func main() {
	inputFilename := "input.txt"
	input := getInputFromInputFile(inputFilename)
	countTwos, countThrees := 0, 0
	// part one-1: count number of words that have a letter occurring twice
	for x := 0; x < len(input); x++ {
		wordToCheck := input[x]
		countTwos += countOccurrences(wordToCheck, 2)
	}
	// part one-2: count number of words that have a letter occurring thrice
	for x := 0; x < len(input); x++ {
		wordToCheck := input[x]
		countThrees += countOccurrences(wordToCheck, 3)
	}
	fmt.Printf("\nMultiplying %d by %d yields %d", countTwos, countThrees, countTwos*countThrees)

	// part two: find two words that have exactly only differing character
	var commonLetters string
	for x := 0; x < len(input); x++ {
		wordToCheck := input[x]
		commonLetters = compareWordToOthers(wordToCheck, input)
		if len(commonLetters) > 0 {
			fmt.Printf("\nFound common letters %v", commonLetters)
			os.Exit(0)
		}
	}
}
