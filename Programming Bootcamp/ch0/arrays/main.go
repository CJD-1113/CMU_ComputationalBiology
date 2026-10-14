//Arrays are all of the same time
//Arrays have a fixed length
//Slices are an array with variable length

package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices.")

	var list [6]int
	list[0] = -8
	i := 3
	list[2*i-4] = 17
	list[len(list)-1] = 43
	fmt.Println(list)

	var aSlice []int
	n := 4
	aSlice = make([]int, n)
	aSlice[0] = 6
	aSlice[2] = -33
	fmt.Println(aSlice)

	bSlice := make([]int, n+2)
	fmt.Println(bSlice)

	fmt.Println(FactorialArray(6))

	var c [6]int
	d := make([]int, 6)

	//Array is pass by value
	//	Array makes a copy of itself and
	//	destroys itself after the program runs
	//Slice is pass by reference
	//	When you pass in a slice you are accessing
	//	the orginal slice

	ChangeFirstElementArray(c)
	ChangeFirstElementSlice(d)
	fmt.Println("Array: ", c)
	fmt.Println("Slice: ", d)

	fmt.Println(MinIntegers(3, 5, 7, 1, 5, -43))
}

//Factorial  takes as input an integer n and it
// returns an array of length n+1 whose kth
// element is k!
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Negative Input on Factorial Array")
	}

	fact := make([]int, n+1)
	fact[0] = 1

	for k := 1; k <= n; k++ {
		fact[k] = k * fact[k-1]
	}

	return fact
}

func ChangeFirstElementArray(a [6]int) {
	a[0] = 1
}

func ChangeFirstElementSlice(a []int) {
	a[0] = 1
}

//MinIntegerArray takes as input a slice of integers and returns the minimum calue in that slice
func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: Empty slice given to Min Integer Array")
	}

	var min int
	// for i := 1; i < len(list); i++ {
	// 	if list[i] < min {
	// 		min = list[i]
	// 	}
	// }

	for i := range list {
		if i == 0 || list[i] < min {
			min = list[i]
		}
	}

	return min
}

func MinIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: Empty slice given to Min Integer Array")
	}

	return MinIntegerArray(numbers)
}
