package main

import (
	"bufio"
	"log"
	"os"
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

// pad lines => dont care about index checking anymore
func padLines(lines []string) []string {
	ret := make([]string, 0)

	padding := make([]byte, 0)
	for i := 0; i < len(lines[0])+2; i += 1 {
		padding = append(padding, '.')
	}
	ret = append(ret, string(padding))
	for _, line := range lines {
		ret = append(ret, strings.Join([]string{".", line, "."}, ""))
	}
	ret = append(ret, string(padding))

	return ret
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSpecial(c byte) bool {
	return !isDigit(c) && c != '.'
}
