package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)


func TestParseLine(t *testing.T) {
	line := "Card 123: 44 22 11 | 11 22 33"
	id := 123
	winning := []int{44, 22, 11}
	mine := []int{11, 22, 33}

	card := ParseLine(line)
	if card.id != id {
		t.Errorf("Parsed card id incorrectly. Got: %d\n", card.id)
	}
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
	card := Card{1, []int{1, 2, 3}, []int{2, 5, 6}, 1}
	want := 1

	got := FindScore(card)
	if want != got {
		t.Errorf("Got incorrect score for single winning number. Got: %d\n", got)
	}
}

func TestFindScoreNoMatch(t *testing.T) {
	card := Card{1, []int{1, 2, 3}, []int{12, 5, 6}, 1}
	want := 0

	got := FindScore(card)
	if want != got {
		t.Errorf("Got incorrect score for no winning numbers. Got: %d\n", got)
	}
}

func TestFindScoreMultipleMatches(t *testing.T) {
	card := Card{1, []int{1, 2, 3}, []int{2, 3, 3, 5}, 1}
	want := 4

	got := FindScore(card)
	if want != got {
		t.Errorf("Got incorrect score for multiple winning numbers. Got: %d\n", got)
	}
}