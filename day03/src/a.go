package main

func partA(lines []string) uint64 {

	var sum uint64 = 0

	for row := 1; row < len(lines)-1; row += 1 {
		previous := false
		valid_part := false
		var parsed_value uint64 = 0
		for col := 1; col < len(lines[0])-1; col += 1 {
			// check neighbors
			current := isSpecial(lines[row][col]) || isSpecial(lines[row-1][col]) || isSpecial(lines[row+1][col])

			// if is digit
			if isDigit(lines[row][col]) { // construct number
				valid_part = valid_part || previous || current
				parsed_value *= 10
				parsed_value += uint64(lines[row][col] - '0')
			} else if parsed_value != 0 { // if is not digit but we had parsed a number
				valid_part = valid_part || current
				if valid_part {
					sum += parsed_value
				}

				// reset
				parsed_value = 0
				valid_part = false
			}
			previous = current
		}
		// leftovers
		if valid_part && parsed_value != 0 {
			sum += parsed_value
		}
	}

	return sum
}
