package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
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

func parseLine(line string) map[string][]int {
	cubes := map[string][]int{
		"red":   {0},
		"green": {0},
		"blue":  {0},
	}

	// parse red
	r := regexp.MustCompile(`(\d+) red`)
	match := r.FindAllStringSubmatch(line, -1)
	for i := 0; i < len(match); i += 1 {
		v, _ := strconv.Atoi(match[i][1])
		cubes["red"] = append(cubes["red"], v)
	}

	// parse green
	r = regexp.MustCompile(`(\d+) green`)
	match = r.FindAllStringSubmatch(line, -1)
	for i := 0; i < len(match); i += 1 {
		v, _ := strconv.Atoi(match[i][1])
		cubes["green"] = append(cubes["green"], v)
	}

	// parse blue
	r = regexp.MustCompile(`(\d+) blue`)
	match = r.FindAllStringSubmatch(line, -1)
	for i := 0; i < len(match); i += 1 {
		v, _ := strconv.Atoi(match[i][1])
		cubes["blue"] = append(cubes["blue"], v)
	}

	return cubes
}

func max(lst []int) int {
	ret := lst[0]
	for _, v := range lst[1:] {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func processLine(line string) map[string][]int {
	cubes := parseLine(line)
	return cubes
}
