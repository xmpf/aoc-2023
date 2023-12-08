package main

func partA(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i++ {
		winning_cards, my_cards := getCards(lines[i])
		sum += 1 << getScore(winning_cards, my_cards) >> 1
	}
	return sum
}
