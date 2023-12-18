package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
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
	for _, c := range s {
		if unicode.IsDigit(c) {
			num = num*10 + int64(c-'0')
		}
	}
	return num
}

func computeScore(s string) int64 {
	var score int64 = 0
	for _, c := range s {
		score = score*int64(len(CardMapping)) + int64(CardMapping[c])
	}
	return score
}

func processLine(line string) Hand {
	h := Hand{
		Line: line,
		Cards: func(s string) [5]CardType {
			var cards [5]CardType
			for ix, c := range s {
				cards[ix] = CardMapping[c]
			}
			return cards
		}(line[0:5]),
		CardScore: computeScore(line[0:5]),
		Bid:       parseBid(line[6:]),
	}

	h.Analyze()

	// fmt.Printf("%+v\n", h)

	return h
}
