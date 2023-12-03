package main

func partB(buf []string) int {

	mapping := map[string]int{
		"0":     0,
		"zero":  0,
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}

	sum := 0
	for _, line := range buf {
		sum += firstInt(line, mapping)*10 + lastInt(line, mapping)
	}
	return sum
}
