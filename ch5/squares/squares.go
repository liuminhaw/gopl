package main

import "fmt"

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	g := squares()

	fmt.Println("f:", f())
	fmt.Println("f:", f())
	fmt.Println("g:", g())
	fmt.Println("f:", f())
	fmt.Println("f:", f())
	fmt.Println("g:", g())
}
