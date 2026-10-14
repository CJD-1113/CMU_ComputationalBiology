package main

import (
	"fmt"      //formats output in the terminal
	"io"       //read and write to files
	"log"      //logs file error
	"net/http" //for accessing URLS
)

func main() {
	fmt.Println("Clumps")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Recieved non-OK status: %v", resp.Status)
	}

	genomeSymbols, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("The number of nucleatides in genome is", len(genomeSymbols))

	defer resp.Body.Close() //happens at end of func main due to defer

	EcoliGenome := string(genomeSymbols)

	k := 9
	L := 500
	t := 3
	clumps := FindClumpsFaster(EcoliGenome, k, L, t)

	fmt.Println("Found", len(clumps), "total patterns occuring as clumps")
}

func FindClumpsFaster(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	foundPatterns := make(map[string]bool)
	n := len(text)

	firstWindow := text[:L]
	freqMap := FrequencyTable(firstWindow, k)
	for s, freq := range freqMap {
		if freq >= t {
			patterns = append(patterns, s)
			foundPatterns[s] = true
		}
	}

	for i := 1; i < n-L+1; i++ {
		oldPattern := text[i-1 : i-1+k]
		freqMap[oldPattern]--

		if freqMap[oldPattern] == 0 {
			delete(freqMap, oldPattern)
		}

		newPattern := text[i+L-k : i+L]
		freqMap[newPattern]++
		for s, freq := range freqMap {
			if freq >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}

	}

	return patterns
}

// FindClumps takes as input a string text, and integers k, L, and t and returns a slice of strings of all strings of length k that appear at least t times in a "window" of length L in the string
func FindClumps(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	foundPatterns := make(map[string]bool)
	n := len(text)

	for i := 0; i < n-L+1; i++ {
		window := text[i : i+L]
		freqMap := FrequencyTable(window, k)
		for s, val := range freqMap {
			if val >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}

	}

	return patterns
}

// Contains takes as input a slice of strings and a single string and returns true if contained in the slice, false otherwise
func Contains(pattern []string, s string) bool {

	for _, pattern := range pattern {
		if pattern == s {
			return true
		}
	}

	return false
}

func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)

	for i := 0; i < n-k+1; i++ {
		pattern := text[i : i+k]
		freqMap[pattern]++
	}

	return freqMap
}
