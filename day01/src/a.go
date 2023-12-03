package main

func partA(buf []string) int {

	mapping := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	sum := 0
	for _, line := range buf {
		sum += firstInt(line, mapping)*10 + lastInt(line, mapping)
	}
	return sum
}
