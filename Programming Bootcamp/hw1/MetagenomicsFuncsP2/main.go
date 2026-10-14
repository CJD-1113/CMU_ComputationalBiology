package main

import "fmt"

func main() {
	//425 4.5 hr
	//485 3.5 hr
	fmt.Println("MetaFuncs Hw1 Part 2")
}

func SumOfValues(sample map[string]int) int {
	sum := 0

	for _, val := range sample {
		sum += val
	}

	return sum
}

func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//SumOfMinima takes as input two maps of strings to int, then returns the sum of the minimum of each map

// So start by for loop of first map, then for loop on the inside for second. If first string == second string then add whichever is a minimum sum to the return
func SumOfMinima(sample1, sample2 map[string]int) int {
	sum := 0

	for aniOne, valOne := range sample1 {
		for aniTwo, valTwo := range sample2 {
			if aniOne == aniTwo {
				sum += Min2(valOne, valTwo)
			}
		}
	}

	return sum
}

// SumOfMaxima takes two maps of strings to ints and returns the sum of all of the maximum integers values associated with a key
// make a third map, add the key and value to map. If the key is already there, replace with higher value. Then sum all of the values and return
func SumOfMaxima(sample1, sample2 map[string]int) int {
	combinedSample := make(map[string]int)

	for key := range sample1 {
		combinedSample[key] = sample1[key]
	}

	for key := range sample2 {
		if sample2[key] > combinedSample[key] {
			combinedSample[key] = sample2[key]
		}
	}

	return SumOfValues(combinedSample)
}

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	min := float64(SumOfMinima(sample1, sample2))
	avgTotal := float64(SumOfValues(sample1)+SumOfValues(sample2)) / 2.0

	return 1.0 - (min / avgTotal)
}
