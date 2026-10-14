package main

import (
	"os"      // for reading from files
	"strconv" // our old friend string conversion package
	"strings" // for working with strings
)

//ReadElectoralVotes processes the number of electoral votes for each state. Takes a filename string as input and outputs a map of string states to uint of electoral votes

func ReadElectoralVotes(filename string) map[string]uint {
	electoralVotes := make(map[string]uint)

	fileContents, err := os.ReadFile(filename)
	//fileContents is a slice of bytes
	Check(err)

	giantString := string(fileContents)
	lines := strings.Split(giantString, "\n")

	for _, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")
		stateName := lineElements[0]
		numVotes, err := strconv.Atoi(lineElements[1])
		Check(err)

		electoralVotes[stateName] = uint(numVotes)
	}

	return electoralVotes
}

func ReadPollingData(filename string) map[string]float64 {
	candidate1Percentage := make(map[string]float64)

	fileContents, err := os.ReadFile(filename)
	//fileContents is a slice of bytes
	Check(err)

	giantString := string(fileContents)
	lines := strings.Split(giantString, "\n")

	for _, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")
		stateName := lineElements[0]
		percentage1, err := strconv.ParseFloat(lineElements[1], 64)
		Check(err)

		candidate1Percentage[stateName] = percentage1 / 100.0
	}

	return candidate1Percentage
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
