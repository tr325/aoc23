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

func TestParseMappeTitleLine(t *testing.T) {
	line := "water-to-light map:"
	wantSource := "water"
	wantDestination := "light"

	gotSource, gotDestination := ParseMappeTitleLine(line)
	if gotSource != wantSource {
		t.Errorf("Failed to parse Mappe source. Got: %s\n", gotSource)
	}
	if gotDestination != wantDestination {
		t.Errorf("Failed to parse Mappe destination. Got: %s\n", gotDestination)
	}
}

func TestParseMappeDirectiveLine(t *testing.T) {
	line := "12 13 14"
	wantSourceStart := 13
	wantDestinationStart := 12
	wantRangeLength := 14

	source, destination, length := ParseMappeDirectiveLine(line)
	if source != wantSourceStart {
		t.Errorf("Failed to parse Mappe source range start. Got: %d\n", source)
	}
	if destination != wantDestinationStart {
		t.Errorf("Failed to parse Mappe destination range start. Got: %d\n", destination)
	}
	if length != wantRangeLength {
		t.Errorf("Failed to parse Mappe range length. Got: %d\n", length)
	}
}

// ------------------------------------------------------------
// Part 1

// ------------------------------------------------------------
// Part 2
