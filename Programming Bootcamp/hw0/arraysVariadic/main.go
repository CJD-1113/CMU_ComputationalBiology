package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

// FibonacciArray takes an integer n and returns an array of length n+1 whose kth element is the kth in the sequence
func FibonacciArray(n int) []int {
	fib := make([]int, n+1)
	fib[0] = 1

	if n > 0 {
		fib[1] = 1
	}

	for i := 2; i < n+1; i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}

	return fib
}

// DividesAll() takes as input an array of integers a and integer d, returns true if d is a divisor of every element in a
func DividesAll(a []int, d int) bool {
	if d == 0 {
		return false
	}
	for x := range a {
		if a[x]%d != 0 {
			return false
		}
	}
	return true
}

// MaxIntegerArray takes as input an array of integers list, and returns the maximum value in that array
func MaxIntegerArray(list []int) int {
	max := list[0]

	for x := range list {
		if list[x] > max {
			max = list[x]
		}
	}

	return max
}

// MaxIntegers takes as input intgers and returns the maximum value of all of the inputs
func MaxIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: Empty slice given to Max Integer Array")
	}

	return MaxIntegerArray(numbers)
}

//SumIntegers takes as input a collection of integers and returns the sum of all of the inputs

func SumIntegers(numbers ...int) int {
	sum := 0
	for x := range numbers {
		sum += numbers[x]
	}
	return sum
}

// GCDArray takes an array of integers a and returns the greatest common divisor of all of the integets
func GCDArray(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	gcf := EuclidGCD(a[0], a[1])

	for i := 1; i < len(a); i++ {
		gcf = EuclidGCD(gcf, a[i])
	}

	return gcf
}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
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
