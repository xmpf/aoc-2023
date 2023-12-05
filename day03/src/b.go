package main

import (
	"math"
)

func partB(lines []string) int64 {

	gears := make([][]int, 0)
	numbers_store := make(map[int][][]int64, 0)

	var sum int64 = 0

	for row := 1; row < len(lines)-1; row += 1 {
		var parsed_value int64
		var start int

		for col := 1; col < len(lines[0])-1; col += 1 {
			if lines[row][col] == '*' { // find gears
				gears = append(gears, []int{row, col})
			}

			if isDigit(lines[row][col]) {
				if parsed_value == 0 {
					parsed_value = int64(lines[row][col] - '0')
					start = int(col)
				} else {
					parsed_value *= 10
					parsed_value += int64(lines[row][col] - '0')
				}
			} else if parsed_value != 0 {
				numbers_store[row] = append(numbers_store[row], []int64{parsed_value, int64(start), int64(col) - 1})

				// reset
				parsed_value = 0
				start = 0
			}
		}
		if parsed_value != 0 {
			numbers_store[row] = append(numbers_store[row], []int64{parsed_value, int64(start), int64(len(lines[0]) - 2)})
		}
	}

	// fmt.Printf("Numbers store: %+v\n", numbers_store)
	// fmt.Printf("Gears map: %+v\n", gears)

	// var gear_ratios []int

	for i := 0; i < len(gears); i += 1 {
		x, y := gears[i][0], gears[i][1]
		sum += findAdjacent(numbers_store, x, y)
	}

	return sum
}

func findAdjacent(numbers_store map[int][][]int64, x, y int) int64 {
	adjacents := make([]int64, 0)

	// fmt.Printf("x = %d, y = %d\n", x, y)
	// fmt.Printf("x = %d, y = %d, store = %+v\n", x, y, numbers_store)

	for delta := -1; delta <= 1; delta += 1 {
		numbers := numbers_store[x+delta]

		// fmt.Printf("numbers[x-1] = %+v\n", numbers_store[x-1])

		if len(numbers) > 0 {
			for _, number := range numbers {
				start := number[1]
				end := number[2]
				var diff float64

				if end <= int64(y) {
					diff = math.Abs(float64(end - int64(y)))
				} else {
					diff = math.Abs(float64(start - int64(y)))
				}

				// fmt.Printf("diff = %.2f\n", diff)

				if diff <= 1.0 { // -1, 0, +1
					adjacents = append(adjacents, number[0])
				}
			}
		}
	}

	if len(adjacents) == 2 {
		prod := adjacents[0] * adjacents[1]
		// fmt.Printf("Adjacents: %+v\n", adjacents)
		// fmt.Printf("product: %d\n", prod)
		return prod
	}
	return 0
}
