package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("CH 1 3.2 Substrings and subslices")

	s := "Hi Lovers!"
	fmt.Println(s[1:4])
	fmt.Println(s[2:4])
	fmt.Println(s[:7])
	fmt.Println(s[4:])

	a := make([]int, 6)
	for i := range a {
		a[i] = 2*i + 1
	}

	fmt.Println(a)
	fmt.Println(a[3:5])
	fmt.Println(a[:3])
	fmt.Println(a[4:])

	pattern := "ATA"
	text := "ATATA"

	fmt.Println(PatternCount(pattern, text))  //yes overlap
	fmt.Println(strings.Count(text, pattern)) //no overlap
}

// PaternCount takes as input two strings pattern and text, returns the number of times that pattern occurs as a substring of text, even overlaps
func PatternCount(pattern, text string) int {
	// count := 0

	// /*
	// 	for i := 0; i < len(text)-len(pattern)+1; i++ {
	// 		if text[i:i+len(pattern)] == pattern {
	// 			count++
	// 		}
	// 	}
	// */

	// for i := range len(text) - len(pattern) + 1 {
	// 	if text[i:i+len(pattern)] == pattern {
	// 		count++
	// 	}
	// }

	// return count
	return len(StartingIndices(pattern, text))

}

// StartingIndices takes as input two strings pattern and text. It reatuurns the collection of all starting positions of pattern as a substring of text, including overlaps
func StartingIndices(pattern, text string) []int {
	positions := make([]int, 0)
	n := len(text)
	k := len(pattern)

	for i := 0; i < n-k+1; i++ {
		if pattern == text[i:i+k] {
			positions = append(positions, i)
		}
	}

	return positions
}
