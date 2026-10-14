package main

import "fmt"

func main() {
	fmt.Println("Functions!")

	x := 3
	n := SumTwoInts(x, x)
	fmt.Println(n)

	m := 17
	fmt.Println(AddOne(m))
	fmt.Println(m)

	PrintHi()

}

//AddOne takes integer k as input and returns k+1
func AddOne(k int) int {
	k = k + 1
	return k
}

//SumTwoInts takes two inputs as input and returns their sum
func SumTwoInts(a int, b int) int {
	return a + b
}

//DoubleAndDuplicate takes as input a float64
//Returns two copies of that input variable.
func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0 * x, 2.0 * x
}

//Pi takes no inputs and returns the value of pi
func Pi() float64 {
	return 3.14
}

//PrintHi prints "Hi" to the console
func PrintHi() {
	fmt.Println("Hi")
}
