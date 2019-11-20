// Dup1 prints the text of each line that appears more than
// once in the standard output, precede by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// Ctrl-D to send EOF (Terminate Scan)
	for input.Scan() {
		counts[input.Text()]++
	}

	// Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
