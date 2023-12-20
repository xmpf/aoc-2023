package main

import (
	"sort"
)

type CardType int

const (
	Joker CardType = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type HandType int

const (
	_ HandType = iota
	HighCard
	SinglePair
	TwoPairs
	ThreeKinds
	FullHouse
	FourKinds
	FiveKinds
)

var CardMapping = map[rune]CardType{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Jack,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

var CardMappingJoker = map[rune]CardType{
	'J': Joker,
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

type Hand struct {
	Cards     [5]CardType
	CardScore int64
	Type      HandType
	Bid       int64
}

func parseCards(s string, countJokers bool) [5]CardType {
	var cards [5]CardType

	var mapping = CardMapping
	if countJokers {
		mapping = CardMappingJoker
	}

	for ix, c := range s {
		cards[ix] = mapping[c]
	}

	return cards
}

func (h *Hand) Analyze(countJokers bool) {

	sorted_cards := h.Cards

	// sort descending
	sort.SliceStable(sorted_cards[:], func(i, j int) bool {
		return sorted_cards[i] > sorted_cards[j]
	})

	// if the highest cars is Joker
	if countJokers && sorted_cards[0] == Joker {
		h.Type = FiveKinds
		return
	}

	// compute frequencies
	freq := [5]int{1, 0, 0, 0, 0}

	jokers := 0
	j := 0
	for i := 1; i < 5; i += 1 {
		// joker(s) will always be at positions [1, 4]
		if countJokers && sorted_cards[i] == Joker {
			jokers += 1
			continue
		}

		if sorted_cards[i-1] != sorted_cards[i] {
			j += 1
		}

		freq[j] += 1
	}

	// sort descending
	sort.SliceStable(freq[:], func(i, j int) bool {
		return freq[i] > freq[j]
	})

	// fmt.Printf("%+v\n", freq)

	if freq[0] == 5 {
		h.Type = FiveKinds
	} else if freq[0] == 4 {
		h.Type = FourKinds
		if jokers == 1 {
			h.Type = FiveKinds
		}
	} else if freq[0] == 3 && freq[1] == 2 {
		h.Type = FullHouse
	} else if freq[0] == 3 && freq[1] != 2 {
		h.Type = ThreeKinds
		if jokers == 1 {
			h.Type = FourKinds
		} else if jokers == 2 {
			h.Type = FiveKinds
		}
	} else if freq[0] == 2 && freq[1] == 2 {
		h.Type = TwoPairs
		if jokers == 1 {
			h.Type = FullHouse
		}
	} else if freq[0] == 2 && freq[1] != 2 {
		h.Type = SinglePair
		if jokers == 1 {
			h.Type = ThreeKinds
		} else if jokers == 2 {
			h.Type = FourKinds
		} else if jokers == 3 {
			h.Type = FiveKinds
		}
	} else {
		h.Type = HighCard
		if jokers == 1 {
			h.Type = SinglePair
		} else if jokers == 2 {
			h.Type = ThreeKinds
		} else if jokers == 3 {
			h.Type = FourKinds
		} else if jokers == 4 {
			h.Type = FiveKinds
		}
	}
}

func solve(lines []string, countJokers bool) int64 {
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand := processLine(line, countJokers)
		hands = append(hands, hand)
	}

	// sort ascending
	sort.SliceStable(hands, func(i, j int) bool {

		if hands[i].Type == hands[j].Type {
			return hands[i].CardScore < hands[j].CardScore
		}

		return hands[i].Type < hands[j].Type
	})

	var prod int64 = 0
	for idx, hand := range hands {
		prod = prod + hand.Bid*int64(idx+1)
	}

	return prod
}
