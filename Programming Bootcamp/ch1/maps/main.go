package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps in Go.")
	var polls map[string]float64
	polls = make(map[string]float64)
	//of use polls := make(map[string]float64)

	polls["Pennsylvania"] = 0.517
	polls["Ohio"] = 0.488
	polls["Texas"] = 0.378
	polls["Florida"] = 0.5
	fmt.Println("The number of states in the map is", len(polls))

	delete(polls, "Florida")

	fmt.Println("The number of states in the map is", len(polls))

	/*
		//Array and Slice Literals
		dnaAlphabet := [4]byte{'A', 'C', 'G', 'T'}
		primes := []int{2,3,5,7,11}
	*/

	//Map Literal
	electoralVotes := map[string]uint{
		"Pennsylvania": 20,
		"Ohio":         18,
		"Texas":        38,
	}

	//Pass by reference not value
	UpdateVotes(electoralVotes)

	fmt.Println("The number of electoral votes in Pennsylvania is now", electoralVotes["Pennsylvania"])
	fmt.Println()

	// for state, votes := range electoralVotes {
	// 	fmt.Println("The number of electoral votes in", state, "is", votes)
	// }
	// fmt.Println()

	// for state := range electoralVotes {
	// 	fmt.Println("The number of electoral votes in", state, "is", electoralVotes[state])
	// }
	// fmt.Println()

	// for _, votes := range electoralVotes {
	// 	fmt.Println("The number of electoral votes in state is", votes)
	// }
	// fmt.Println()

	PrintMapAlphabetical(electoralVotes)

}

func PrintMapAlphabetical(dict map[string]uint) {
	keys := make([]string, len(dict))

	i := 0

	for key := range dict {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("The value associated with", key, "is", dict[key])
	}
}

func UpdateVotes(electoralVotes map[string]uint) {
	electoralVotes["Pennsylvania"] = 19
	electoralVotes["Ohio"] = 17
	electoralVotes["Texas"] = 40
}
