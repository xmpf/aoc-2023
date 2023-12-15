package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func readFile(filename string) []string {
	input_file, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error opening input file!")
	}

	defer input_file.Close()

	fileScanner := bufio.NewScanner(input_file)
	fileScanner.Split(bufio.ScanLines)

	buf := make([]string, 0)
	for fileScanner.Scan() {
		buf = append(buf, fileScanner.Text())
	}

	return buf
}

func processLine(line string) []int {
	numbers := make([]int, 0)

	number := 0
	for i := 0; i < len(line); i += 1 {
		if line[i] >= '0' && line[i] <= '9' {
			number = number*10 + int(line[i]-'0')
			continue
		}

		if number != 0 {
			numbers = append(numbers, number)
			number = 0
		}
	}

	if number != 0 {
		numbers = append(numbers, number)
		number = 0
	}

	return numbers
}

func processLine_b(line string) int {
	number := 0
	for i := 0; i < len(line); i += 1 {
		if line[i] >= '0' && line[i] <= '9' {
			number = number*10 + int(line[i]-'0')
			continue
		}
	}

	return number
}

func solve_quadratic(a, b, c float64) (float64, float64) {
	delta := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	x1 := (-b + delta) / (2 * a)
	x2 := (-b - delta) / (2 * a)
	return x1, x2
}
