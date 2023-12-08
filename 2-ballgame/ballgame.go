package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


// Part 1
func GetIdForGame(gameTitle string) (int, error) {
	gameNumber := strings.Replace(gameTitle, "Game ", "", -1)
	return strconv.Atoi(gameNumber)
}

func IsHandfulPossible(handful string) bool {
	limits := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	hh := strings.Replace(handful, ";", "", -1)
	cubes := strings.Split(hh, ", ")
	for _, c := range cubes {
		totAndColour := strings.Split(c, " ")
		total, _ := strconv.Atoi(totAndColour[0])
		if total > limits[totAndColour[1]] {
			return false
		}
	}
	return true
}

func IsGamePossible(gameScore string) bool {
	handfuls := strings.Split(gameScore, "; ")
	for _, h := range handfuls {
		if !IsHandfulPossible(h) {
			return false
		}
	}
	return true
}

// ------------------------------------------------------------

// Part 2

func GetGamePower(gameScore string) int {
	required := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	handfuls := strings.Split(gameScore, "; ")
	for _, h := range handfuls {
		cubes := strings.Split(h, ", ")
		for _, c := range cubes {
			totAndColour := strings.Split(c, " ")
			total, _ := strconv.Atoi(totAndColour[0])
			if total > required[totAndColour[1]] {
				required[totAndColour[1]] = total
			}
		}
	}

	return required["red"]*required["green"]*required["blue"]
}

// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("ballgame.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var possibleIdsSum = 0
	var powerSum = 0
	for fileScanner.Scan() {
		gameRecord := fileScanner.Text()
		idAndGame := strings.Split(gameRecord, ": ")

		// Part 1
		validresult := IsGamePossible(idAndGame[1])
		if validresult == true {
			id, _ := GetIdForGame(idAndGame[0])
			possibleIdsSum = possibleIdsSum + id
		}

		// Part 2
		powerSum = powerSum + GetGamePower(idAndGame[1])
	}

	fmt.Printf("Sum of all the possible game IDs: %d\n", possibleIdsSum)
	fmt.Printf("Sum of all the game powers: %d\n", powerSum)
}
