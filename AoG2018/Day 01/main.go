package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"
)

//go:embed test_input.txt
//go:embed 2test_input.txt
//go:embed input.txt
var filesystem embed.FS

func frequencySeenBefore(frequencies []int64, value int64) bool {
	for _, v := range frequencies {
		if v == value {
			return true
		}
	}

	return false
}

func processFrequencies(seenFrequencies []int64, frequency int64) ([]int64, int64) {
	// Instantiate variables. ParseInt() only wants to return an int64, but who cares about memory footprint?
	var value int64

	// Read file, log error if breakage
	f, err := filesystem.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		// ParseInt supports signed integers, so let's use those
		value, err = strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		frequency += value

		// The program is done once it's seen something twice, otherwise it should keep iterating forever
		if frequencySeenBefore(seenFrequencies, frequency) {
			fmt.Printf("Got duplicate number %d\n", frequency)
			os.Exit(0)
		} else {
			seenFrequencies = append(seenFrequencies, frequency)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return seenFrequencies, frequency
}

func main() {
	seenFrequencies := []int64{0}
	var frequency int64 = 0
	x := 0

	for true {
		seenFrequencies, frequency = processFrequencies(seenFrequencies, frequency)
		// Print the outcome of the first iteration to solve part 1
		if x == 0 {
			fmt.Printf("Final number %d\n", frequency)
		}
		x++
	}
}
