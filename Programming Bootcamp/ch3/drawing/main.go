package main

import (
	"canvas/canvas"
	"fmt"
)

func main() {
	fmt.Println("Drawing a head.")

	DrawFace()
}

func DrawFace() {

	c := canvas.CreateNewCanvas(1000, 2000)
	black := canvas.MakeColor(0, 0, 0)
	c.SetFillColor(black)
	c.Clear()
	c.SaveToPNG("fun.png")
}
