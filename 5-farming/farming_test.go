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

	destination, source, length := ParseMappeDirectiveLine(line)
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

func TestMapValueInRange(t *testing.T) {
	mappe := Mappe{
		"", "",
		[]Directive{
			Directive{14, 12, 5},
			Directive{5, 7, 5},
			Directive{22, 3, 4},
		},
	}
	input := 9
	want := 11

	got := MapValue(mappe, input)
	if got != want {
		t.Errorf("Mapping within directive range incorrect. Got: %d\n", got)
	}
}

func TestMapValueAtStartOfRange(t *testing.T) {
	mappe := Mappe{
		"", "",
		[]Directive{
			Directive{14, 12, 5},
			Directive{5, 7, 5},
			Directive{22, 3, 4},
		},
	}
	input := 22
	want := 3

	got := MapValue(mappe, input)
	if got != want {
		t.Errorf("Mapping at start of directive range incorrect. Got: %d\n", got)
	}
}

func TestMapValueAtEndOfRange(t *testing.T) {
	mappe := Mappe{
		"", "",
		[]Directive{
			Directive{14, 12, 5},
			Directive{5, 7, 5},
			Directive{22, 3, 4},
		},
	}
	input := 10
	want := 12

	got := MapValue(mappe, input)
	if got != want {
		t.Errorf("Mapping at end of directive range incorrect. Got: %d\n", got)
	}
}

func TestMapValueOutsideRange(t *testing.T) {
	mappe := Mappe{
		"", "",
		[]Directive{
			Directive{14, 12, 5},
			Directive{5, 7, 5},
			Directive{22, 3, 4},
		},
	}
	input := 2
	want := 2

	got := MapValue(mappe, input)
	if got != want {
		t.Errorf("Default mapping (value outside of directive ranges) incorrect. Got: %d\n", got)
	}
}

func TestFindMultiplyMappedValue(t *testing.T) {
	mappe1 := Mappe{
		"a", "b",
		[]Directive{
			Directive{14, 12, 5},
			Directive{5, 7, 5},
			Directive{22, 3, 4},
		},
	}
	mappe2 := Mappe{
		"b", "c",
		[]Directive{
			Directive{14, 10, 5},
			Directive{5, 7, 3},
			Directive{22, 3, 4},
		},
	}
	mapOfMappes := map[string]Mappe{
		"b": mappe2,
		"a": mappe1,
	}
	input := 16
	want := 10

	got := FindMappedValue(mapOfMappes, input, "a", "c")
	if got != want {
		t.Errorf("Mapping through two mappes incorrect. Got: %d\n", got)
	}
}

func TestFindLowest(t *testing.T) {
	values := []int{123, 33, 12, 9999}
	want := 12

	got := FindLowest(values)
	if got != want {
		t.Errorf("Failed to find lowest value. Got: %d\n", got)
	}
}

// ------------------------------------------------------------
// Part 2
