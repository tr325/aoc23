package main

import (
	"bufio"
	"github.com/samber/lo"
	"os"
	"strings"
	"strconv"
)

type Card struct {
	id int
	winningNumbers []int
	myNumbers []int
}

func parse(fileScanner *bufio.Scanner) []Card {
	// TODO
	return []Card{}
}

func parseNumberList(list string) []int {
	return lo.Map(strings.Split(list, " "), func(s string, _ int) int {
		number, _ := strconv.Atoi(s)
		return number
	})
}

func ParseLine(line string) Card {
	idAndGame := strings.Split(line, ": ")
	idStr := strings.Replace(idAndGame[0], "Card ", "", -1)
	id, _ := strconv.Atoi(idStr)

	numberLists := strings.Split(idAndGame[1], " | ")
	return Card{
		id,
		parseNumberList(numberLists[0]),
		parseNumberList(numberLists[1]),
	}
}

// ------------------------------------------------------------
// Part 1

// ------------------------------------------------------------
// Part 2


// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("cards.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// cards := parse(fileScanner)

	// fmt.Printf("Sum of all card scores: %d\n", cardScoreSum)
}
