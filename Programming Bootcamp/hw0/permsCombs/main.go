package main

import (
	"fmt"
)

func main() {
	fmt.Println("2.1 Permutations and combinations")
	fmt.Println(Factorial(6))
	fmt.Println(Permutation(12, 1))
	fmt.Println(Permutation(6, 6))
	fmt.Println(Permutation(10, 0))
}

// Permutation takes integer n and k as input and returns the permutation statistic
func Permutation(n, k int) int {
	//returns n! * (n-k)!
	// we can use recursion to compute both factorials
	return Factorial(n) / Factorial(n-k)
}

// Combination takes integer n and k as input and returns n! / ( (n-k)! * k!)
func Combination(n, k int) int {
	return Factorial(n) / (Factorial(n-k) * Factorial(k))
}

// Factorial takes integer n and returns 1 * 2 * 3 * ... n!
func Factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}
