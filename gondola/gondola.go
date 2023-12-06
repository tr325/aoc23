package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Part struct {
	number int
	row int
	start int
	end int
}

type SymbolLocation struct {
	row int
	col int
}

func abs(x int) int {
    if x < 0 {
            return -x
    }
    return x
}

func IsPartNumberRelevant(part Part, symbols []SymbolLocation) bool {
	for _, symbol := range symbols {
		withinWidth := ((symbol.col + 1) >= part.start) && ((symbol.col - 1) <= part.end)
		adjacentRows := (abs(symbol.row - part.row) <= 1)
		if withinWidth && adjacentRows {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------

func main() {
	readFile, _ := os.Open("engine.txt")

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	numberFinder := regexp.MustCompile(`[0-9]*`)
	symbolFinder := regexp.MustCompile(`[^A-Za-z0-9\.]`)

	var row = 0
	parts := []Part{}
	symbolLocations := []SymbolLocation{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := numberFinder.FindAllString(line, -1)
		numberLocs := numberFinder.FindAllStringIndex(line, -1)
		symbols := symbolFinder.FindAllStringIndex(line, -1)
		for i, num := range numbers {
			number, _ := strconv.Atoi(num)
			parts = append(parts, Part{number, row, numberLocs[i][0], numberLocs[i][1]-1})
		}
		for _, loc := range symbols {
			symbolLocations = append(symbolLocations, SymbolLocation{row, loc[0]})
		}
		row = row + 1
	}

	var partNumbersSum = 0
	for _, part := range parts {
		if IsPartNumberRelevant(part, symbolLocations) {
			partNumbersSum = partNumbersSum + part.number
		}
	}

	fmt.Printf("Sum of all the active part numbers: %d\n", partNumbersSum)
}
