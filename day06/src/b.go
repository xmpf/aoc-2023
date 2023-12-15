package main

import "math"

func partB(lines []string) int {
	b := processLine_b(lines[0])
	c := processLine_b(lines[1])
	x1, x2 := solve_quadratic(float64(1.0), float64(-b), float64(c))
	possible_solutions := math.Floor(math.Nextafter(x1, x2)) - math.Ceil(math.Nextafter(x2, x1)) + 1
	return int(possible_solutions)
}
