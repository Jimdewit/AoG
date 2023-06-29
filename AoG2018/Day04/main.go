package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type sleepCounter struct {
	minuteMap map[int]int
}

type guardMap struct {
	guard map[string]sleepCounter
}

func sleepCounterMaker() sleepCounter {
	newMinuteMap := make(map[int]int)
	for x := 0; x < 60; x++ {
		newMinuteMap[x] = 0
	}
	newSleepCounter := sleepCounter{newMinuteMap}
	return newSleepCounter
}

func processNaps(guardId string, minuteAsleep int, minuteAwake int, allGuards guardMap) guardMap {
	_, ok := allGuards.guard[guardId]
	if !ok {
		allGuards.guard[guardId] = sleepCounterMaker()
	}
	for x := minuteAsleep; x < minuteAwake; x++ {
		allGuards.guard[guardId].minuteMap[x] += 1
	}

	return allGuards
}

func findMostSleepyMinute(guard sleepCounter) int {
	keys := make([]int, 0, len(guard.minuteMap))

	for key := range guard.minuteMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return guard.minuteMap[keys[i]] > guard.minuteMap[keys[j]]
	})

	return keys[0]
}

func findMostSleepyGuard(allGuards guardMap) {
	var mostSleepyGuard = "0"
	var mostMinutesSlept = 0
	for guardId, guardSleepCounter := range allGuards.guard {
		totalMinutesSlept := 0
		for x := range guardSleepCounter.minuteMap {
			totalMinutesSlept += guardSleepCounter.minuteMap[x]
		}
		if totalMinutesSlept > mostMinutesSlept {
			mostMinutesSlept = totalMinutesSlept
			mostSleepyGuard = guardId
		}
	}
	fmt.Printf("Most sleepy guard: %v\n", mostSleepyGuard)
	calculator, _ := strconv.Atoi(mostSleepyGuard)
	fmt.Println(findMostSleepyMinute(allGuards.guard[mostSleepyGuard]) * calculator)
}

func findMostOftenSleptMinute(allGuards guardMap) {
	var mostOftenSleepingGuard = "0"
	var mostTimesSlept = 0
	var mostSleepyMinute = 0
	for guardId, guardSleepCounter := range allGuards.guard {
		for x := range guardSleepCounter.minuteMap {
			if guardSleepCounter.minuteMap[x] > mostTimesSlept {
				mostTimesSlept = guardSleepCounter.minuteMap[x]
				mostOftenSleepingGuard = guardId
				mostSleepyMinute = x
			}
		}
	}
	fmt.Printf("Most sleepy guard: %v\n", mostOftenSleepingGuard)
	calculator, _ := strconv.Atoi(mostOftenSleepingGuard)
	fmt.Println(mostSleepyMinute * calculator)
}

func doSolve(input []string) {
	// Sort input to properly order the data
	sort.Strings(input)

	// Initialize map of guards
	allGuards := guardMap{map[string]sleepCounter{}}

	// Declare regexes for matching
	var fallsAsleepRegex = regexp.MustCompile(`falls asleep`)
	var wakesUpRegex = regexp.MustCompile(`wakes up`)
	var beginsShiftRegex = regexp.MustCompile(`begins shift`)

	// Declare globals for usage during iteration
	var guardId string
	var minuteAsleep int
	var minuteAwake int

	minuteFindRegex := regexp.MustCompile(`:(?P<minute>\d{2})`)
	guardFindRegex := regexp.MustCompile(`#(?P<guardId>\d+)`)

	for l := 0; l < len(input); l++ {
		line := input[l]
		if beginsShiftRegex.MatchString(line) {
			guardId = guardFindRegex.FindStringSubmatch(line)[1]
		} else if fallsAsleepRegex.MatchString(line) {
			minuteAsleep, _ = strconv.Atoi(minuteFindRegex.FindStringSubmatch(line)[1])
		} else if wakesUpRegex.MatchString(line) {
			minuteAwake, _ = strconv.Atoi(minuteFindRegex.FindStringSubmatch(line)[1])
			allGuards = processNaps(guardId, minuteAsleep, minuteAwake, allGuards)
		}
	}
	findMostSleepyGuard(allGuards)
	findMostOftenSleptMinute(allGuards)
}

func main() {
	inputFilename := "input.txt"
	input := getInputFromInputFile(inputFilename)
	doSolve(input)
}
