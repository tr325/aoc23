package main

import (
	"testing"
)


func TestParseLine(t *testing.T) {
	line := "....123...$.."
	row := 10
	parts := []Part{}
	symbols := []SymbolData{}

	ParseLine(line, &parts, &symbols, row)
	if len(parts) != 1 {
		t.Error("Failed to find part")
	}
	if parts[0].row != row {
		t.Errorf("Part placed in wrong row. got: %d\n", parts[0].row)
	}
	if parts[0].start != 4 {
		t.Errorf("Part start incorrect. got: %d\n", parts[0].start)
	}
	if parts[0].end != 6 {
		t.Errorf("Part end incorrect. got: %d\n", parts[0].end)
	}
	if len(symbols) != 1 {
		t.Error("Failed to find symbol")
	}
	if symbols[0].symbol != "$" {
		t.Errorf("Found incorrect symbol. Got: %s\n", symbols[0].symbol)
	}
	if symbols[0].location.row != row {
		t.Errorf("Found incorrect row for symbol. Got: %d\n", symbols[0].location.row)
	}
	if symbols[0].location.col != 10 {
		t.Errorf("Found incorrect column for symbol. Got: %d\n", symbols[0].location.col)
	}
}

// ------------------------------------------------------------
// Part 1

func TestPartWithinWidthIsRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{3, 4}

	if !IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Failed to find within width symbol")
	}
}

func TestPartOneLeftOfWidthIsRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{3, 2}

	if !IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Failed to find symbol left of part")
	}
}

func TestPartOneRightOfWidthIsRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{3, 6}

	if !IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Failed to find symbol right of part")
	}
}

func TestPartTwoRightOfWidthIsNotRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{3, 7}

	if IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Included number two left of symbol")
	}
}

func TestPartTwoLeftOfWidthIsNotRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{3, 1}

	if IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Included number two right of symbol")
	}
}

func TestPartTwoBelowIsNotRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{2, 4}

	if IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Included number two below symbol")
	}
}

func TestPartTwoAboveIsNotRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5}
	symbolLocation := SymbolLocation{6, 4}

	if IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Included number two above symbol")
	}
}

// ------------------------------------------------------------
// Part 2

func TestFindSingleAdjacentPart(t *testing.T) {
	parts := []Part{Part{123, 4, 3, 5}}
	symbol := SymbolData{[]Part{}, "*", SymbolLocation{3, 2}}
	want := 1

	AddAdjacentParts(&symbol, parts)
	if len(symbol.adjacentParts) != want {
		t.Error("Failed to find adjacent part")
	}
}

// To make sure the reference isn't overwriting the adjacentParts array
func TestFindTwoAdjacentParts(t *testing.T) {
	parts := []Part{Part{123, 4, 3, 5}, Part{333, 2, 1, 3}}
	symbol := SymbolData{[]Part{}, "*", SymbolLocation{3, 2}}
	want := 2

	AddAdjacentParts(&symbol, parts)
	if len(symbol.adjacentParts) != want {
		t.Error("Failed to find adjacent parts")
	}
}

func TestFindTwoOfThreeAdjacentParts(t *testing.T) {
	parts := []Part{Part{123, 4, 3, 5}, Part{333, 12, 1, 3}, Part{333, 2, 1, 3}}
	symbol := SymbolData{[]Part{}, "*", SymbolLocation{3, 2}}
	want := 2

	AddAdjacentParts(&symbol, parts)
	got := len(symbol.adjacentParts)
	if got != want {
		t.Errorf("Failed to find correct number of adjacent parts. Got %d\n", got)
	}
}

func TestIsGear(t *testing.T) {
	symbol := SymbolData{
		[]Part{Part{123, 4, 3, 5}, Part{333, 2, 1, 3}},
		"*",
		SymbolLocation{3, 2},
	}

	if !IsGear(symbol) {
		t.Error("Failed to identify gear")
	}
}

func TestGearIsWrongSymbol(t *testing.T) {
	symbol := SymbolData{
		[]Part{Part{123, 4, 3, 5}, Part{333, 2, 1, 3}},
		"$",
		SymbolLocation{3, 2},
	}

	if IsGear(symbol) {
		t.Error("Incorrectly identified $ symbol as gear")
	}
}

func TestGearHasOnlyOneAdjacentPart(t *testing.T) {
	symbol := SymbolData{
		[]Part{Part{123, 4, 3, 5}},
		"*",
		SymbolLocation{3, 2},
	}

	if IsGear(symbol) {
		t.Error("Incorrectly identified symbol with one ajacent part as gear")
	}
}

func TestFindsPartsAndIsGear(t *testing.T) {
	parts := []Part{Part{123, 4, 3, 5}, Part{333, 2, 1, 3}}
	symbol := SymbolData{[]Part{}, "*", SymbolLocation{3, 2}}

	AddAdjacentParts(&symbol, parts)
	if !IsGear(symbol) {
		t.Error("Failed to find adjacent parts and identify as gear")
	}
}

func TestGetGearRatio(t *testing.T) {
	symbol := SymbolData{
		[]Part{Part{12, 4, 3, 5}, Part{30, 2, 1, 3}},
		"*",
		SymbolLocation{3, 2},
	}
	want := 360

	got := GetGearRatio(symbol)
	if want != got {
		t.Errorf("Got incorrect gear ratio. Got: %d\n", got)
	}
}
