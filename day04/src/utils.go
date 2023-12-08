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

func getCards(line string) (map[int]struct{}, map[int]struct{}) {

	winning_numbers := make(map[int]struct{}, 0)
	my_numbers := make(map[int]struct{}, 0)

	cards := strings.Split(line, ":")
	cards = strings.Split(cards[1], "|")

	winning_cards := strings.Trim(cards[0], " ")
	for _, number := range strings.Split(strings.Trim(winning_cards, " "), " ") {
		trimmed := strings.Trim(number, " ")
		if trimmed == "" {
			continue
		}
		parsed_number, _ := strconv.Atoi(trimmed)
		winning_numbers[parsed_number] = struct{}{}
	}

	my_cards := strings.Trim(cards[1], " ")
	for _, number := range strings.Split(strings.Trim(my_cards, " "), " ") {
		trimmed := strings.Trim(number, " ")
		if trimmed == "" {
			continue
		}
		parsed_number, _ := strconv.Atoi(trimmed)
		my_numbers[parsed_number] = struct{}{}
	}

	return winning_numbers, my_numbers
}

func getScore(winning_cards, my_cards map[int]struct{}) int {
	score := 0
	for k, _ := range my_cards {
		_, ok := winning_cards[k]
		if ok {
			score += 1
		}
	}

	return score
}
