package main

import (
	"fmt"
)

func main() {
	lines := readFile("./data/input")
	fmt.Printf("Part A: %d\n", partA(lines))
	fmt.Printf("Part B: %d\n", partB(lines))
}
