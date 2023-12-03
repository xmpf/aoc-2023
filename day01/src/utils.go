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

// first int
func firstInt(s string, mapping map[string]int) int {
	if s == "" {
		return 0
	}

	for key, value := range mapping {
		if strings.HasPrefix(s, key) {
			return value
		}
	}

	return firstInt(s[1:], mapping)
}

// last int
func lastInt(s string, mapping map[string]int) int {
	if s == "" {
		return 0
	}

	for key, value := range mapping {
		if strings.HasSuffix(s, key) {
			return value
		}
	}

	return lastInt(s[0:len(s)-1], mapping)
}
