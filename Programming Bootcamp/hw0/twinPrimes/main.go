package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("2.5 Twin primes")
}

func NextTwinPrimes(n int) (int, int) {
	var counter int
	if n <= 1 {
		counter = 3
	} else {
		counter = n + 1
	}

	for !(IsPrime(counter) && IsPrime(counter+2)) {
		counter++
	}

	return counter, counter + 2
}

func IsPrime(p int) bool {

	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}
