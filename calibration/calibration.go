package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func GetDigitFromMatchedLocation(line string, loc []int) int {
	var myMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var digit int

	substr := string(line[loc[0]:loc[1]])
	if loc[0] == loc[1]-1 {
		// Digit found
		digit, _ = strconv.Atoi(substr)
	} else {
		digit = myMap[substr]
	}
	return digit
}

func GetFirstAndLastNumbers(line string) (int, int) {

	pattern := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[0-9])`)
	locs := pattern.FindAllIndex([]byte(line), -1)

	firstMatch := locs[0]
	lastMatch := locs[len(locs)-1]
	first := GetDigitFromMatchedLocation(line, firstMatch)
	last := GetDigitFromMatchedLocation(line, lastMatch)

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
