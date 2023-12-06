package main

import (
	"testing"
)

func TestGetIdFromGameRecord(t *testing.T) {
	gameRecord := "Game 13"
	want := 13

	got, _ := GetIdForGame(gameRecord)
	if got != want {
		t.Errorf("Didn't find the game id correctly. Found: %d\n", got)
	}
}
