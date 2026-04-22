package main

// TrimNetwork takes in an overlap graph adjList and a max iteration maxK
// It removes all n-transitivity edges in adjList, n <= maxK.
// It returns a trimmed copy of the original graph.
func TrimNetwork(adjList map[string][]string, maxK int) map[string][]string {
	adjList2 := make(map[string][]string)

	for key := range adjList {
		adjList2[key] = GetTrimmedNeighbors(key, adjList, maxK)
	}
	return adjList2
}
