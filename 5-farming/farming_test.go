package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParseSeedsInput(t *testing.T) {
	line := "seeds: 123 4 56"
	want := []int{123, 4, 56}

	got := ParseSeedsInputLine(line)
	if !cmp.Equal(got, want) {
		t.Error("Failed to parse seed input line")
	}
}

// ------------------------------------------------------------
// Part 1

// ------------------------------------------------------------
// Part 2
