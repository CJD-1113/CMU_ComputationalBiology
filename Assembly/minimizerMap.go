package main

// StringIndex is a type that will map a minimizer string to its list of indices
// in a read set corresponding to reads with this minimizer.
type StringIndex map[string][]int

// BuildMinimizerMap takes a collection of reads, integer k and integer windowLength.
// It returns a map of k-mers to the indices of the reads in the list having this k-mer minimizer.
func BuildMinimizerMap(reads []string, k int, windowLength int) StringIndex {
	dict := make(StringIndex)

	for i, read := range reads {
		n := len(read)

		for j := 0; j < n-windowLength+1; j++ {
			currentWindow := read[j : j+windowLength]
			m := Minimizer(currentWindow, k)
			if !IsElementOf(dict[m], i) {
				dict[m] = append(dict[m], i)
			}

		}
	}

	return dict
}

func IsElementOf(a []int, givenValue int) bool {
	for _, val := range a {
		if val == givenValue {
			return true
		}
	}

	return false
}
