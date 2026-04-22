package main

// GetTrimmedNeighbors takes in a string pattern (read), an adjacency list and maxK.
// It returns all n-transitivity edges in the adjList of the current read (pattern) for n <= maxK.
func GetTrimmedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := adjList[pattern]
	extendedNeighbors := GetExtendedNeighbors(pattern, adjList, maxK)

	for i := len(neighbors) - 1; i >= 0; i-- {
		if Contains(extendedNeighbors, neighbors[i]) {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
		}
	}

	return neighbors
}

func Contains(patterns []string, element string) bool {
	for _, pattern := range patterns {
		if pattern == element {
			return true
		}
	}

	return false
}
