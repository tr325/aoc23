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

func GamePossible(gameRecord string) bool {
	// TODO
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
		validresult := GamePossible(idAndGame[1])
		if validresult == true {
			id, _ := GetIdForGame(idAndGame[0])
			sum = sum + id
		}
	}

	fmt.Printf("Sum of all the calibration numbers: %d\n", sum)
}
