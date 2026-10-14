package main

import "math/rand"

// SimulateReadsClean takes a genome along with a read length and a probability.
// It returns a collection of strings resulting from simulating clean reads,
// where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {
	n := len(genome)
	reads := make([]string, 0, n-readLength+1)

	for i := 0; i < n-readLength+1; i++ {
		randNumber := rand.Float64()
		if randNumber < probability {
			currentRead := genome[i : i+readLength]
			reads = append(reads, currentRead)
		}
	}

	return reads
}

func SimulateReadsMessy(genome string, readLength int, probability, substititonErrorRate, insertionErrorRate, deletetionErrorRate float64) []string {
	n := len(genome)
	reads := make([]string, 0, n-readLength+1)

	for i := 0; i < n-readLength+1; i++ {
		randNumber := rand.Float64()
		if randNumber < probability {
			currentRead := genome[i : i+readLength]
			reads = append(reads, currentRead)
		}
	}

	return reads
}
