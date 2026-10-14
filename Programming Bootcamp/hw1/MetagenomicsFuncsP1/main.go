package main

import "math"

func main() {
	//380 5 hr
	//425 4.5 hr
}

// Richness takes as input a frequency table and returns the number of elements with frequency over 0
func Richness(sample map[string]int) int {
	elements := 0

	for _, val := range sample {
		if val > 0 {
			elements++
		}
	}

	return elements
}

// SumOfVales takes as input a frequency table and returns the sum of all the occurances
func SumOfValues(sample map[string]int) int {
	sum := 0

	for _, val := range sample {
		sum += val
	}

	return sum
}

//SimpsonsIndex takes as input a frequenc table and returns the simpsons index

// Simpsonos index is eequal to each value divided by sum of all and then added together
func SimpsonsIndex(sample map[string]int) float64 {
	index := 0.0
	total := float64(SumOfValues(sample))

	for _, val := range sample {
		index = index + (math.Pow(float64(val)/total, 2))
	}

	return index
}
