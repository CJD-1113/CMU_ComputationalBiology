package main

// AverageOutDegree takes the adjacency list of a directed network.
// It returns the average outdegree of each node in the network.
// It does not include nodes with no outgoing edges in the average.
func AverageOutDegree(adjList map[string][]string) float64 {
	count := 0
	totalNodes := 0

	for currentRead := range adjList {
		numNeighbors := len(adjList[currentRead])
		count += numNeighbors
		if numNeighbors > 0 {
			totalNodes++
		}
	}

	return float64(count) / float64(totalNodes)
}

// AverageOutDegreeAllNodes takes as input the adjacency list of a directed network.
// It returns the average outdegree of each node in the network, including nodes with outdegree zero.
func AverageOutDegreeAllNodes(adjList map[string][]string) float64 {
	count := 0

	for currentRead := range adjList {
		count += len(adjList[currentRead])
	}

	return float64(count) / float64(len(adjList))
}
