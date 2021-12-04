package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	pos := 0
	depth := 0
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		dir := tokens[0]
		units := toInt(tokens[1])
		switch dir {
		case "forward":
			pos += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}
	return pos * depth
}

func partTwo(lines []string) int {
	pos := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		dir := tokens[0]
		units := toInt(tokens[1])
		switch dir {
		case "forward":
			pos += units
			depth += units * aim
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	return pos * depth
}

func main() {
	fpath := parseArgs()
	lines := readLines(fpath)
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}
