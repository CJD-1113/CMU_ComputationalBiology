package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(GenerateLinearCongruenceSequence(1, 5, 1, 8192))
}

func CountNumDigits(x int) int {
	if x < 0 {
		x *= -1
	}

	if x > 9 {
		return 1 + CountNumDigits(x/10)
	} else {
		return 1
	}
}

func SquareMiddle(x, numDigits int) int {
	if numDigits%2 == 1 || x < 0 || numDigits < 0 || CountNumDigits(x) > numDigits {
		return -1
	}
	x *= x
	s := strconv.Itoa(x)

	for len(s) < 2*numDigits {
		s = "0" + s
	}

	a := (numDigits / 2)
	b := a + numDigits

	i, err := strconv.Atoi(s[a:b])
	if err != nil {
		panic(err)
	}
	return i
}

func HasRepeat(a []int) bool {

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}

	return true

}

func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	nums := make([]int, 0)

	for seed > 0 {
		nums = append(nums, seed)
		seed = SquareMiddle(seed, numDigits)
		if HasRepeat(nums) {
			return nums
		}
	}

	// fmt.Println(356)
	// fmt.Println(SquareMiddle(356,6))
	// fmt.Println(SquareMiddle(126,6))
	// fmt.Println(SquareMiddle(15,6))
	// fmt.Println(SquareMiddle(0,6))
	return nums
}

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
	seq := make([]int, 1)
	seq[0] = seed

	for !HasRepeat(seq) {
		seed = (a*seed + c) % m
		seq = append(seq, seed)
	}

	return seq
}
