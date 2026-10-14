package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Finding primes!")
	n := 11
	fmt.Println(TrivialPrimeFinder(n))
	fmt.Println(SieveOfEratoshemes(n))
	fmt.Println(ListPrimes(n))
}

// TrivialPrimeFinder takes as input an integer n and returns a slice of booleans ariables storing the primality of each nonnegative integer up to and including n
func TrivialPrimeFinder(n int) []bool {
	primeBooleans := make([]bool, n+1)

	for p := 2; p <= n; p++ {
		primeBooleans[p] = IsPrime(p)
	}

	return primeBooleans
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

// SieveOfEratosthemes takes as input an integer n and returns a slice of boolean variables storing the primarlity of each nonnegative interger up to and including n
func SieveOfEratoshemes(n int) []bool {
	primeBooleans := make([]bool, n+1)
	//0, 1, 2  are false and then 2 onward is true

	for p := 2; p <= n; p++ {
		primeBooleans[p] = true
	}

	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		//if p is prime, make all multiples false
		if primeBooleans[p] == true {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}
	}

	return primeBooleans
}

//CrossOffMultiples takes as input a boolean slice prime Booleans and an integer and returns a slice in which all multiples of the integer have been set to false

func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans)

	//look at every multiple of p, make it false
	//start with 2p

	for i := 2 * p; i < n; i += p {
		primeBooleans[i] = false
	}

	return primeBooleans
}

// ListPrimes takes an integer and returns a slice containing all prime numbers up to and including n
func ListPrimes(n int) []int {
	primeList := make([]int, 0)
	primeBooleans := SieveOfEratoshemes(n)

	for p := range primeBooleans {
		if primeBooleans[p] {
			primeList = append(primeList, p)
		}
	}

	return primeList
}
