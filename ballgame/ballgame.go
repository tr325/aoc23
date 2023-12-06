package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetIdForGame(gameRecord string) int {
	// TODO
	return 1
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
			sum = sum + GetIdForGame(gameRecord)
		}
	}

	fmt.Printf("Sum of all the calibration numbers: %d\n", sum)
}