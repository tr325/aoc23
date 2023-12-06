package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetIdForGame(gameRecord string) (int, error) {
	pattern := regexp.MustCompile(`Game \d\d`)
	substr := pattern.Find([]byte(gameRecord))
	gameNumber := strings.Replace(string(substr), "Game ", "", -1)
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
		validresult := GamePossible(gameRecord)
		if validresult == true {
			id, _ := GetIdForGame(gameRecord)
			sum = sum + id
		}
	}

	fmt.Printf("Sum of all the calibration numbers: %d\n", sum)
}