package main

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed test_input.txt
//go:embed input.txt
var filesystem embed.FS

func getInput() fs.File {
	// Read file, log error if breakage
	f, err := filesystem.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func main() {
	scanner := bufio.NewScanner(getInput())
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
