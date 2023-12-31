package main

import (
	"bufio"
	"embed"
	"io/fs"
	"log"
)

//go:embed test_input.txt
//go:embed input.txt
var filesystem embed.FS

func readInputFile(fileName string) fs.File {
	// Read file, log error if something breaks
	f, err := filesystem.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func getInputFromInputFile(fileName string) [][]string {
	var lines [][]string
	scanner := bufio.NewScanner(readInputFile(fileName))
	for scanner.Scan() {
		var thisLine = []string{scanner.Text()}
		lines = append(lines, thisLine)
	}
	return lines
}
