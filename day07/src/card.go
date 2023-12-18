package main

import (
	"sort"
)

type CardType int

const (
	Joker CardType = iota
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
	Line      string // for debugging
	Cards     [5]CardType
	CardScore int64
	Type      HandType
	Bid       int64
}

func (h *Hand) Analyze() {

	sorted_cards := h.Cards

	// sort descending
	sort.SliceStable(sorted_cards[:], func(i, j int) bool {
		return sorted_cards[i] > sorted_cards[j]
	})

	// compute frequencies
	freq := [5]int{1, 0, 0, 0, 0}

	j := 0
	for i := 1; i < 5; i += 1 {
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
	} else if freq[0] == 3 && freq[1] == 2 {
		h.Type = FullHouse
	} else if freq[0] == 3 && freq[1] != 2 {
		h.Type = ThreeKinds
	} else if freq[0] == 2 && freq[1] == 2 {
		h.Type = TwoPairs
	} else if freq[0] == 2 && freq[1] != 2 {
		h.Type = SinglePair
	} else {
		h.Type = HighCard
	}
}
