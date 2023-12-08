package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strings"
	"strconv"
)

type Card struct {
	id int
	winningNumbers []int
	myNumbers []int
}

func parse(fileScanner *bufio.Scanner) []Card {
	cards := []Card{}
	for fileScanner.Scan() {
		cards = append(cards, ParseLine(fileScanner.Text()))
	}
	return cards
}

func parseNumberList(list string) []int {
	whiteSpace := regexp.MustCompile(`\s+`)
	split := whiteSpace.Split(list, -1)
	filteredList := lo.Filter(split, func(s string, _ int) bool {
		return s != ""
	})
	return lo.Map(filteredList, func(s string, _ int) int {
		number, _ := strconv.Atoi(s)
		return number
	})
}

func ParseLine(line string) Card {
	idAndGame := strings.Split(line, ":")
	idStr := strings.Replace(idAndGame[0], "Card ", "", -1)
	id, _ := strconv.Atoi(idStr)

	numberLists := strings.Split(idAndGame[1], "|")
	return Card{
		id,
		parseNumberList(numberLists[0]),
		parseNumberList(numberLists[1]),
	}
}

// ------------------------------------------------------------
// Part 1

func FindScore(card Card) int {
	var score = 0
	for _, n := range card.myNumbers {
		if -1 != lo.IndexOf(card.winningNumbers, n) {
			if score == 0 {
				score = 1
			} else {
				score = score * 2
			}
		}
	}
	return score
}

// ------------------------------------------------------------
// Part 2


// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("cards.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cards := parse(fileScanner)

	var totalCardScores = 0
	for _, c := range cards {
		totalCardScores = totalCardScores + FindScore(c)
	}

	fmt.Printf("Sum of all card scores: %d\n", totalCardScores)
}
