package main

import (
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")() // Don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // Simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
