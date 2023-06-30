package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type element struct {
	charRune        rune
	charIsUpperCase bool
}

func toRune(char string) rune {
	characterRune, _ := utf8.DecodeRuneInString(char)
	return characterRune
}

func elementsAreReactive(first element, second element) bool {
	if unicode.ToLower(first.charRune) != unicode.ToLower(second.charRune) {
		return false
	} else if first.charIsUpperCase == second.charIsUpperCase {
		return false
	} else {
		return true
	}
}

func solvePartOneAndPrepareForTwo(input []string) int {
	charIndex := 0
	for true {
		if charIndex >= len(input)-1 {
			break
		}
		var left, right element = element{toRune(input[charIndex]), unicode.IsUpper(toRune(input[charIndex]))},
			element{toRune(input[charIndex+1]), unicode.IsUpper(toRune(input[charIndex+1]))}
		if elementsAreReactive(left, right) {
			input = append(input[:charIndex], input[charIndex+2:]...)
			// Prevent ugly edge case the silly way
			if charIndex > 0 {
				charIndex -= 1
			}
		} else {
			charIndex += 1
		}
	}

	return len(input)
}

func solvePartTwo(inputFilename string) int {
	shortestPolymer := 100000
	var elementsToCheck = make([]rune, 0)
	for r := 'a'; r < 'z'; r++ {
		elementsToCheck = append(elementsToCheck, r)
	}

	for _, r := range elementsToCheck {
		input := strings.Split(getInputFromInputFile(inputFilename)[0], "")
		charIndex := 0
		for true {
			if charIndex > len(input)-1 {
				break
			}
			var charToCheck = element{toRune(input[charIndex]), unicode.IsUpper(toRune(input[charIndex]))}
			if unicode.ToLower(charToCheck.charRune) == r {
				input = append(input[:charIndex], input[charIndex+1:]...)
			} else {
				charIndex += 1
			}

		}
		charIndex = 0
		for true {
			if charIndex >= len(input)-1 {
				break
			}
			var left, right element = element{toRune(input[charIndex]), unicode.IsUpper(toRune(input[charIndex]))},
				element{toRune(input[charIndex+1]), unicode.IsUpper(toRune(input[charIndex+1]))}
			if elementsAreReactive(left, right) {
				input = append(input[:charIndex], input[charIndex+2:]...)
				// Prevent ugly edge case the silly way
				if charIndex > 0 {
					charIndex -= 1
				}
			} else {
				charIndex += 1
			}
		}
		if len(input) < shortestPolymer {
			shortestPolymer = len(input)
		}
	}

	return shortestPolymer
}

func main() {
	inputFilename := "input.txt"
	input := strings.Split(getInputFromInputFile(inputFilename)[0], "")
	first := solvePartOneAndPrepareForTwo(input)
	fmt.Printf("Length after initial reactions: %d\n", first)
	second := solvePartTwo(inputFilename)
	fmt.Printf("Shortest possible: %d", second)
}
