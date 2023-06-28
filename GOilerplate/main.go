package main

import (
	"fmt"
)

func main() {
	inputFilename := "input.txt"
	input := getInputFromInputFile(inputFilename)
	fmt.Print(input)
}
