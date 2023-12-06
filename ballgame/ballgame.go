package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	readFile, _ := os.Open("ballgame.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var sum = 0
	for fileScanner.Scan() {
		gameRecord := fileScanner.Text()
		idAndGame := strings.Split(gameRecord, ": ")
		validresult := IsGamePossible(idAndGame[1])
		if validresult == true {
			id, _ := GetIdForGame(idAndGame[0])
			sum = sum + id
		}
	}

	fmt.Printf("Sum of all the calibration numbers: %d\n", sum)
}
