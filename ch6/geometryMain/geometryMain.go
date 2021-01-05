package main

import (
	"fmt"

	"github.com/liuminhaw/gopl/ch6/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	fmt.Println("Function distance: ", geometry.Distance(p, q)) // function call
	fmt.Println("Method distance: ", p.Distance(q))             // method call

	perim := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Println("Path distance: ", perim.Distance())
}
