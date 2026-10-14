package main

import "fmt"

func main() {
	SARSOverlapNetworkMinimizerTrim()
}

func SARSOverlapNetworkMinimizerTrim() {
	fmt.Println("Read in SARS-CoV-2 genome.")
	genome := ReadGenomeFromFASTA("Data/SARS-Cov-2_genome.fasta")
	fmt.Println("Genome read")

	readLength := 150
	probability := 0.1
	substititonErrorRate := 0.01
	insertionErrorRate := 0.01
	deletetionErrorRate := 0.01
	reads := SimulateReadsMessy(genome, readLength, probability, substititonErrorRate, insertionErrorRate, deletetionErrorRate)
	fmt.Println("Reads generated. We have", len(reads), "reads")

	k := 10
	windowLength := 20
	minimizerDictionary := BuildMinimizerMap(reads, k, windowLength)
	fmt.Println("Minimizer map made. It conains", len(minimizerDictionary), "total keys")

	match := 1.0
	mismatch := 1.0
	gap := 5.0
	threshold := 40.0
	adjList := MakeOverlapNetworkMinimizers(reads, minimizerDictionary, match, mismatch, gap, threshold)
	fmt.Println("Overlap network made!")
	fmt.Println("The network has", len(adjList), "total keys.")
	fmt.Println("The average outdegree is", AverageOutDegree(adjList))

	maxK := 3
	trimmedAdjacencyList := TrimNetwork(adjList, maxK)
	fmt.Println("Graph now has", len(trimmedAdjacencyList), "total reads")
	fmt.Println("Average outdegree is", AverageOutDegree(trimmedAdjacencyList))
}

func SARSOverlapNetworkMinimizer() {
	fmt.Println("Read in SARS-CoV-2 genome.")
	genome := ReadGenomeFromFASTA("Data/SARS-Cov-2_genome.fasta")
	fmt.Println("Genome read")

	readLength := 150
	probability := 0.1
	reads := SimulateReadsClean(genome, readLength, probability)
	fmt.Println("Reads generated. We have", len(reads), "reads")

	k := 10
	windowLength := 20
	minimizerDictionary := BuildMinimizerMap(reads, k, windowLength)
	fmt.Println("Minimizer map made. It conains", len(minimizerDictionary), "total keys")

	match := 1.0
	mismatch := 1.0
	gap := 5.0
	threshold := 40.0
	adjList := MakeOverlapNetworkMinimizers(reads, minimizerDictionary, match, mismatch, gap, threshold)
	fmt.Println("Overlap network made!")
	fmt.Println("The network has", len(adjList), "total keys.")
	fmt.Println("The average outdegree is", AverageOutDegree(adjList))
}

func RunGenomeAssembly() {
	fmt.Println("Genome assembly!")

	genomeLength := 1000
	k := 20
	originalGenome := GenerateRandomGenome(genomeLength)
	reads := KmerComposition(originalGenome, k)
	reads = ShuffleStrings(reads)

	assembledGenome := GreedyAssembler(reads)
	fmt.Println(assembledGenome)
	kmers := KmerComposition(assembledGenome, k)
	if StringSliceEquals(reads, kmers) {
		fmt.Println("Success.")
	}

	fmt.Println("Program exiting.")
}
