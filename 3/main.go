package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseArgs() string {
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " <input_file>")
		os.Exit(1)
	}

	return os.Args[1]
}

func openFile(fpath string) *os.File {
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println("failed to open file '" + fpath + "'")
		os.Exit(1)
	}
	return file
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func toDecimal(b string) int64 {
	n, _ := strconv.ParseInt(b, 2, 64)
	return n
}

func readLines(fpath string) []string {
	file := openFile(fpath)
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return lines
}

func getBitCounts(lines []string) map[int]map[rune]int {
	bitCounts := map[int]map[rune]int{}
	for _, line := range lines {
		for i, bit := range line {
			if _, ok := bitCounts[i]; !ok {
				bitCounts[i] = map[rune]int{}
			}
			bitCounts[i][bit] += 1
		}
	}
	return bitCounts
}

func getMostCommonBit(pos int, bitCounts map[int]map[rune]int) byte {
	if bitCounts[pos]['0'] > bitCounts[pos]['1'] {
		return '0'
	}
	return '1'
}

func getLeastCommonBit(pos int, bitCounts map[int]map[rune]int) byte {
	if bitCounts[pos]['0'] > bitCounts[pos]['1'] {
		return '1'
	}
	return '0'
}

func partOne(lines []string) int64 {
	bitCounts := getBitCounts(lines)

	gamma := ""
	epsilon := ""
	for i := 0; i < len(bitCounts); i++ {
		counts := bitCounts[i]
		if counts['1'] > counts['0'] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	return toDecimal(gamma) * toDecimal(epsilon)
}

func getRating(lines []string, getCommonBit func(pos int, bitCounts map[int]map[rune]int) byte) int64 {
	for pos := 0; len(lines) > 1; pos++ {
		bitCounts := getBitCounts(lines)
		commonBit := getCommonBit(pos, bitCounts)

		numbers := []string{}
		for _, line := range lines {
			bit := line[pos]
			if bit == commonBit {
				numbers = append(numbers, line)
			}
		}

		lines = make([]string, len(numbers))
		copy(lines, numbers)
	}

	rating := lines[0]
	return toDecimal(rating)
}

func partTwo(lines []string) int64 {
	oxygen := getRating(lines, getMostCommonBit)
	co2 := getRating(lines, getLeastCommonBit)
	return oxygen * co2
}

func main() {
	fpath := parseArgs()
	lines := readLines(fpath)
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}
