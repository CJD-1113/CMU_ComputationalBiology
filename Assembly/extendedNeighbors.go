package main

// GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
// It returns the extendedNeighbors list. For each neighbor *n* in this list,
// distance between n and pattern is between 2 to maxK.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {

	answer := make([]string, 0)
	vis := InitMap(adjList)
	currNeighbors := make([]string, 0)
	for _, v := range adjList[pattern] {
		currNeighbors = append(currNeighbors, v)
	}
	newNbrs := make([]string, 0)

	for steps := 1; steps <= maxK-1; steps++ {
		for _, u := range currNeighbors {
			for _, v := range adjList[u] {
				if v == pattern {
					continue
				}
				if !vis[v] {
					vis[v] = true
					answer = append(answer, v)
				}
				newNbrs = append(newNbrs, v)
			}
			currNeighbors = newNbrs
		}
	}

	return answer
}

func InitMap(adjList map[string][]string) map[string]bool {
	mp := make(map[string]bool)
	for k := range adjList {
		mp[k] = false
	}
	return mp
}
