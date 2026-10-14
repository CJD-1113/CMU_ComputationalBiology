package main

import (
	"fmt"
	"math/rand" // this should be helpful!
)

func main() {
	fmt.Println("Running...")
	fmt.Println(RelativelyPrimeProbability(1, 10, 100))
}

func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	count := 0

	for i := 0; i < numPairs; i++ {
		n1 := lowerBound + rand.Intn(upperBound-lowerBound+1)
		n2 := lowerBound + rand.Intn(upperBound-lowerBound+1)
		if RelativelyPrime(n1, n2) {
			count++
		}
	}

	return float64(count) / float64(numPairs)
}

func RelativelyPrime(a, b int) bool {
	if EuclidGCD(a, b) == 1 {
		return true
	}
	return false
}

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
