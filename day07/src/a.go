package main

import (
	"sort"
)

func partA(lines []string) int64 {
	hands := make([]Hand, 0)
	for _, line := range lines {
		hands = append(hands, processLine(line))
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
