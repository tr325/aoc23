package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFirstAndLastNumbers(line string) (int, int) {
	var first = -1
	var last = -1

	s := bufio.NewScanner(strings.NewReader(line))
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			// ignore
		} else {
			if first == -1 {
				first = i
			}
			last = i
		}
	}

	return first, last
}

func ConcatTwoDigits(first int, second int) int {
	concat, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
	return concat
}

func getSumOfFirstAndLastNumbers(line string) int {
	first, last := GetFirstAndLastNumbers(line)
	return ConcatTwoDigits(first, last)
}

func main() {
	readFile, _ := os.Open("calibration.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var sum = 0
	for fileScanner.Scan() {
		sum = sum + getSumOfFirstAndLastNumbers(fileScanner.Text())
	}

	fmt.Printf("Sum of all the calibration numbers: %d\n", sum)
}
