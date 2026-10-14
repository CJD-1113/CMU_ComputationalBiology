package main

import (
	"fmt"
	"math"
	"math/rand" // this should be helpful!
)

func main() {
	fmt.Println("Pi is", OptimizedEstimatePi(10000))
}

func WeightedDie() int {
	ran := rand.Float64()

	if ran <= 0.5 {
		return 3
	} else if ran <= 0.6 {
		return 1
	} else if ran <= 0.7 {
		return 2
	} else if ran <= 0.8 {
		return 4
	} else if ran <= 0.9 {
		return 5
	} else {
		return 6
	}
}

func OptimizedEstimatePi(numPoints int) float64 {
	count := 0

	for i := 0; i < numPoints; i++ {
		x := rand.Float64()*2.0 - 1
		y := rand.Float64()*2.0 - 1
		if math.Sqrt((x*x)+(y*y)) <= 1 {
			count++
		}
	}

	return (4.0 * float64(count)) / float64(numPoints)
}

func EstimatePi(numPoints int) float64 {
	xs, ys := GeneratePoints(numPoints)
	count := 0

	for i := range xs {
		if DistanceFromOrigin(xs[i], ys[i]) <= 1 {
			count++
		}
	}

	return (4.0 * float64(count)) / float64(numPoints)
}

func GeneratePoints(numPoints int) ([]float64, []float64) {
	xs := make([]float64, 0)
	ys := make([]float64, 0)

	for i := 0; i < numPoints; i++ {
		xs = append(xs, rand.Float64()*2.0-1)
		ys = append(ys, rand.Float64()*2.0-1)
	}

	return xs, ys
}

func DistanceFromOrigin(x1, y1 float64) float64 {
	return math.Sqrt((x1 * x1) + (y1 * y1))
}
