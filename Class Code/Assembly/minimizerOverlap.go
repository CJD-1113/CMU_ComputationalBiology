package main

import "fmt"

func MakeOverlapNetworkMinimizers(reads []string, minimizerDictionary StringIndex, match, mismatch, gap, threshold float64) map[string][]string {

	numReads := len(reads)
	overlapMatrix := InitializeFloatMatrix(numReads, numReads)
	bigNegative := threshold - 100000

	for r := range overlapMatrix {
		for c := range overlapMatrix[r] {
			overlapMatrix[r][c] = bigNegative
		}
	}

	counter := 0

	for _, readIndices := range minimizerDictionary {
		if counter%100 == 0 {
			fmt.Println("Now considering element", counter, "of the minimizer")
		}
		for i := range readIndices {
			for j := i + 1; j < len(readIndices); j++ {

				index1 := readIndices[i]
				index2 := readIndices[j]

				read1 := reads[index1]
				read2 := reads[index2]

				if overlapMatrix[index1][index2] == bigNegative {
					overlapMatrix[index1][index2] = ScoreOverlapAlignment(read1, read2, match, mismatch, gap)
				}

				if overlapMatrix[index2][index1] == bigNegative {
					overlapMatrix[index2][index1] = ScoreOverlapAlignment(read2, read1, match, mismatch, gap)
				}

			}
		}
		counter++
	}

	b := BinarizeMatrix(overlapMatrix, threshold)
	return ConvertAdjacencyMatrixToList(reads, b)
}
