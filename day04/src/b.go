package main

func getScoreAlt(scratchcards map[int]int, position int, winning_numbers, my_numbers map[int]struct{}) map[int]int {
	score := getScore(winning_numbers, my_numbers)
	for i := position + 1; i <= position+score; i++ {
		scratchcards[i] += scratchcards[position]
	}
	return scratchcards
}

func partB(lines []string) int {
	sum := 0
	scratchcards := make(map[int]int, 0)
	for i := 0; i < len(lines); i++ {
		scratchcards[i] += 1
		winning_numbers, my_numbers := getCards(lines[i])
		scratchcards = getScoreAlt(scratchcards, i, winning_numbers, my_numbers)
	}

	for _, v := range scratchcards {
		sum += v
	}
	return sum
}
