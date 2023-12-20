package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func parseBid(s string) int64 {
	var num int64 = 0
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic("parseBid")
	}
	return num
}

// compute score base-14
func computeScore(s string, countJokers bool) int64 {
	var score int64 = 0

	var mapping = CardMapping
	if countJokers {
		mapping = CardMappingJoker
	}

	for _, c := range s {
		score = score*int64(1+len(mapping)) + int64(mapping[c])
	}

	return score
}

func processLine(line string, countJokers bool) Hand {
	h := Hand{
		Cards:     parseCards(line[0:5], countJokers),
		CardScore: computeScore(line[0:5], countJokers),
		Bid:       parseBid(line[6:]),
	}

	h.Analyze(countJokers)

	// fmt.Printf("%+v\n", h)

	return h
}
