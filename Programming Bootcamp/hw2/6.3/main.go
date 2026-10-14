package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	num := 1 + rand.Intn(366)
	// 	if num == 1 {
	// 		fmt.Println("hello")
	// 	} else if num == 365 {
	// 		fmt.Println("Yello")
	// 	}
	// }
	fmt.Println(SharedBirthdayProbability(5000, 40000))
	fmt.Println(SmallestValue(1000))
}

func SmallestValue(numTrials int) int {
	people := 2
	chance := 0.0

	for chance > .5 {
		people++
		chance = SharedBirthdayProbability(people, numTrials)
	}

	return people
}

func SharedBirthdayProbability(numPeople, numTrials int) float64 {
	count := 0

	for i := 0; i < numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			count++
		}
	}

	return float64(count) / float64(numTrials)
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

func SimulateOneBirthdayTrial(num_people int) bool {
	bdays := make([]int, 0)
	for i := 0; i < num_people; i++ {
		bdays = append(bdays, 1+rand.Intn(366))
	}
	return HasRepeat(bdays)
}
