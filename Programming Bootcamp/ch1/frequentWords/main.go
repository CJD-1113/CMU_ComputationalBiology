package main

import "fmt"

func main() {
	fmt.Println("Finding frequent words.")

	text := "ATCAATGATCAACGTAAGCTTCTAAGCATGATCAAGGTGCTCACACAGTTTATCCACAACCTGAGTGGATGACATCAAGATAGGTCGTTGTATCTCCTTCCTCTCGTACTCTCATGACCACGGAAAGATGATCAAGAGAGGATGATTTCTTGGCCATATCGCAATGAATACTTGTGACTTGTGCTTCCAATTGACATCTTCAGCGCCATATTGCGCTGGCCAAGGTGACGGAGCGGGATTACGAAAGCATGATCATGGCTGTTGTTCTGTTTATCTTGTTTTGACTGAGACTTGTTAGGATAGACGGTTTTTCATCACTGACTAGCCAAAGCCTTACTCTGCCTGACATCGACCGTAAATTGATAATGAATTTACATGCTTCCGCGACGATTTACCTCTTGATCATCGATCCGATTGAAGATCTTCAATTGTTAATTCTCTTGCCTCGACTCATAGCCATGATGAGCTCTTGATCATGTTTCCTTAACCCTCTATTTTTTACGGAAGAATGATCAAGCTGCTGCTCTTGATCATCGTTTC"
	k := 9

	freqMap := FrequencyTable(text, k)
	freqPatterns := FindFrequentWords(text, k)

	fmt.Println("Frequent Pattern(s) when k is", k, "are", freqPatterns)
	fmt.Println("The maximum number of occurences is", freqMap[freqPatterns[0]])
}

//FindFrequentWords takes as input a string text and an integer k. It returns a slice of strings corresponding to the substring(s) of length k that occur most frequently in text
func FindFrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)

	freqMap := FrequencyTable(text, k)
	max := MaxMapValue(freqMap)

	for pattern, val := range freqMap {
		if val == max {
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

//MaxMapValue takes as input a map of strings to integers and returns the maximum value found in the map
func MaxMapValue(dict map[string]int) int {
	m := 0
	firstTime := true

	//range over map and update m for larger values

	for _, val := range dict {
		if firstTime || val > m {
			firstTime = false
			m = val
		}
	}

	return m
}

//FrequencyTable takes text and k as input and produces a map associating each substring with length k as a key with its value which is the mode of the substring
func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)

	for i := 0; i < n-k+1; i++ {
		pattern := text[i : i+k]
		freqMap[pattern]++
		//if it is a key it already gets incremented
		//if not a key it gets initialized to 1

		/*
			//exists is true if already a key, and false if not a ket
			_, exists := freqMap[pattern]
			if exists == false {
				freqMap[pattern] = 1
			} else {
				freqMap[pattern]++
			}
		*/
	}

	return freqMap
}
