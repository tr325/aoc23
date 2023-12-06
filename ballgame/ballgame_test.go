package main

import ("testing")

func TestGetIdFromGameRecord(t *testing.T) {
	gameRecord := "Game 13: 2 green, 2 red, 3 blue; 3 blue, 3 red, 3 green;"
	want := 13

	got, _ := GetIdForGame(gameRecord)
	if got != want {
		t.Errorf("Didn't find the game id correctly. Found: %d\n", got)
	}
}
