package main

import "fmt"

func main() {
	fmt.Println("The minimum of 3 and 4 is", Min2(3, 4))
	fmt.Println(WhichIsGreater(3, 5))   //-1
	fmt.Println(WhichIsGreater(42, 42)) //0
	fmt.Println(WhichIsGreater(-2, -7)) //1

	fmt.Println(SameSign(-2, -7)) //t
	fmt.Println(SameSign(2, 7))   //t
	fmt.Println(SameSign(2, -7))  //f
	fmt.Println(SameSign(-2, 7))  //f
}

// Min2 takes two ints as inputs and returns their min
func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//WhichIsGreater compares two input integers and returns 1
//the first input is larger, -1 the second is larger
//and 0 if they are equal
func WhichIsGreater(x, y int) int {
	if x == y {
		return 0
	} else if x > y {
		return 1
	} else {
		return -1
	}
}

//SameSign take as input two integers
//Returns true if the two integers have the same sign
//Returns false if the two integers have different signs
func SameSign(x, y int) bool {
	return x*y >= 0
}

//PositiveDifference takes as input two integers
//Returns the abs value of the difference
func PositiveDifference(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return a - b
	} else {
		return b - a
	}
}
