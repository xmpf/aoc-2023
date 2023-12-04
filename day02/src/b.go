package main

func partB(lines []string) int {
	power_line := []int{}
	for _, line := range lines {
		prod := 1
		cubes := processLine(line)
		for _, v := range cubes {
			prod *= max(v)
		}
		power_line = append(power_line, prod)
	}

	sum := 0
	for _, v := range power_line {
		sum += v
	}

	return sum
}
