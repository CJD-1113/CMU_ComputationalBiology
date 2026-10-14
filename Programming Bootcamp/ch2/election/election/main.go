package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Let's simulate an election!")

	electoralVoteFile := "data/electoralVotes.csv"
	pollFile := "data/debates.csv"

	electoralVotes := ReadElectoralVotes(electoralVoteFile)
	polls := ReadPollingData(pollFile)

	// fmt.Println(electoralVotes)
	// fmt.Println(polls)

	numTrials := 1000000 //100 Million
	marginOfError := 0.1 //5%
	fmt.Println("Running...")
	probability1, probability2, probabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)
	fmt.Println("Probability of Clinton:", probability1)
	fmt.Println("Probability of Trump:", probability2)
	fmt.Println("Probability of Tie:", probabilityTie)

}

// SimulateMultipleElections runs a Monte Carlo simulation with multiple trials to simulate a presidential election. Takes as input a map of polling data to states and a map of state names to electoral college votes. Output is going to be three probabilities, likelihood of each candidate winning and tie
func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]uint, numTrials int, marginOfError float64) (float64, float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	for i := 0; i < numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)
		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}

	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	probabilityTie := float64(tieCount) / float64(numTrials)

	return probability1, probability2, probabilityTie
}

// SimulateOneElection simulates one election. Takes as input a map of polling data to states and a map of state names to electoral college votes. Outputs the votes for candidate 1 and 2 as uint
func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint, marginOfError float64) (uint, uint) {
	var collegeVotes1 uint
	var collegeVotes2 uint

	for state, pollingValue := range polls {
		numVotes := electoralVotes[state]

		adjustedPoll := AddNoise(pollingValue, marginOfError)

		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		}
	}

	return collegeVotes1, collegeVotes2
}

// AddNoise adjusts a polling percentage based on a random sampled from a normal distribution. Takes as input two float64s, a pollling value and a margin of error. Returns an adjusted float64 polling value
func AddNoise(pollingValue, marginOfError float64) float64 {
	x := rand.NormFloat64()
	x /= 2.0
	x *= marginOfError

	return x + pollingValue
}

//505 3.5
