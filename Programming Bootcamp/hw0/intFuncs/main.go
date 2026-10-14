package main

import "fmt"

func main() {
	fmt.Println("2.2 More practice with integer functions")
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

// SumProperDivisors takes an integer n and returns the sum of all divisors of n
func SumProperDivisors(n int) int {
	//for loop, up to but excluding n
	//if n% i == 0 then add i to sum

	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum
}
