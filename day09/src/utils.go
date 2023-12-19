package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

func processLine(line string) []int64 {
	result := make([]int64, 0)

	for _, n := range strings.Split(line, " ") {
		number, err := strconv.Atoi(n)
		if err != nil {
			panic("processLine: strconv.Atoi()")
		}
		result = append(result, int64(number))
	}

	return result
}

func extrapolateSequence(sequence []int64) int64 {
	diffs := make([][]int64, 0)
	diffs = append(diffs, sequence)

	done := false
	for !done {
		done = true
		diff := make([]int64, 0)
		for i := 1; i < len(sequence); i += 1 {
			delta := sequence[i] - sequence[i-1]
			if delta != 0 {
				done = false
			}
			diff = append(diff, delta)
		}
		sequence = diff
		diffs = append(diffs, diff)
	}

	// fmt.Printf("%+v\n", diffs)

	diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)
	for i := len(diffs) - 2; i >= 0; i -= 1 {
		prevSeq := diffs[i+1]
		seq := diffs[i]

		extrapolate := seq[len(seq)-1] + prevSeq[len(prevSeq)-1]
		diffs[i] = append(diffs[i], extrapolate)
	}

	// fmt.Printf("%+v\n", diffs)

	return diffs[0][len(diffs[0])-1]
}

func listReverse(input []int64) []int64 {
	if len(input) == 0 {
		return input
	}
	return append(listReverse(input[1:]), input[0])
}

func solve(lines []string, reverse bool) int64 {
	var sum int64 = 0
	for _, line := range lines {
		sequence := processLine(line)
		if reverse {
			sequence = listReverse(sequence)
		}
		sum += extrapolateSequence(sequence)
	}

	return sum
}
