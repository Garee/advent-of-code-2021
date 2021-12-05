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

func readLines(fpath string) ([]int, [][][]int) {
	file := openFile(fpath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbers := []int{}
	numberStrs := strings.Split(scanner.Text(), ",")
	for _, n := range numberStrs {
		numbers = append(numbers, toInt(n))
	}
	scanner.Scan()

	cards := [][][]int{}

	card := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			cards = append(cards, card)
			card = [][]int{}
		}

		row := make([]int, 5)
		rowRefs := make([]interface{}, 5)
		for i := range row {
			rowRefs[i] = &row[i]
		}

		_, err := fmt.Sscan(line, rowRefs...)
		if err == nil {
			card = append(card, row)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cards = append(cards, card)

	return numbers, cards
}

func isWinner(card [][]int) bool {
	nRow := len(card[0])
	nCol := nRow

	for r := 0; r < nRow; r++ {
		rWin := true
		cWin := true
		for c := 0; c < nCol; c++ {
			if card[r][c] != -1 {
				rWin = false
			}
			if card[c][r] != -1 {
				cWin = false
			}
		}
		if rWin || cWin {
			return true
		}
	}

	return false
}

func findWinningCard(numbers []int, cards [][][]int) (int, int) {
	for _, n := range numbers {
		for j, card := range cards {
			for _, row := range card {
				for l, cell := range row {
					if cell == n {
						row[l] = -1
						if isWinner(card) {
							return j, n
						}
					}
				}
			}
		}
	}

	return -1, -1
}

func sumCard(card [][]int) int {
	sum := 0
	for _, row := range card {
		for _, col := range row {
			if col != -1 {
				sum += col
			}
		}
	}
	return sum
}

func partOne(numbers []int, cards [][][]int) int {
	idx, winningNumber := findWinningCard(numbers, cards)
	if idx != -1 {
		card := cards[idx]
		return sumCard(card) * winningNumber
	}

	return -1
}

func partTwo(numbers []int, cards [][][]int) int {
	var idx int
	var winningNumber int
	for len(cards) > 1 {
		idx, winningNumber = findWinningCard(numbers, cards)
		cards = append(cards[:idx], cards[idx+1:]...)
	}

	idx, winningNumber = findWinningCard(numbers, cards)
	return sumCard(cards[idx]) * winningNumber
}

func main() {
	fpath := parseArgs()
	numbers, cards := readLines(fpath)

	partOneCards := make([][][]int, len(cards))
	copy(partOneCards, cards)
	fmt.Println(partOne(numbers, partOneCards))

	partTwoCards := make([][][]int, len(cards))
	copy(partTwoCards, cards)
	fmt.Println(partTwo(numbers, partTwoCards))
}
