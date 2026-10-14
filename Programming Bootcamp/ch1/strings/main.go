package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ch 1 Code Along 3.1")

	// fmt.Println(string('A'))
	// fmt.Println(string(45))

	// fmt.Println(strconv.Itoa(45))
	// j, err := strconv.Atoi("23")
	// pi, err2 := strconv.ParseFloat("3.14", 64)

	// if err != nil {
	// 	panic(err)
	// }

	// if err2 != nil {
	// 	panic(err2)
	// }

	// fmt.Println(j)
	// fmt.Println(pi)

	// var s string = "Hi"
	// t := "Lovers"
	// u := s + t

	// fmt.Println(u)
	// fmt.Println(s, t)
	// fmt.Println("The first symbol of u is", string(u[0]))
	// fmt.Println("The final symbol of u is", string(u[len(u)-1]))

	// if t[2] == 'v' {
	// 	fmt.Println("The symbol at position 2 of t is v.")
	// }

	dna := "ACCGAT"

	fmt.Println(Complement(dna))        //TGGCTA
	fmt.Println(Reverse(dna))           //TAGCCA
	fmt.Println(ReverseComplement(dna)) //ATCGGT
}

// ReverseComplement takes as input a string patter of dna symbols and returns the complimentary nucletide in reverse
func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

// Reverse takes as input a string and returns it in reverse
func Reverse(pattern string) string {
	// pattern2 := ""
	// for x := range pattern {
	// 	pattern2 = string(pattern[x]) + pattern2
	// }

	// return pattern2

	rev := make([]byte, len(pattern))

	for i := range pattern {
		rev[i] = pattern[len(pattern)-1-i]
	}

	return string(rev)
}

// Complement takes as input a dna string and returns it with its complementary nucleatides
func Complement(dna string) string {
	dna2 := make([]byte, len(dna))
	for i, symbol := range dna {
		switch symbol {
		case 'A':
			dna2[i] = 'T'
		case 'C':
			dna2[i] = 'G'
		case 'G':
			dna2[i] = 'C'
		case 'T':
			dna2[i] = 'A'
		default:
			panic("Invalid Symbol given to complement()")
		}
	}

	return string(dna2)
}
