package main

import (
	"testing"
)

func TestGettingNumericalDigitFromLocation(t *testing.T) {
	line := "ab2cd3ef4"
	location := []int{2, 3}
	want := 2

	got := GetDigitFromMatchedLocation(line, location)
	if got != want {
		t.Errorf("Didn't find the digit at the location correctly. Found: %d\n", got)
	}
}

func TestGettingTextDigitFromLocation(t *testing.T) {
	line := "abone2cd3ef4"
	location := []int{2, 5}
	want := 1

	got := GetDigitFromMatchedLocation(line, location)
	if got != want {
		t.Errorf("Didn't find the digit at the location correctly. Found: %d\n", got)
	}
}

func TestGetFirstAndLastDigit(t *testing.T) {
	line := "ab2cd3ef4"
	wantFirst := 2
	wantLast := 4

	first, last := GetFirstAndLastNumbers(line)
	if first != wantFirst {
		t.Errorf("Didn't get the desired first number. Got %d\n", first)
	}
	if last != wantLast {
		t.Errorf("Didn't get the desired last number. Got %d\n", last)
	}
}

func TestGetFirstAndLastTextDigit(t *testing.T) {
	line := "abthree2cd3ef4nine"
	wantFirst := 3
	wantLast := 9

	first, last := GetFirstAndLastNumbers(line)
	if first != wantFirst {
		t.Errorf("Didn't get the desired first number. Got %d\n", first)
	}
	if last != wantLast {
		t.Errorf("Didn't get the desired last number. Got %d\n", last)
	}
}

func TestGetFirstAndLastTextDigitWithOverlap(t *testing.T) {
	line := "abthree2cd3ef4nineight"
	wantFirst := 3
	wantLast := 8

	first, last := GetFirstAndLastNumbers(line)
	if first != wantFirst {
		t.Errorf("Didn't get the desired first number. Got %d\n", first)
	}
	if last != wantLast {
		t.Errorf("Didn't get the desired last number. Got %d\n", last)
	}
}

func TestConcatDigits(t *testing.T) {
	one := 2
	two := 5
	want := 25

	got := ConcatTwoDigits(one, two)
	if got != want {
		t.Errorf("Didn't concat digits correctly. Got %d\n", got)
	}
}
