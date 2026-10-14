package main

import "fmt"

func main() {
	fmt.Println("GCD Algorithms")
	fmt.Println(TrivialGCD(63, 42))
}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TrivialGCD(a, b int) int {
	d := 1
	m := Min2(a, b)

	for p := 1; p < m; p++ {
		if a%p == 0 && b%p == 0 {
			d = p
		}
	}
	return d
}

func EuclidGCCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return a
}
