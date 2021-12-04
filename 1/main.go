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

func partOne(lines []string) int {
	count := 0
	for i := 1; i < len(lines); i++ {
		measurement := toInt(lines[i])
		prev := toInt(lines[i-1])
		if measurement > prev {
			count++
		}
	}
	return count
}

func partTwo(lines []string) int {
	count := 0
	for i := 0; i < len(lines)-3; i++ {
		prev := toInt(lines[i]) + toInt(lines[i+1]) + toInt(lines[i+2])
		window := toInt(lines[i+1]) + toInt(lines[i+2]) + toInt(lines[i+3])
		if window > prev {
			count++
		}
	}
	return count
}

func main() {
	fpath := parseArgs()
	lines := readLines(fpath)
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}
