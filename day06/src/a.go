package main

import (
	"math"
)

// Solution from: https://simontoth.substack.com/p/daily-bite-of-c-advent-of-code-day-04f

func partA(lines []string) int {
	times := processLine(lines[0])
	distances := processLine(lines[1])

	prod := 1
	for i := 0; i < len(times); i += 1 {
		b := times[i]
		c := distances[i]

		x1, x2 := solve_quadratic(float64(1.0), float64(-b), float64(c))

		possible_solutions := math.Floor(math.Nextafter(x1, x2)) - math.Ceil(math.Nextafter(x2, x1)) + 1
		prod *= int(possible_solutions)
	}

	return prod
}
