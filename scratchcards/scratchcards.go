package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []int
	myNumbers      []int
	matches        int
	copies         int
}

func parse(fileScanner *bufio.Scanner) []*Card {
	cards := []*Card{}
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

func ParseLine(line string) *Card {
	idAndGame := strings.Split(line, ":")
	numberLists := strings.Split(idAndGame[1], "|")
	return &Card{
		parseNumberList(numberLists[0]),
		parseNumberList(numberLists[1]),
		0,
		1,
	}
}

// ------------------------------------------------------------
// Part 1

func FindScore(card *Card) int {
	card.matches = len(lo.Intersect(card.winningNumbers, card.myNumbers))

	if card.matches > 0 {
		return int(math.Pow(2, float64(card.matches-1)))
	}
	return 0
}

// ------------------------------------------------------------
// Part 2

func AddCopies(cards []*Card, currentIndex int) {
	currentCard := *cards[currentIndex]
	for i := 1; i <= currentCard.matches; i++ {
		copiedCard := *cards[currentIndex+i]
		copiedCard.copies = copiedCard.copies + currentCard.copies
		// Need to reassign after modifying when working with slices of references
		cards[currentIndex+i] = &copiedCard
	}
}

// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("cards.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cards := parse(fileScanner)

	var totalCardScores = 0
	for i, c := range cards {
		// Part 1
		score := FindScore(c)
		totalCardScores = totalCardScores + score

		// Part 2
		if score != 0 {
			AddCopies(cards, i)
		}
	}
	var totalCards = 0
	for _, c := range cards {
		card := *c
		totalCards = totalCards + card.copies
	}

	fmt.Printf("Sum of all card scores: %d\n", totalCardScores)
	fmt.Printf("Total number of cards: %d\n", totalCards)
}
