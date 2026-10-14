package main

import (
	"fmt"
)

func main() {
	fmt.Println("Two-dimensional arrays (and slices).")

	/*
		var a [7][4]int
		a[0][0] = 42
		a[1][2] = 19
		a[6][3] = 100

		fmt.Println(a)
		fmt.Println(a[1])
		fmt.Println("a has", len(a), "rows and", len(a[0]), "columns")

		//var board [][]bool //2d slice
		numRows := 4
		board := make([][]bool, numRows)
		for row := range board {
			board[row] = make([]bool, row)
		}

		fmt.Println(board[2][0])
		PrintBoard(board)

		for row := range board {
			board[row] = append(board[row], true)
		}

		newRow := make([]bool, 5)
		board = append(board, newRow)

		PrintBoard(board)
	*/

	rPentomino := make([][]bool, 5)
	for row := range rPentomino {
		rPentomino[row] = make([]bool, 5)
	}

	rPentomino[1][2] = true
	rPentomino[1][3] = true
	rPentomino[2][1] = true
	rPentomino[2][2] = true
	rPentomino[3][2] = true

	PrintBoard(rPentomino)
}

func PrintBoard(board [][]bool) {
	for _, row := range board {
		PrintRow(row)
	}
}

func PrintRow(row []bool) {
	for _, val := range row {
		PrintCell(val)
	}
	fmt.Println()
}

func PrintCell(value bool) {
	if value {
		fmt.Print("⬛️")
	} else {
		fmt.Print("⬜️")
	}
}
