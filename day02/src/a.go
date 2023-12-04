package main

// 12 red cubes, 13 green cubes, and 14 blue cubes
var CUBES map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func validateCubes(cubes map[string][]int) bool {
	for k, v := range cubes {
		// k => color
		// v => list of cubes for this color
		if max(v) > CUBES[k] {
			return false
		}
	}
	return true
}

func partA(lines []string) int {
	valid_lines := make([]int, 0)

	for idx, line := range lines {
		if validateCubes(processLine(line)) {
			valid_lines = append(valid_lines, idx+1)
		}
	}

	sum := 0
	for _, v := range valid_lines {
		sum += v
	}

	return sum
}
