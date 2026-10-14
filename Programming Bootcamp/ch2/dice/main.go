package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Rolling Dice")
	numTrials := 1000

	fmt.Println("Estimated House Edge with", numTrials, "trials is", ComputeCrapsHouseEdge(numTrials))
}

// PlayCrapsOnce simulated one game of craps. returns true if the game is a win and false if it is a loss
func PlayCrapsOnce() bool {
	firstRoll := SumDice(2)

	if firstRoll == 7 || firstRoll == 11 {
		return true
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false
	} else {
		for { //white true
			newRoll := SumDice(2)
			if newRoll == firstRoll {
				return true
			} else if newRoll == 7 {
				return false
			}
		}
	}
}

// ComputeCrapsHouseEdge estimated the house edge in craps over multiple simulations. Takes as input an integer # of simulations and outputs the house edge over those simulations
func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0

	for i := 0; i < numTrials; i++ {
		if PlayCrapsOnce() {
			count++
		} else {
			count--
		}
	}

	return float64(count) / float64(numTrials)
}

func RollDie() int {
	return rand.Intn(6) + 1 //[0,5] + 1
}

func SumTwoDice() int {
	return RollDie() + RollDie()
}

func SumDice(numDice int) int {
	sum := 0

	for i := 0; i < numDice; i++ {
		sum += RollDie()
	}

	return sum
}
