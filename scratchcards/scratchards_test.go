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
		t.Error("Failed to parse winning numbers")
	}
}

// ------------------------------------------------------------
// Part 1
