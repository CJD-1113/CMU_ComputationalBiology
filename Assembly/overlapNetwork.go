package main

// MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
// It returns adjacency lists of reads; edges are only included
// in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {

	overlapMatrix := OverlapScoringMatrix(reads, match, mismatch, gap)
	binarizedMatrix := BinarizeMatrix(overlapMatrix, threshold)
	return ConvertAdjacencyMatrixToList(reads, binarizedMatrix)

}

func ConvertAdjacencyMatrixToList(reads []string, b [][]int) map[string][]string {
	adjacencyList := make(map[string][]string)

	for r := range b {
		for c := range b[r] {
			if b[r][c] == 1 {
				currentRead := reads[r]
				adjacencyList[currentRead] = append(adjacencyList[currentRead], reads[c])
			}
		}
	}

	return adjacencyList
}
