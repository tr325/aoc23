package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "Card 123: 44 22 11 | 11 22 33"
	winning := []int{44, 22, 11}
	mine := []int{11, 22, 33}

	card := ParseLine(line)
	if !cmp.Equal(card.winningNumbers, winning) {
		t.Error("Failed to parse winning numbers")
	}
	if !cmp.Equal(card.myNumbers, mine) {
		t.Error("Failed to parse my numbers")
	}
}

// ------------------------------------------------------------
// Part 1

func TestFindScoreSingleMatch(t *testing.T) {
	card := Card{[]int{1, 2, 3}, []int{2, 5, 6}, 0, 1}
	want := 1

	got := FindScore(&card)
	if want != got {
		t.Errorf("Got incorrect score for single winning number. Got: %d\n", got)
	}
}

func TestFindScoreAddsMatchesToCard(t *testing.T) {
	card := Card{[]int{1, 2, 3}, []int{2, 5, 6}, 0, 1}
	want := 1

	FindScore(&card)
	got := card.matches
	if want != got {
		t.Errorf("Did not add matches correctly to card. Got: %d\n", got)
	}
}

func TestFindScoreNoMatch(t *testing.T) {
	card := Card{[]int{1, 2, 3}, []int{12, 5, 6}, 0, 1}
	want := 0

	got := FindScore(&card)
	if want != got {
		t.Errorf("Got incorrect score for no winning numbers. Got: %d\n", got)
	}
}

func TestFindScoreMultipleMatches(t *testing.T) {
	card := Card{[]int{1, 2, 3}, []int{2, 3, 3, 5}, 0, 1}
	want := 4

	got := FindScore(&card)
	if want != got {
		t.Errorf("Got incorrect score for multiple winning numbers. Got: %d\n", got)
	}
}

// ------------------------------------------------------------
// Part 2

func TestAddingSingleMatch(t *testing.T) {
	cards := []*Card{
		&Card{[]int{}, []int{}, 1, 1},
		&Card{[]int{}, []int{}, 0, 1},
	}
	want := 2

	AddCopies(cards, 0)
	got := cards[1].copies
	if got != want {
		t.Errorf("Failed to add single copy. Number of copies: %d\n", got)
	}
}

func TestAddingSingleMatchForMultipleOriginals(t *testing.T) {
	cards := []*Card{
		&Card{[]int{}, []int{}, 1, 2},
		&Card{[]int{}, []int{}, 0, 1},
	}
	want := 3 // 2 copies of [0], 1 match --> 1 original of [1] + 2

	AddCopies(cards, 0)
	got := cards[1].copies
	if got != want {
		t.Errorf("Failed to add two copies when two originals with one match. Number of copies: %d\n", got)
	}
}

func TestAddingMultipleMatches(t *testing.T) {
	cards := []*Card{
		&Card{[]int{}, []int{}, 2, 1},
		&Card{[]int{}, []int{}, 0, 1},
		&Card{[]int{}, []int{}, 0, 1},
	}
	want := 2

	AddCopies(cards, 0)
	got1 := cards[1].copies
	got2 := cards[2].copies
	if got1 != want {
		t.Errorf("Failed to add first copy. Number of copies: %d\n", got1)
	}
	if got2 != want {
		t.Errorf("Failed to add second copy. Number of copies: %d\n", got2)
	}
}

func TestAddingSequentialMatches(t *testing.T) {
	cards := []*Card{
		&Card{[]int{}, []int{}, 2, 1},
		&Card{[]int{}, []int{}, 1, 1},
		&Card{[]int{}, []int{}, 0, 1},
	}
	// 2 matches in [0] -> 2 copies of [1] and [2]
	// (1 match in [1]) * 2 copies --> 2 more copies of [2] --> 4 copies of [2]
	want := 4

	for i, _ := range cards {
		AddCopies(cards, i)
	}
	got := cards[2].copies
	if got != want {
		t.Errorf("Failed to add sequential matches. Number of copies: %d\n", got)
	}
}
