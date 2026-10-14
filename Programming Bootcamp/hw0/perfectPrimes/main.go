package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("2.4 Perfect numbers and Mersenne primes")
}

// IsPerfect takes an int and returns true if it is perfect
// it is perfect if the sum of the divisors equals the number
func IsPerfect(n int) bool {
	if n == SumProperDivisors(n) {
		return true
	}
	return false
}

// SumProperDivisors takes an integer n and returns the sum of all divisors of n
func SumProperDivisors(n int) int {
	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum
}

// NextPerfectNumber finds the next perfect number after the number inputted
func NextPerfectNumber(n int) int {
	counter := n + 1
	for !IsPerfect(counter) {
		counter++
	}

	return counter
}

// IsPrime takes as input an integer p and returns true if p is prime and false otherwise
func IsPrime(p int) bool {

	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

// Power takes integer a and b as input, a is the base while b is the exponent, returns a^b
func Power(a, b int) int {
	if b == 0 {
		return 1
	} else if b == 1 {
		return a
	} else {
		return a * Power(a, b-1)
	}
}

// ListMersennePrimes takes an integer n and returns an array of all primes that are in the form of 2^p-1 where 1 < p <= n
func ListMersennePrimes(n int) []int {
	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		candidate := Power(2, i) - 1
		if IsPrime(candidate) {
			primes = append(primes, candidate)
		}
	}

	return primes
}
