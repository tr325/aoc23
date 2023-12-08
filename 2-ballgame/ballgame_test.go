package main

import (
	"testing"
)

func TestGetIdFromGameTitle(t *testing.T) {
	gameRecord := "Game 13"
	want := 13

	got, _ := GetIdForGame(gameRecord)
	if got != want {
		t.Errorf("Didn't find the game id correctly. Found: %d\n", got)
	}
}

func TestHandfulPossible(t *testing.T) {
	handful := "4 blue, 1 green, 2 red"
	want := true

	if want != IsGamePossible(handful) {
		t.Error("Valid handful registered as invalid")
	}
}

func TestHandfulPossibleLimits(t *testing.T) {
	handful := "14 blue, 13 green, 12 red"
	want := true

	if want != IsGamePossible(handful) {
		t.Error("Limit-valid handful registered as invalid")
	}
}

func TestGetGamePossibleFromGameScore(t *testing.T) {
	gameScore := "4 blue, 1 green, 2 red; 5 red, 11 blue, 6 green; 9 green, 11 blue;"
	want := true

	got := IsGamePossible(gameScore)
	if got != want {
		t.Error("Game is possible, but wasn't reported as such")
	}
}

func TestGetGameImpossibleTooManyRed(t *testing.T) {
	gameScore := "4 blue, 1 green, 13 red; 5 red, 11 blue, 6 green; 9 green, 11 blue;"
	want := false

	got := IsGamePossible(gameScore)
	if got != want {
		t.Error("Game is impossible (too many red) but wasn't reported as such")
	}
}

func TestGetGameImpossibleTooManyGreen(t *testing.T) {
	gameScore := "4 blue, 14 green, 2 red; 5 red, 11 blue, 6 green; 9 green, 11 blue;"
	want := false

	got := IsGamePossible(gameScore)
	if got != want {
		t.Error("Game is impossible (too many green) but wasn't reported as such")
	}
}

func TestGetGameImpossibleTooManyBlue(t *testing.T) {
	gameScore := "4 blue, 1 green, 2 red; 5 red, 11 blue, 6 green; 9 green, 15 blue;"
	want := false

	got := IsGamePossible(gameScore)
	if got != want {
		t.Error("Game is impossible (too many blue) but wasn't reported as such")
	}
}

// ------------------------------------------------
// Part 2

func TestGetGamePower(t *testing.T) {
	gameScore := "4 blue, 1 green, 2 red; 5 red, 3 blue, 6 green; 2 green, 3 blue;"
	want := 120

	got := GetGamePower(gameScore)
	if got != want {
		t.Errorf("Got incorrect game power. Got: %d\n", got)
	}
}
