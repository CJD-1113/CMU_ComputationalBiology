package main

/*
import (
	"fmt"      //formats output in the terminal
	"io"       //read and write to files
	"log"      //logs file error
	"net/http" //for accessing URLS

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	fmt.Println("Clumps")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Recieved non-OK status: %v", resp.Status)
	}

	genomeSymbols, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	genome := string(genomeSymbols)

	fmt.Println("The number of nucleatides in genome is", len(genome))

	EcoliSkewArray := SkewArray(genome)
	minSkewPositions := MinimumSkew(genome)
	fmt.Println("The minimum skew value of", EcoliSkewArray[minSkewPositions[0]], "occurs at positions", minSkewPositions)
	MakeSkewDiagram(EcoliSkewArray)
	fmt.Println("Skew diagram drawn")
}

func MakeSkewDiagram(skewArray []int) {
	p := plot.New()
	p.Title.Text = "Skew Diagram"
	p.X.Label.Text = "Genome Position"
	p.Y.Label.Text = "Skew Value"
	p.Legend.Top = false

	points := make(plotter.XYs, len(skewArray))

	for i, skewValue := range skewArray {
		points[i].X = float64(i)
		points[i].Y = float64(skewValue)
	}

	line, err := plotter.NewLine(points)

	if err != nil {
		panic(err)
	}

	p.Add(line)

	unitOfLength := vg.Centimeter
	p.X.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Y.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Title.TextStyle.Font.Size = 4 * unitOfLength
	p.X.Tick.Label.Font.Size = 2 * unitOfLength
	p.Y.Tick.Label.Font.Size = 2 * unitOfLength

	err = p.Save(100*unitOfLength, 60*unitOfLength, "SkewDiagram.png")

	if err != nil {
		panic(err)
	}
}

// SkewArray takes a genome string and returns a slice of integers corresponding to the G-C skew of the genome
func SkewArray(genome string) []int {
	n := len(genome)

	array := make([]int, n+1)
	array[0] = 0

	for i := 1; i < n+1; i++ {
		array[i] = array[i-1] + Skew(genome[i-1])
	}

	return array
}

// Skew takes a input a symbol and outputs a 1 if a g or -1 if a c and 0 otherwise
func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	} else {
		return 0
	}
}

// MinimumSkew identifies where the skew array reaches a minimum, takes as input a dna string and returns a slice of integer indices corresponding to minimums
func MinimumSkew(genome string) []int {
	indices := make([]int, 0)
	array := SkewArray(genome)
	m := MinIntegerArray(array)

	for i, val := range array {
		if val == m {
			indices = append(indices, i)
		}
	}
	return indices
}

func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: Empty slice given to Min Integer Array")
	}

	var min int

	for i := range list {
		if i == 0 || list[i] < min {
			min = list[i]
		}
	}

	return min
}
*/
