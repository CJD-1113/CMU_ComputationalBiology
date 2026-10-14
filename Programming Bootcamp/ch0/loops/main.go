package main

import "fmt"

func main() {
	fmt.Println("Loops!")

	var count uint = 10

	for ; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println("Blast Off!")
}

//While Loop Video

// Factorial takes 1 integer n as input and returns n!
func Factorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial()")
	}

	p := 1 //stores product
	i := 1 //counter

	//go has no keyword while and uses for instead
	for i <= n {
		p = p * i
		i += 1
	}

	return p
}

//SumFirstNIntegers takes as input an integer n
//and returns the sum of the first n possitive integers

func SumFirstNIntegers(n int) int {
	if n < 0 {
		panic("Error: negative input given to SumFirstNIntegers()")
	}

	sum := 0
	i := 0

	for i <= n {
		sum += i
		i++
	}

	return sum
}

//For Loop Video
//------------------------

//AnotherFactorial takes as input an integer n and returns n!
func AnotherFactorial(n int) int {

	p := 1

	for i := 1; i <= n; i++ {
		p *= i
	}

	return p
}

//AnotherSum takes an integer n as input and returns the sum
//of the first n positive integers

func AnotherSum(n int) int {

	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}

func YetAnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to YetAnotherFactorial")
	}

	p := 1

	for i := n; i >= 1; i-- {
		p += i
	}

	return p
}

func SumEvent(k int) int {
	sum := 0

	for i := 2; i <= k; i = i + 2 {
		sum += i
	}

	return sum
}

func GaussSum(n int) int {
	return n * (n + 1) / 2
}
