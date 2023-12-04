package main

import (
	"fmt"
	"os"
)

func main() {

	filename := "./data/input"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	// read input file
	buf := readFile(filename)

	// solution A
	fmt.Printf("Part A = %d\n", partA(buf))

	// solution B
	fmt.Printf("Part B = %d\n", partB(buf))
}
