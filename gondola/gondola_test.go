package main

import (
	"testing"
)

func TestPartWithinWidthIsRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5,}
	symbolLocation := SymbolLocation{3, 4,}

	if !IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Failed to find within width symbol")
	}
}

func TestPartOneLeftOfWidthIsRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5,}
	symbolLocation := SymbolLocation{3, 2,}

	if !IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Failed to find symbol left of part")
	}
}

func TestPartOneRightOfWidthIsRelevant(t *testing.T) {
	part := Part{123, 4, 3, 5,}
	symbolLocation := SymbolLocation{3, 6,}

	if !IsPartNumberRelevant(part, []SymbolLocation{symbolLocation}) {
		t.Error("Failed to find symbol left of part")
	}
}