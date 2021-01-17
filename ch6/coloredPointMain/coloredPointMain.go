package main

import (
	"fmt"
	"image/color"

	"github.com/liuminhaw/gopl/ch6/coloredPoint"
)

func main() {
	var cp coloredPoint.ColoredPoint

	cp.X = 1
	cp.Point.Y = 2
	fmt.Println("Point X position: ", cp.Point.X)
	fmt.Println("Point Y position: ", cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := coloredPoint.ColoredPoint{coloredPoint.Point{1, 1}, red}
	q := coloredPoint.ColoredPoint{coloredPoint.Point{5, 4}, blue}
	fmt.Println("p q distance: ", p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println("p q distance: ", p.Distance(q.Point))
}
